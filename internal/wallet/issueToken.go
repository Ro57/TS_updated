package wallet

import (
	"context"
	"token-strike/tsp2p/server/rpcservice"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddPeer append new peer to peer slice
func (s Server) IssueToken(context.Context, *rpcservice.IssueTokenRequest) (*rpcservice.IssueTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implement now")
}
