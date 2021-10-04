package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetBlockSequence(ctx context.Context, req *replicator.GetBlockSequenceRequest) (*replicator.GetUrlTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetBlockSequence not implimented")
}
