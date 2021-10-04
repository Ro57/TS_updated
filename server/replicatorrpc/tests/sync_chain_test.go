package replicatorrpc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"token-strike/tsp2p/server/replicator"
)

func (suite *TestSuite) TestSyncChain() {
	type args struct {
		ctx context.Context
		req *replicator.SyncChainRequest
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "valid",
			args: args{
				ctx: context.Background(),
				req: &replicator.SyncChainRequest{
					Name:   "t1",
					Blocks: nil,
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			c, err := suite.grpcClient.SyncChain(tt.args.ctx)
			require.NoError(t, err)
			err = c.Send(tt.args.req)
			if tt.wantErr {
				require.EqualError(t, err, tt.wantErrMsg)
				return
			}
			require.NoError(t, err)
		})
	}

}
