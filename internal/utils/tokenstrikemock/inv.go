package tokenstrikemock

import (
	"bytes"
	"context"
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

		if checkNeedHash(inv.EntityHash) {
			resp.Needed = append(resp.Needed, NeedData)
			continue
		}

		resp.Needed = append(resp.Needed, DontNeedData)
	}

	return resp, nil
}

func checkNeedHash(hash []byte) bool {
	//we pretend that we take goodHash from db and then we compare it
	return bytes.Compare(hash, GoodHash) == 0
}
