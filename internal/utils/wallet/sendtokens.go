package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

func (s *SimpleWallet) SendTokens(tokenId string, lockId []byte, secretHex []byte) ([]byte, error) {
	transferTokens := &tokenstrike.TransferTokens{
		Htlc: secretHex[:],
		Lock: lockId[:],
	}

	transferTokensB, err := proto.Marshal(transferTokens)
	if err != nil {
		return nil, err
	}

	transferTokensHash := sha256.Sum256(transferTokensB)

	blockHash, err := hex.DecodeString(tokenId)
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
					Data: &tokenstrike.Data_Transfer{Transfer: transferTokens},
				}

				//send selected lock and NOW skip check of warning
				_, err := s.issuerInvSlice[0].PostData(context.TODO(), DataReq)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return transferTokensHash[:], nil
}
