package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUrlSequence(ctx context.Context, req *replicator.GetUrlSequenceRequest) (*replicator.GetUrlSequenceResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetUrlSequence not implimented")
}
