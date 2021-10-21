package tokenstrikemock

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"token-strike/tsp2p/server/tokenstrike"
)

func (t TokenStrikeMock) Inv(ctx context.Context, req *tokenstrike.InvReq) (*tokenstrike.InvResp, error) {
	if req.Invs == nil {
		return nil, errors.New("empty Invs list")
	}

	invs := req.Invs
	resp := &tokenstrike.InvResp{}

	for _, inv := range invs {
		resp.Needed = append(resp.Needed, t.selectNeeded(inv))
	}

	return resp, nil
}

func (t TokenStrikeMock) selectNeeded(inv *tokenstrike.Inv) bool {
	if !t.containToken(hex.EncodeToString(inv.Parent)) {
		return DontNeedData
	}

	// TODO: rework with correct data
	if !bytes.Equal(inv.EntityHash, GoodHash) {
		return DontNeedData
	}

	return NeedData
}

func (t TokenStrikeMock) containToken(token string) bool {
	issuerTokens, err := t.bboltDB.GetIssuerTokens()
	if err != nil {
		return false
	}

	tokenSlice := issuerTokens.GetToken(t.issuer.String())

	for _, t := range tokenSlice {
		if t == token {
			return true
		}
	}

	return false
}
