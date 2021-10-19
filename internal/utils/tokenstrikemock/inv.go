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
	var resp *tokenstrike.InvResp
	for _, inv := range invs{
		if bytes.Compare(inv.Parent, GoodParent) != 0 && bytes.Compare(inv.EntityHash, GoodHash) != 0{
			resp.Needed = append(resp.Needed, NeedData)
			continue
		}
		resp.Needed = append(resp.Needed, DontNeedData)
	}

	return nil, nil
}