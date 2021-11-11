package utils_test

import (
	"context"
	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/ptypes/empty"
)

type walletMockServer struct {
	invs []*tokenstrike.Inv
	data *tokenstrike.Data
}

func (w *walletMockServer) SendToken(ctx context.Context, request *rpcservice.TransferTokensRequest) (*rpcservice.TransferTokensResponse, error) {
	panic("implement me")
}

func (w *walletMockServer) IssueToken(ctx context.Context, request *rpcservice.IssueTokenRequest) (*rpcservice.IssueTokenResponse, error) {
	panic("implement me")
}

func (w *walletMockServer) LockToken(ctx context.Context, request *rpcservice.LockTokenRequest) (*rpcservice.LockTokenResponse, error) {
	panic("implement me")
}

func (w *walletMockServer) AddPeer(ctx context.Context, request *rpcservice.PeerRequest) (*empty.Empty, error) {
	panic("implement me")
}

func (w *walletMockServer) Inv(ctx context.Context, req *tokenstrike.InvReq) (*tokenstrike.InvResp, error) {
	var response = &tokenstrike.InvResp{
		Needed: make([]bool, len(req.Invs)),
	}
	for i, inv := range req.Invs {
		response.Needed[i] = true
		w.invs = append(w.invs, inv)
	}
	return response, nil
}

func (w *walletMockServer) PostData(ctx context.Context, data *tokenstrike.Data) (*tokenstrike.PostDataResp, error) {
	w.data = data
	return &tokenstrike.PostDataResp{}, nil
}

func (w *walletMockServer) GetTokenStatus(ctx context.Context, req *tokenstrike.TokenStatusReq) (*tokenstrike.TokenStatus, error) {
	panic("implement me")
}
