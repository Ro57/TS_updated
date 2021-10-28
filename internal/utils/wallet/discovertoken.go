package wallet

import (
	"context"
	"errors"
	"token-strike/tsp2p/server/tokenstrike"
)

const emptyIssue = "empty issue url collection"

func (s SimpleWallet) DiscoverToken(tokenID string) error {

	req := &tokenstrike.TokenStatusReq{
		Tokenid: tokenID,
	}

	tokenStatus, err := s.issuerInvSlice[0].GetTokenStatus(context.Background(), req)
	if err != nil {
		return err
	}
	if tokenStatus == nil {
		return errors.New("token not found")
	}

	return nil
}
