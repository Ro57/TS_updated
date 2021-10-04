package replicatorrpc

import (
	"context"
	"testing"
	"time"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/replicator"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestGetHeaders() {
	type args struct {
		ctx context.Context
		req *replicator.GetHeadersRequest
	}

	now := time.Now().Unix()

	firstBlockDB := "block 1"

	tests := []struct {
		name       string
		args       args
		want       *replicator.GetHeadersResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "get headers from first",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetHeadersRequest{
					TokenId: "smt",
					Hash:    "some hash in blockchain",
				},
			},
			want: &replicator.GetHeadersResponse{
				Token: &DB.Token{
					Count:        100,
					Creation:     now,
					Expiration:   1000000,
					IssuerPubkey: "someHash",
					Urls: []string{
						domain,
					},
				},
				Blocks: []*replicator.MerkleBlock{
					&replicator.MerkleBlock{
						Hash:     firstBlockDB,
						PrevHash: "",
					},
					&replicator.MerkleBlock{
						Hash:     "block 2",
						PrevHash: "block 1",
					},
					&replicator.MerkleBlock{
						Hash:     "block 3",
						PrevHash: "block 2",
					},
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.grpcClient.GetHeaders(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.EqualError(t, err, tt.wantErrMsg)
				return
			}

			require.Equal(t, got, tt.want)
		})
	}
}
