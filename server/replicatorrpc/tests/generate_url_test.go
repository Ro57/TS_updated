package replicatorrpc

import (
	"context"
	"testing"

	"token-strike/tsp2p/server/replicator"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			name: "not implemented",
			args: args{
				ctx: context.Background(),
				req: &replicator.GenerateURLRequest{
					Name: "",
				},
			},
			want:       nil,
			wantErr:    true,
			wantErrMsg: status.Error(codes.Unimplemented, "GenerateURL not implemented").Error(),
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
