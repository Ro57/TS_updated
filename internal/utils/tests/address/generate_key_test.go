package utils_test

import (
	"crypto/ed25519"
	"encoding/hex"
	"token-strike/internal/utils/simple"
)

const (
	aliceIndex = iota
	bobIndex
	christyIndex
)

const byte32Size = 32

func randomSeed(offset int) [32]byte {
	byte32 := [byte32Size]byte{}

	for i := 0; i < byte32Size; i++ {
		byte32[i] = byte(i + offset)
	}

	return byte32
}

func (suite *TestSuite) TestGenerateKey() {
	seedSlice := [][byte32Size]byte{randomSeed(aliceIndex), randomSeed(bobIndex), randomSeed(christyIndex)}

	type args struct {
		seed [byte32Size]byte
	}

	tests := []struct {
		name     string
		args     []args
		wantKeys []ed25519.PrivateKey
		want     func(value bool, msgAndArgs ...interface{}) bool
	}{
		{
			name: "Single key genrated correct",
			args: []args{
				{
					seed: seedSlice[aliceIndex],
				},
			},
			wantKeys: []ed25519.PrivateKey{
				ed25519.NewKeyFromSeed(seedSlice[aliceIndex][:]),
			},
			want: suite.True,
		},
		{
			name: "Single key generated incorrect",
			args: []args{
				{
					seed: seedSlice[aliceIndex],
				},
			},

			wantKeys: []ed25519.PrivateKey{
				ed25519.NewKeyFromSeed(seedSlice[bobIndex][:]),
			},
			want: suite.False,
		},
		{
			name: "All keys generated correct",
			args: []args{
				{
					seed: seedSlice[aliceIndex],
				},
				{
					seed: seedSlice[bobIndex],
				},
				{
					seed: seedSlice[christyIndex],
				},
			},
			wantKeys: []ed25519.PrivateKey{
				ed25519.NewKeyFromSeed(seedSlice[aliceIndex][:]),
				ed25519.NewKeyFromSeed(seedSlice[bobIndex][:]),
				ed25519.NewKeyFromSeed(seedSlice[christyIndex][:]),
			},
			want: suite.True,
		},
		{
			name: "All keys genrated incorrect",
			args: []args{
				{
					seed: seedSlice[aliceIndex],
				},
				{
					seed: seedSlice[bobIndex],
				},
				{
					seed: seedSlice[christyIndex],
				},
			},
			wantKeys: []ed25519.PrivateKey{
				ed25519.NewKeyFromSeed(seedSlice[bobIndex][:]),
				ed25519.NewKeyFromSeed(seedSlice[christyIndex][:]),
				ed25519.NewKeyFromSeed(seedSlice[aliceIndex][:]),
			},
			want: suite.False,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			for i, a := range tt.args {
				key := suite.addressScheme.GenerateKey(a.seed)
				wantKey := simple.SimplePrivateKey{Key: tt.wantKeys[i]}

				tt.want(key.Equal(wantKey), "error in test %v (private) want %v but got %v", tt.name, wantKey, key)

				wantPublic := tt.wantKeys[i].Public().(ed25519.PublicKey)
				wantPublicHash := hex.EncodeToString(wantPublic)
				gotPublic := simple.NewSimpleAddress(key.GetPublicKey()).String()

				tt.want(gotPublic == wantPublicHash, "error in test %v (public) want %v but got %v", tt.name, wantPublicHash, gotPublic)
			}
		})

	}

}
