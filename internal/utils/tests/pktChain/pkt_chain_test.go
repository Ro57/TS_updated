package utils_test

import (
	"testing"
	"token-strike/internal/utils/pktchain"

	"github.com/stretchr/testify/suite"
)

// Make a test dummy implementation MockPktChain of this where:
// CurrentHeight() -> (unixTime() - 1566269808) / 60
// BlockHashAtHeight(height) -> if height > CurrentHeight() { nil } else { sha256(height) }
// AnnounceData(data) -> go func() { loop { sleepSeconds(random(30, 90)); channel <- AnnProof { num: CurrentHeight() } } }
// VerifyProof(ap) -> return ap.num

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	chain pktchain.SimplePktChain
}

func (suite *TestSuite) SetupTest() {
	suite.chain = pktchain.SimplePktChain{}
}
