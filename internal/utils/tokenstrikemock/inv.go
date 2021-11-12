package tokenstrikemock

import (
	"context"
	"encoding/hex"
	"errors"
	"token-strike/tsp2p/server/tokenstrike"
)

func (t *TokenStrikeMock) Inv(ctx context.Context, req *tokenstrike.InvReq) (*tokenstrike.InvResp, error) {
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

func (t *TokenStrikeMock) selectNeeded(inv *tokenstrike.Inv) bool {
	entity := hex.EncodeToString(inv.EntityHash)
	// parent := hex.EncodeToString(inv.Parent)

	// TODO: check only for issuer
	// if !t.isStoreToken(parent) {
	// 	return DontNeedData
	// }

	if _, ok := t.mempoolEntries[entity]; ok {
		return DontNeedData
	}

	// TODO: move invCache out of this function
	t.mempoolEntries[entity] = &MempoolEntry{
		ParentHash: string(inv.Parent),
		Expiration: 0,
		Type:       inv.Type,
		Message:    nil,
	}

	return NeedData
}

func (t TokenStrikeMock) isStoreToken(token string) bool {
	issuerTokens, err := t.bboltDB.GetIssuerTokens()
	if err != nil {
		return false
	}

	tokenSlice := issuerTokens.GetToken(t.address.String())

	for _, t := range tokenSlice {
		if t == token {
			return true
		}
	}

	return false
}
