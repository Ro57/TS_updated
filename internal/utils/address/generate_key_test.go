package utils

import (
	"crypto/ed25519"
)

func randomSeed(l int) []byte {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(i)
	}
	return bytes
}

func (suite *TestSuite) TestGenerateKey() {
	seed := randomSeed(32)
	wantKey := ed25519.NewKeyFromSeed(seed)

	addressHelper := Address{}
	key := addressHelper.GenerateKey(seed)

	suite.True(key.Equal(wantKey))
}
