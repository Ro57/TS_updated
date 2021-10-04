package replicatorrpc

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
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
	now := time.Now().Unix()

	var rpcSequances *replicator.GetUrlTokenResponse
	var jsonSequances *replicator.GetUrlTokenResponse

	tokenResponse := &replicator.GetUrlTokenResponse{
		State: &DB.State{
			Token: &DB.Token{
				Count:        100,
				Creation:     now,
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
			name: "get block sequances by rpc",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetBlockSequenceRequest{
					Name: "smt",
				},
			},
			want:       tokenResponse,
			wantRpc:    true,
			wantJson:   false,
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "get block sequances by json",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetBlockSequenceRequest{
					Name: "smt",
				},
			},
			want:       tokenResponse,
			wantRpc:    false,
			wantJson:   true,
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "json and rpc equal",
			args: args{
				ctx: context.Background(),
				req: &replicator.GetBlockSequenceRequest{
					Name: "smt",
				},
			},
			want:       tokenResponse,
			wantRpc:    false,
			wantJson:   false,
			wantErr:    true,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			if tt.wantRpc {
				got, err := suite.grpcClient.GetBlockSequence(tt.args.ctx, tt.args.req)
				if tt.wantErr {
					require.EqualError(t, err, tt.wantErrMsg)
					return
				}

				rpcSequances = got

				require.Equal(t, got.String(), tt.want.String())
			}
			if tt.wantJson {
				res, err := http.Get(domain + "/v2/replicator/blocksequence/smt")
				if tt.wantErr {
					require.EqualError(t, err, tt.wantErrMsg)
					return
				}
				if err != nil {
					t.Fatal(err)
					return
				}

				body, err := ioutil.ReadAll(res.Body)
				if tt.wantErr {
					require.EqualError(t, err, tt.wantErrMsg)
					return
				}
				if err != nil {
					t.Fatal(err)
					return
				}

				err = json.Unmarshal(body, &jsonSequances)
				if err != nil {
					t.Fatal(err)
					return
				}

				wantJson, err := json.Marshal(tt.want)
				if err != nil {
					t.Fatal(err)
					return
				}

				require.Equal(t, body, string(wantJson))
			}

			if tt.Equals {
				require.Equal(t, jsonSequances, rpcSequances)
			}
		})
	}
}
