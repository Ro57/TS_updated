package utils_test

import (
	"os"
	"testing"
	"token-strike/internal/database/repository"
	addressScheme "token-strike/internal/utils/simple"
	utils "token-strike/internal/utils/tests"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	tokenStrike tokenstrikemock.TokenStrikeMock

	sendingMessageTestData struct {
		block  *DB.Block
		invs   []*tokenstrike.Inv
		hash   []byte
		wallet *walletMockServer
	}
}

func randomSeed(l, offset int) [32]byte {
	bytes := [32]byte{}
	for i := 0; i < l; i++ {
		bytes[i] = byte(i + offset)
	}
	return bytes
}

func (suite *TestSuite) SetupTest() {
	db, path := utils.InitTempDatabase(suite.T())
	defer os.RemoveAll(path)
	defer utils.CloseDB(db, suite.T())

	tokendb := repository.NewBbolt(db)

	activeAddressScheme := &addressScheme.SimpleAddressScheme{}
	privKey := activeAddressScheme.GenerateKey(randomSeed(32, 0))

	suite.tokenStrike = *tokenstrikemock.New(tokendb, privKey.Address())
}
