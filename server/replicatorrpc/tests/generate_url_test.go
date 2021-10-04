package replicatorrpc

import (
	"context"
	"reflect"
	"testing"

	"token-strike/tsp2p/server/replicator"
)

func (suite *TestSuite) TestGenerateURL() {

	type args struct {
		ctx context.Context
		req *replicator.GenerateURLRequest
	}
	tests := []struct {
		name       string
		args       args
		want       *replicator.GenerateURLResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "",
			args: args{
				ctx: nil,
				req: &replicator.GenerateURLRequest{
					Name: "",
				},
			},
			want:       nil,
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.grpcClient.GenerateURL(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
