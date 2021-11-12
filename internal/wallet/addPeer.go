package wallet

import (
	"context"
	"token-strike/tsp2p/server/rpcservice"

	"github.com/golang/protobuf/ptypes/empty"
)

// AddPeer append new peer to peer slice
func (s Server) AddPeer(ctx context.Context, req *rpcservice.PeerRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.inv.AddPeer(req.Url)
}
