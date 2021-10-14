package utils

func (suite *TestSuite) TestAnnounceData() {

	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		args    args
		want    func(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
		wantNum int32
	}{
		{
			name: "Valid annoncment num",
			args: args{
				payload: []byte{},
			},
			want:    suite.Equal,
			wantNum: suite.chain.CurrentHeight(),
		},
		{
			name: "Invalid annoncment num",
			args: args{
				payload: []byte{},
			},
			want:    suite.NotEqual,
			wantNum: 0,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			annChan := suite.chain.AnnounceData(tt.args.payload)
			tempAnn := <-annChan
			tt.want(tt.wantNum, tempAnn.Num, "error in %v want %v but got %v", tt.name, tt.wantNum, tempAnn.Num)
		})
	}
}
