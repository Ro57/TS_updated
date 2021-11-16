package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"token-strike/internal/utils/idgen"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/DB"
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
		ParentHash: req.TokenId,
		Expiration: 123,
		Type:       tokenstrike.TYPE_TX,
		Message:    transferTokens,
	})

	wait := s.dispatcher.WaitBlockAction(func(b DB.Block) (string, error) {
		number, err := idgen.EntityIndex(b, transferTokensHash[:])
		if err != nil {
			return "", err
		}

		txBlockBytes, err := b.GetHash()
		if err != nil {
			return "", err
		}

		txBlockHash := hex.EncodeToString(txBlockBytes)

		id := idgen.Encode(txBlockHash, number)

		return id, nil
	})

	txID := <-wait

	if txID.Err != nil {
		return &rpcservice.TransferTokensResponse{}, txID.Err
	}
	return &rpcservice.TransferTokensResponse{Txid: txID.ID}, nil
}
