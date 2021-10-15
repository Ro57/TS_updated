package utils_test

import (
	"crypto/ed25519"
)

func (suite *TestSuite) TestSign() {
	seed := randomSeed(32, 0)
	wantMsg := "Hello pkt"
	wantKey := ed25519.NewKeyFromSeed(seed)
	wantSingature := ed25519.Sign(wantKey, []byte(wantMsg))

	key := suite.address.GenerateKey(seed)

	singature := suite.address.Sign(wantKey, []byte(wantMsg))
	suite.True(key.Equal(wantKey))

	for i, b := range singature {
		suite.Equal(b, wantSingature[i], "error with signature bytes")
	}
}
