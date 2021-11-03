package utils_test

import (
	"testing"
	addressScheme "token-strike/internal/utils/address_scheme"

	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	addressScheme addressScheme.SimpleAddressScheme
}

func (suite *TestSuite) SetupTest() {
	suite.addressScheme = addressScheme.SimpleAddressScheme{}
}
