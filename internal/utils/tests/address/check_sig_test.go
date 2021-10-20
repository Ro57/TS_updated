package utils_test

import (
	"math/rand"
	"time"
	"token-strike/internal/types"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomBytes() [32]byte {
	rand.Seed(time.Now().UnixNano())

	const size = 32
	bytes := [size]byte{}

	for i := 0; i < size; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}

	return bytes
}

func (suite *TestSuite) TestCheckSig() {
	someData := []byte("some data for sig")
	seedByteOne, seedByteTwo := randomBytes(), randomBytes()
	keyOne, keyTwo := suite.addressScheme.GenerateKey(seedByteOne), suite.addressScheme.GenerateKey(seedByteTwo)
	sigOne := keyOne.Sign(someData)
	sigTwo := keyTwo.Sign(someData)
	addressOne := keyOne.Address()
	addressTwo := keyTwo.Address()

	type args struct {
		signature []byte
		data      []byte
	}
	tests := []struct {
		name    string
		address types.Address
		args    args
		want    bool
	}{
		{
			name:    "Valid sig",
			address: addressOne,
			args: args{
				signature: sigOne,
				data:      someData,
			},
			want: true,
		},
		{
			name:    "Invalid incorrect address",
			address: addressOne,
			args: args{
				signature: sigTwo,
				data:      someData,
			},
			want: false,
		},
		{
			name:    "Invalid incorrect data",
			address: addressTwo,
			args: args{
				signature: sigTwo,
				data:      []byte("incorrect"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got := tt.address.CheckSig(tt.args.data, tt.args.signature)

			suite.Require().Equal(tt.want, got)
		})
	}

}
