package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"
)

func (s *Server) GetToken(ctx context.Context, req *replicator.GetTokenRequest) (*replicator.GetTokenResponse, error) {
	token, err := s.db.GetToken(req.TokenId)
	if err != nil {
		return nil, err
	}

	return &replicator.GetTokenResponse{
		Token: &token,
	}, nil
}
