package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"token-strike/internal/utils/idgen"
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

	dispatcher := s.inv.Subscribe(req.TokenId)
	txBlock := <-dispatcher.Block

	number, err := idgen.EntityIndex(txBlock.Content, transferTokensHash[:])
	if err != nil {
		return nil, err
	}

	txBlockBytes, err := txBlock.Content.GetHash()
	if err != nil {
		return nil, err
	}

	txBlockHash := hex.EncodeToString(txBlockBytes)

	id := idgen.Encode(txBlockHash, number)

	return &rpcservice.TransferTokensResponse{Txid: id}, nil
}
