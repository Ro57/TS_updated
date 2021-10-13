package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	address Address
}

func (suite *TestSuite) SetupTest() {
	suite.address = Address{}
}
