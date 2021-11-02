package wallet

import (
	"context"
	"token-strike/tsp2p/server/rpcservice"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddPeer append new peer to peer slice
func (s Server) AddPeer(ctx context.Context, req *rpcservice.PeerRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "not implement now")
}
