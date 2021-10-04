package replicatorrpc

import (
	"context"
	"testing"

	"token-strike/tsp2p/server/replicator"

	"github.com/stretchr/testify/require"
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
			name: "valid",
			args: args{
				ctx: context.Background(),
				req: &replicator.GenerateURLRequest{
					Name: "tokenA",
				},
			},
			want: &replicator.GenerateURLResponse{
				Url: domain + "/v2/replicator/blocksequence/tokenA",
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.grpcClient.GenerateURL(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.EqualError(t, err, tt.wantErrMsg)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}
