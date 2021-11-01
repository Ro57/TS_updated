package wallet

import (
	"context"
	"token-strike/tsp2p/server/tokenstrike"
)

func (s Server) Inv(ctx context.Context, req *tokenstrike.InvReq) (*tokenstrike.InvResp, error) {
	return s.inv.Inv(ctx, req)
}

// PostData — send full data to replication
func (s Server) PostData(ctx context.Context, req *tokenstrike.Data) (*tokenstrike.PostDataResp, error) {
	return s.inv.PostData(ctx, req)
}

// GetTokenStatus — response with information about token
func (s Server) GetTokenStatus(ctx context.Context, req *tokenstrike.TokenStatusReq) (*tokenstrike.TokenStatus, error) {
	return s.inv.GetTokenStatus(ctx, req)
}
