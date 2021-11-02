package issuer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

func (i *Issuer) IssueToken(ctx context.Context, request *rpcservice.IssueTokenRequest) (*rpcservice.IssueTokenResponse, error) {
	token := &DB.Token{
		Count:        10,
		Expiration:   request.Expiration,
		Creation:     time.Now().Unix(),
		IssuerPubkey: i.address.String(),
		Urls: []string{
			"http://localhost:3333/token1",
		},
	}

	state := &DB.State{
		Token:  token,
		Owners: request.Owners,
		Locks:  nil,
	}

	stateBytes, err := state.GetHash()
	if err != nil {
		return nil, err
	}

	block := &DB.Block{
		PrevBlock:      "",
		Justifications: nil,
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
		PktBlockHash:   string(i.config.Chain.BlockHashAtHeight(i.config.Chain.CurrentHeight())),
		PktBlockHeight: i.config.Chain.CurrentHeight(),
		Height:         0,
	}

	err = block.Sing(i.private)
	if err != nil {
		return nil, err
	}

	blockSigned, err := proto.Marshal(block)
	if err != nil {
		return nil, err
	}

	blockHash := sha256.Sum256(blockSigned)
	tokenID := hex.EncodeToString(blockHash[:])

	err = i.tokendb.SaveIssuerTokenDB(tokenID, i.address.String())
	if err != nil {
		return nil, err
	}

	err = i.tokendb.IssueTokenDB(tokenID, token, block, state)
	if err != nil {
		return nil, err
	}

	return &rpcservice.IssueTokenResponse{
		TokenId: tokenID,
	}, i.sendBlock(tokenID, blockHash, block)
}

func (i *Issuer) sendBlock(tokenID string, blockHash [32]byte, block *DB.Block) error {
	var genError error

	for index, peer := range i.peers {
		conn, err := grpc.DialContext(
			context.TODO(),
			peer,
			grpc.WithInsecure(),
		)
		if err != nil {
			genError = fmt.Errorf("%v : %s /n %s", index, err, genError)
		}

		client := rpcservice.NewRPCServiceClient(conn)

		resp, err := client.Inv(
			context.Background(),
			&tokenstrike.InvReq{Invs: []*tokenstrike.Inv{
				{
					Parent:     blockHash[:],
					Type:       tokenstrike.TYPE_BLOCK,
					EntityHash: blockHash[:],
				},
			}},
		)
		if err != nil {
			genError = fmt.Errorf("%v : %s /n %s", index, err, genError)
		}

		if resp.Needed != nil {
			for _, need := range resp.Needed {
				if need {
					DataReq := &tokenstrike.Data{
						Data: &tokenstrike.Data_Block{Block: block},
					}

					//send selected lock and NOW skip check of warning
					_, err := client.PostData(context.TODO(), DataReq)
					if err != nil {
						genError = fmt.Errorf("%v : %s /n %s", index, err, genError)
					}
				}
			}
		}
	}

	return genError
}
