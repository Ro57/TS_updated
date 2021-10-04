package replicatorrpc

import (
	"context"
	"testing"
	"time"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/replicator"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestGetIssuerTokens() {
	type args struct {
		ctx context.Context
		req *replicator.GetIssuerTokensRequest
	}
	now := time.Now().Unix()

	pubKey := "issuer_pubKey"

	tests := []struct {
		name       string
		args       args
		want       *replicator.GetIssuerTokensResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "get issuer token list",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetIssuerTokensRequest{
					Issuer: []string{
						pubKey,
					},
					Params: &replicator.Pagination{
						Limit:  10,
						Offset: 0,
					},
				},
			},
			want: &replicator.GetIssuerTokensResponse{
				Tokens: []*replicator.IssuerTokens{
					&replicator.IssuerTokens{
						Name: pubKey,
						Tokens: []*replicator.Token{
							&replicator.Token{
								Name: "smt",
								Token: &DB.Token{
									Count:        100,
									Creation:     now,
									Expiration:   1000000,
									IssuerPubkey: "someHash",
									Urls: []string{
										domain,
									},
								},
							},
						},
					},
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.grpcClient.GetIssuerTokens(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.EqualError(t, err, tt.wantErrMsg)
				return
			}

			require.Equal(t, got, tt.want)
		})
	}
}
