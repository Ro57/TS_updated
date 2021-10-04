package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) IssueToken(ctx context.Context, req *replicator.IssueTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "IssueToken not implimented")
}
