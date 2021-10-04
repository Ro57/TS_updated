package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTokenList(ctx context.Context, req *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetTokenList not implimented")
}
