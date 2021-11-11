package issuer

import (
	"context"
	"encoding/hex"
	"time"
	"token-strike/internal/utils/tokenstrikemock"

	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"
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
		PktBlockHash:   hex.EncodeToString(i.config.Chain.BlockHashAtHeight(i.config.Chain.CurrentHeight())),
		PktBlockHeight: i.config.Chain.CurrentHeight(),
		Height:         0,
	}

	err = block.Sing(i.private)
	if err != nil {
		return nil, err
	}

	blockHash, err := block.GetHash()
	if err != nil {
		return nil, err
	}

	tokenID := hex.EncodeToString(blockHash)

	err = i.tokendb.SaveIssuerTokenDB(tokenID, i.address.String())
	if err != nil {
		return nil, err
	}

	err = i.tokendb.IssueTokenDB(tokenID, token, block, state)
	if err != nil {
		return nil, err
	}

	_ = i.invServer.Insert(
		tokenstrikemock.MempoolEntry{
			Hash:       tokenID,
			Type:       tokenstrike.TYPE_BLOCK,
			Message:    block,
			Expiration: 123,
		})

	return &rpcservice.IssueTokenResponse{
		TokenId: tokenID,
	}, nil
}
