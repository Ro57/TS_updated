package utils_test

import (
	"crypto/ed25519"
	"token-strike/internal/types"
)

const (
	aliceIndex = iota
	bobIndex
	christyIndex
)

func randomSeed(l, offset int) []byte {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(i + offset)
	}
	return bytes
}

func (suite *TestSuite) TestGenerateKey() {
	seedSlice := [][]byte{randomSeed(32, 0), randomSeed(32, 32), randomSeed(32, 64)}
	type args struct {
		seed []byte
	}
	tests := []struct {
		name     string
		args     []args
		wantKeys []types.PrivateKey
		want     func(value bool, msgAndArgs ...interface{}) bool
	}{
		{
			name: "Single key genrated correct",
			args: []args{
				{
					seed: seedSlice[aliceIndex],
				},
			},
			wantKeys: []types.PrivateKey{
				ed25519.NewKeyFromSeed(seedSlice[aliceIndex]),
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

			wantKeys: []types.PrivateKey{
				ed25519.NewKeyFromSeed(seedSlice[bobIndex]),
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
			wantKeys: []types.PrivateKey{
				ed25519.NewKeyFromSeed(seedSlice[aliceIndex]),
				ed25519.NewKeyFromSeed(seedSlice[bobIndex]),
				ed25519.NewKeyFromSeed(seedSlice[christyIndex]),
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
			wantKeys: []types.PrivateKey{
				ed25519.NewKeyFromSeed(seedSlice[bobIndex]),
				ed25519.NewKeyFromSeed(seedSlice[christyIndex]),
				ed25519.NewKeyFromSeed(seedSlice[aliceIndex]),
			},
			want: suite.False,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			for i, a := range tt.args {
				key := suite.address.GenerateKey(a.seed)
				tt.want(key.Equal(tt.wantKeys[i]), "error in test %v want %v but got %v", tt.name, tt.wantKeys[i], key)
			}
		})

	}

}
