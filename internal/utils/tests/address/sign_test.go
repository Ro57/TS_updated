package utils_test

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"token-strike/internal/utils/address"
)

func (suite *TestSuite) TestSign() {
	seed := randomSeed(aliceIndex)
	wantMsg := "Hello pkt"
	wantKey := ed25519.NewKeyFromSeed(seed[:])
	wantSingature := ed25519.Sign(wantKey, []byte(wantMsg))

	key := suite.addressScheme.GenerateKey(seed)

	singature := key.Sign([]byte(wantMsg))

	wantPublic := wantKey.Public().(ed25519.PublicKey)
	wantPublicHash := hex.EncodeToString(wantPublic)

	suite.True(address.NewSimpleAddress(key.GetPublicKey()).String() == wantPublicHash)

	suite.True(bytes.Equal(singature, wantSingature), "error with signature bytes")
}
