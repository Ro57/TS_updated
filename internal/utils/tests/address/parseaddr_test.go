package utils_test

import "token-strike/internal/utils/address"

func (suite *TestSuite) TestParseAddr() {
	seed := randomSeed(aliceIndex)
	privKey := suite.addressScheme.GenerateKey(seed)
	msg := "Hello pkt"
	singature := privKey.Sign([]byte(msg))
	wantAddr := address.NewSimpleAddress(privKey.GetPublicKey())

	parsedAddr, err := suite.addressScheme.ParseAddr(wantAddr.String())
	suite.NoError(err, "on parse addr")

	suite.True(wantAddr.String() == parsedAddr.String())

	isSig := parsedAddr.CheckSig([]byte(msg), singature)

	suite.True(isSig, "sing check failed")

}
