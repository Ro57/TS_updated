package utils_test

import (
	"crypto/sha256"
)

func (suite *TestSuite) TestBlockHashAtHeight() {
	goodHeight := suite.chain.CurrentHeight() - 1000
	badHeight := goodHeight + 100000

	goodHash := sha256.Sum256([]byte(string(goodHeight)))

	tests := []struct {
		name     string
		height   int32
		wantHash []byte
	}{
		{
			name:     "Valid",
			height:   goodHeight,
			wantHash: goodHash[:],
		},
		{
			name:     "Invalid",
			height:   badHeight,
			wantHash: nil,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got := suite.chain.BlockHashAtHeight(tt.height)
			suite.Require().Equal(tt.wantHash, got)
		})
	}

}
