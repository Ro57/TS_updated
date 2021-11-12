package wallet

import (
	"context"
	"crypto/sha256"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

func (s *Server) SendToken(ctx context.Context, req *rpcservice.TransferTokensRequest) (*rpcservice.TransferTokensResponse, error) {
	transferTokens := &tokenstrike.TransferTokens{
		Htlc:   req.Htlc,
		LockId: req.LockId,
	}

	transferTokensB, err := proto.Marshal(transferTokens)
	if err != nil {
		return nil, err
	}

	transferTokensHash := sha256.Sum256(transferTokensB)

	_ = s.inv.Insert(tokenstrikemock.MempoolEntry{
		ParentHash: "req.TokenId",
		Expiration: 123,
		Type:       tokenstrike.TYPE_TX,
		Message:    transferTokens,
	})

	id, err := s.inv.AwaitJustification(req.TokenId, transferTokensHash[:])
	if err != nil {
		return nil, err
	}

	return &rpcservice.TransferTokensResponse{Txid: *id}, nil
}
