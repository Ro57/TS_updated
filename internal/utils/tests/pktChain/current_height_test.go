package utils_test

import "time"

func (suite *TestSuite) TestCurrentHeight() {
	tests := []struct {
		name string
		want int32
	}{
		{
			"Valid",
			int32((time.Now().Unix() - 1566269808) / 60),
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got := suite.chain.CurrentHeight()
			suite.Require().Equal(tt.want, got)
		})
	}

}
