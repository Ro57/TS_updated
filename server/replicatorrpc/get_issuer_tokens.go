package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetIssuerTokens(ctx context.Context, req *replicator.GetIssuerTokensRequest) (*replicator.GetIssuerTokensResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetIssuerTokens not implimented")
}
