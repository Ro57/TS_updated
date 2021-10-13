package utils

import (
	"crypto/ed25519"
)

func (suite *TestSuite) TestSign() {
	seed := randomSeed(32)
	wantMsg := "Hello pkt"
	wantKey := ed25519.NewKeyFromSeed(seed)
	wantSingature := ed25519.Sign(wantKey, []byte(wantMsg))

	addressHelper := Address{}
	key := addressHelper.GenerateKey(seed)

	singature := addressHelper.Sign(wantKey, []byte(wantMsg))
	suite.True(key.Equal(wantKey))

	for i, b := range singature {
		suite.Equal(b, wantSingature[i], "error with signature bytes")
	}
}
