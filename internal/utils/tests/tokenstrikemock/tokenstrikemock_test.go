package utils_test

import (
	"testing"
	"token-strike/internal/database"
	"token-strike/internal/database/repository"
	"token-strike/internal/utils/address"
	addressScheme "token-strike/internal/utils/address_scheme"
	"token-strike/internal/utils/tokenstrikemock"

	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	tokenStrike tokenstrikemock.TokenStrikeMock
}

func randomSeed(l, offset int) [32]byte {
	bytes := [32]byte{}
	for i := 0; i < l; i++ {
		bytes[i] = byte(i + offset)
	}
	return bytes
}

func (suite *TestSuite) SetupTest() {
	db, err := database.Connect("./test.db")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			suite.Error(err)
		}

		err = db.Clear()
		if err != nil {
			suite.Error(err)
		}

	}()

	tokendb := repository.NewBbolt(db)

	activeAddressScheme := &addressScheme.SimpleAddressScheme{}
	privKey := activeAddressScheme.GenerateKey(randomSeed(32, 0))

	suite.tokenStrike = *tokenstrikemock.New(tokendb, address.NewSimpleAddress(privKey.GetPublicKey()))
}
