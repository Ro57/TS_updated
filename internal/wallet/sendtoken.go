package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"token-strike/internal/utils/idgen"
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

	blockHash, err := hex.DecodeString(req.TokenId)
	if err != nil {
		return nil, err
	}

	transferInvs := []*tokenstrike.Inv{
		{
			Parent:     blockHash[:],
			Type:       tokenstrike.TYPE_TX,
			EntityHash: transferTokensHash[:],
		},
	}

	resp, err := s.issuerInvSlice[0].Inv(context.TODO(), &tokenstrike.InvReq{
		Invs: transferInvs,
	})
	if err != nil {
		return nil, err
	}

	if resp.Needed != nil {
		for _, need := range resp.Needed {
			if need {
				DataReq := &tokenstrike.Data{
					Data:  &tokenstrike.Data_Transfer{Transfer: transferTokens},
					Token: req.TokenId,
				}

				//send selected lock and NOW skip check of warning
				_, err := s.issuerInvSlice[0].PostData(context.TODO(), DataReq)
				if err != nil {
					return nil, err
				}
			}
		}
	}

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
