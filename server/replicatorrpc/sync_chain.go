package replicatorrpc

import (
	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SyncChain(stream replicator.Replicator_SyncChainServer) error {
	return status.Error(codes.Unimplemented, "SyncChain not implimented")
}
