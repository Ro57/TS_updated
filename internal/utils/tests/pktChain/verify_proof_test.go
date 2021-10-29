package utils_test

import (
	"token-strike/internal/types/pkt"
)

func (suite *TestSuite) TestVerifyProof() {

	annChan := suite.chain.AnnounceData([]byte{})
	ann := <-annChan

	type args struct {
		ann pkt.AnnProof
	}
	tests := []struct {
		name    string
		args    args
		want    func(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
		wantNum int32
	}{
		{
			name:    "Valid proof",
			args:    args{ann},
			want:    suite.Equal,
			wantNum: suite.chain.CurrentHeight(),
		},
		{
			name:    "Invalid proof",
			args:    args{ann},
			want:    suite.NotEqual,
			wantNum: 0,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			num := suite.chain.VerifyProof(tt.args.ann)
			tt.want(tt.wantNum, num, "error in %v want %v but got %v", tt.name, tt.wantNum, num)
		})
	}

}
