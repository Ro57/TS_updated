package utils_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupTest() {
}
