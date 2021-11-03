package utils_test

func (suite *TestSuite) TestParseAddr() {
	seed := randomSeed(aliceIndex)
	privKey := suite.addressScheme.GenerateKey(seed)
	msg := "Hello pkt"
	singature := privKey.Sign([]byte(msg))
	wantAddr := privKey.Address()

	parsedAddr, err := suite.addressScheme.ParseAddr(wantAddr.String())
	suite.NoError(err, "on parse addr")

	suite.True(wantAddr.String() == parsedAddr.String())

	isSig := parsedAddr.CheckSig([]byte(msg), singature)

	suite.True(isSig, "sing check failed")

}
