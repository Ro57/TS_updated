package replicatorrpc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"token-strike/tsp2p/server/DB"
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
			want:       &replicator.GetTokenResponse{
				Token:                &replicator.Token{
					Name:                 "t1",
					Token:                &DB.Token{
						Count:                1000,
						Expiration:           999999999,
						Creation:             9999999,
						IssuerPubkey:         "some pub key",
					},
					Root:                 "someroot",
				},
				Discredits:           nil,
			},
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
