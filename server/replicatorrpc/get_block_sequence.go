package replicatorrpc

import (
	"context"
	"token-strike/tsp2p/server/replicator"
)

func (s *Server) GetBlockSequence(ctx context.Context, req *replicator.GetBlockSequenceRequest) (*replicator.GetUrlTokenResponse, error) {
	chainInfo, err := s.db.GetChainInfoDB(req.Name)
	if err != nil {
		return nil, err
	}

	return &replicator.GetUrlTokenResponse{
		State:  chainInfo.GetState(),
		Blocks: chainInfo.GetBlocks(),
		Root:   chainInfo.GetRoot(),
	}, nil
}
