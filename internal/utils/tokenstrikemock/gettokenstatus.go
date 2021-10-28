package tokenstrikemock

import (
	"context"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/tokenstrike"
)

func (t *TokenStrikeMock) GetTokenStatus(ctx context.Context, req *tokenstrike.TokenStatusReq) (*tokenstrike.TokenStatus, error) {
	token, err := t.bboltDB.GetToken(req.Tokenid)
	if err != nil {
		return nil, err
	}

	resp := &tokenstrike.TokenStatus{
		CurrentHeight: uint32(t.pktChain.CurrentHeight()),
		CurrentHash:   req.Tokenid,

		// TODO: change to real one
		Dblock0: &DB.Block{},
		State0:  &DB.State{Token: token.Token},
	}

	return resp, nil
}
