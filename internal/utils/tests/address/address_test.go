package utils_test

import (
	"testing"
	"token-strike/internal/utils"

	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	addressScheme utils.SimpleAddressScheme
}

func (suite *TestSuite) SetupTest() {
	suite.addressScheme = utils.SimpleAddressScheme{}
}
