package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/replicator"
)

func (s *Server) GetHeaders(ctx context.Context, req *replicator.GetHeadersRequest) (*replicator.GetHeadersResponse, error) {
	blocks, err := s.db.GetMerkleBlockDB(req.TokenId, req.Hash)
	if err != nil {
		return nil, err
	}

	return &replicator.GetHeadersResponse{
		Token:  &DB.Token{},
		Blocks: blocks,
	}, nil
}
