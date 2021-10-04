package replicatorrpc

import (
	"context"
	"fmt"

	"token-strike/tsp2p/server/replicator"
)

const tokenUrlPattern = "%s/v2/replicator/blocksequence/%s"

func (s *Server) GenerateURL(ctx context.Context, req *replicator.GenerateURLRequest) (*replicator.GenerateURLResponse, error) {
	responseUrl := fmt.Sprintf(tokenUrlPattern, s.domain, req.Name)
	return &replicator.GenerateURLResponse{
		Url: responseUrl,
	}, nil
}
