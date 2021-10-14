package utils

import "token-strike/internal/types"

func (suite *TestSuite) TestVerifyProof() {

	annChan := suite.chain.AnnounceData([]byte{})
	ann := <-annChan

	type args struct {
		ann types.AnnProof
	}
	tests := []struct {
		name    string
		args    args
		want    func(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
		wantNum int32
	}{
		{
			name:    "Valid proof",
			args:    ann,
			want:    suite.Equal,
			wantNum: 0,
		},
		{
			name:    "Invalid proof",
			args:    ann,
			want:    suite.NotEqual,
			wantNum: 1,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			num := suite.chain.VerifyProof(tt.args.ann)
			tt.want(tt.wantNum, num, "error in %v want %v but got %v", tt.name, tt.wantNum, num)
		})
	}

}
