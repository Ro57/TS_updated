package replicatorrpc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/replicator"
)

func (suite *TestSuite) TestGetTokenList() {
	type args struct {
		ctx context.Context
		req *replicator.GetTokenListRequest
	}
	tests := []struct {
		name       string
		args       args
		want       *replicator.GetTokenListResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "valid",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetTokenListRequest{
					IssuerId: "issuerid",
					Params:   nil,
				},
			},
			want: &replicator.GetTokenListResponse{
				Tokens: []*replicator.Token{
					{
						Name: "t1",
						Token: &DB.Token{
							Count:        5,
							Expiration:   99999999,
							Creation:     999,
							IssuerPubkey: "issuerid",
						},
						Root: "some root 1",
					},
					{
						Name: "t2",
						Token: &DB.Token{
							Count:        5,
							Expiration:   99999999,
							Creation:     9999,
							IssuerPubkey: "issuerid",
						},
						Root: "some root 2",
					},
				},
				Total: 10,
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.grpcClient.GetTokenList(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.EqualError(t, err, tt.wantErrMsg)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
