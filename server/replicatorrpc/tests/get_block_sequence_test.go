package replicatorrpc

import (
	"context"
	"testing"
	"time"

	"token-strike/internal/errors"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/replicator"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestGetBlockSequence() {
	type args struct {
		ctx context.Context
		req *replicator.GetBlockSequenceRequest
	}

	tokenResponse := &replicator.GetUrlTokenResponse{
		State: &DB.State{
			Token: &DB.Token{
				Count:        100,
				Creation:     time.Now().Unix(),
				Expiration:   1000000,
				IssuerPubkey: "someHash",
				Urls: []string{
					domain,
				},
			},
			Owners: []*DB.Owner{},
			Locks:  []*lock.Lock{},
		},
		Blocks: []*DB.Block{
			&DB.Block{
				PrevBlock: "",
				Justifications: []*DB.Justification{
					&DB.Justification{
						Content: &DB.Justification_Genesis{
							Genesis: &justifications.Genesis{
								Token: "smt",
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name       string
		args       args
		wantRpc    bool
		want       *replicator.GetUrlTokenResponse
		wantJson   bool
		Equals     bool
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "valid request",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetBlockSequenceRequest{
					Name: "smt",
				},
			},
			want:       tokenResponse,
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "chain not found",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetBlockSequenceRequest{
					Name: "smt2",
				},
			},
			want:       tokenResponse,
			wantErr:    true,
			wantErrMsg: errors.TokenNotFoundErr.Error(),
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.grpcClient.GetBlockSequence(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.EqualError(t, err, tt.wantErrMsg)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
