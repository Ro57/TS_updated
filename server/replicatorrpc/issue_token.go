package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) IssueToken(ctx context.Context, req *replicator.IssueTokenRequest) (*emptypb.Empty, error) {
	err := s.db.IssueTokenDB(req.Name, req.Offer, req.Block, req.Recipient)
	if err != nil {
		return nil, err
	}

	err = s.db.SaveIssuerTokenDB(req.Name, req.Offer)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
