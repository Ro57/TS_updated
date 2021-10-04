package replicatorrpc

import (
	"context"
	"testing"
	"time"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/replicator"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestIssueToken() {
	type args struct {
		ctx context.Context
		req *replicator.IssueTokenRequest
	}
	now := time.Now().Unix()

	smtMock := &DB.Block{
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
		Creation:       now,
		State:          "some hash",
		PktBlockHash:   "pkt hash",
		PktBlockHeight: 120,
		Height:         1,
		Signature:      "Signature",
	}

	block := &DB.Block{
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
		Creation:       now,
		State:          smtMock.State,
		PktBlockHash:   smtMock.PktBlockHash,
		PktBlockHeight: smtMock.PktBlockHeight,
		Height:         1,
		Signature:      smtMock.Signature,
	}

	tests := []struct {
		name       string
		args       args
		want       *DB.Block
		getFromDB  func() *DB.Block
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "issue_smt",
			args: args{
				ctx: context.Background(),
				req: &replicator.IssueTokenRequest{
					Name: "smt",
					Offer: &DB.Token{
						Count:        100,
						Creation:     now,
						Expiration:   1000000,
						IssuerPubkey: "someHash",
						Urls: []string{
							domain,
						},
					},
					Block:     block,
					Recipient: []*DB.Owner{},
				},
			},
			want:       block,
			getFromDB:  func() *DB.Block { return smtMock },
			wantErr:    false,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			_, err := suite.grpcClient.IssueToken(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.EqualError(t, err, tt.wantErrMsg)
				return
			}

			// use string for proto struct
			require.Equal(t, tt.getFromDB().String(), tt.want.String())
		})
	}
}
