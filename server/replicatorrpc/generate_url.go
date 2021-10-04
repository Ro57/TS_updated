package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GenerateURL(ctx context.Context, req *replicator.GenerateURLRequest) (*replicator.GenerateURLResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GenerateURL not implimented")
}
