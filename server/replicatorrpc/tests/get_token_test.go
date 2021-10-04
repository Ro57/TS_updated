package replicatorrpc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"token-strike/tsp2p/server/replicator"
)

func (suite *TestSuite) TestGetToken() {
	type args struct {
		ctx context.Context
		req *replicator.GetTokenRequest
	}
	tests := []struct {
		name       string
		args       args
		want       *replicator.GetTokenResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "valid",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetTokenRequest{
					TokenId: "t1",
				},
			},
			want:       nil,
			wantErr:    false,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.grpcClient.GetToken(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.EqualError(t, err, tt.wantErrMsg)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
