package replicatorrpc

import (
	"context"

	"token-strike/tsp2p/server/replicator"
)

func (s *Server) GetIssuerTokens(ctx context.Context, req *replicator.GetIssuerTokensRequest) (*replicator.GetIssuerTokensResponse, error) {
	var (
		responseTokens = []*replicator.IssuerTokens{}
	)

	tokens, err := s.db.GetIssuerTokens()
	if err != nil {
		return nil, err
	}

	for _, issuer := range req.Issuer {
		issuerTokens := tokens.GetToken(issuer)

		quantityIssuerTokens := len(issuerTokens)
		if quantityIssuerTokens == 0 {
			return &replicator.GetIssuerTokensResponse{}, nil
		}

		issuerResponse := &replicator.IssuerTokens{
			Name:   issuer,
			Tokens: []*replicator.Token{},
		}

		for _, issuerToken := range issuerTokens {
			token, err := s.db.GetToken(issuerToken)
			if err != nil {
				return nil, err
			}
			issuerResponse.Tokens = append(issuerResponse.Tokens, &token)
		}

		responseTokens = append(responseTokens, issuerResponse)
	}

	return &replicator.GetIssuerTokensResponse{
		Tokens: responseTokens,
	}, nil
}
