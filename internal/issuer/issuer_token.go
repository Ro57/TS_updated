package issuer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"time"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/rpcservice"
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

	return &rpcservice.IssueTokenResponse{
		TokenId: tokenID,
	}, i.tokendb.IssueTokenDB(tokenID, token, block, state)
}
