package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"
)

func (s *Server) GetTokenList(ctx context.Context, req *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error) {
	resultList, err := s.db.GetTokenList()
	if err != nil {
		return nil, err
	}

	// Apply pagination
	if req.Params.Offset > 0 {
		if int(req.Params.Offset) <= len(resultList)-1 {
			resultList = resultList[req.Params.Offset:]
		} else {
			resultList = nil
		}
	}
	if req.Params.Limit > 0 {
		if int(req.Params.Limit) <= len(resultList)-1 {
			resultList = resultList[:req.Params.Limit]
		}
	}

	return &replicator.GetTokenListResponse{
		Tokens: resultList,
		Total:  int32(len(resultList)),
	}, nil
}
