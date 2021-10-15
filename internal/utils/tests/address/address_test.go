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
	address utils.Address
}

func (suite *TestSuite) SetupTest() {
	suite.address = utils.Address{}
}
