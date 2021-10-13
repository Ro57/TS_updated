package utils

import (
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) []byte {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return bytes
}

func getPublicBytes(private []byte) []byte {
	publicKey := make([]byte, 32)
	copy(publicKey, private[32:])
	return publicKey
}

func (suite *TestSuite) TestCheckSig() {
	//todo think about how generating data for testing in other way
	seedByteOne, seedByteTwo := randomString(32), randomString(32)
	keyOne, keyTwo := suite.address.GenerateKey(seedByteOne), suite.address.GenerateKey(seedByteTwo)
	someData := []byte("some data for sig")
	sigOne := suite.address.Sign(keyOne, someData)
	sigTwo := suite.address.Sign(keyTwo, someData)
	publicOne := string(getPublicBytes(keyOne))
	publicTwo := string(getPublicBytes(keyTwo))

	type args struct {
		address   string
		signature []byte
		data      []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid sig",
			args: args{
				address:   publicOne,
				signature: sigOne,
				data:      someData,
			},
			want: true,
		},
		{
			name: "Invalid incorrect address",
			args: args{
				address:   publicOne,
				signature: sigTwo,
				data:      someData,
			},
			want: false,
		},
		{
			name: "Invalid incorrect data",
			args: args{
				address:   publicTwo,
				signature: sigTwo,
				data:      []byte("incorrect"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got := suite.address.CheckSig(tt.args.address, tt.args.signature, tt.args.data)
			suite.Require().Equal(tt.want, got)
		})
	}

}
