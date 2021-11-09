package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

func (s Server) LockToken(ctx context.Context, req *rpcservice.LockTokenRequest) (*rpcservice.LockTokenResponse, error) {
	//populate lock with special data todo check the addresses
	lockEl := &lock.Lock{
		Count:          3,
		Recipient:      req.GetRecipient(),
		Sender:         s.address.String(),
		HtlcSecretHash: req.GetSecretHash(),
		ProofCount:     s.pkt.CurrentHeight() + 60,
		PktBlockHash:   s.pkt.BlockHashAtHeight(s.pkt.CurrentHeight()),
		PktBlockHeight: uint32(s.pkt.CurrentHeight()),
		Signature:      "",
	}

	// skip genesis block
	s.inv.AwaitJustification(req.TokenId, nil)

	err := lockEl.Sing(s.privateKey)
	if err != nil {
		return nil, err
	}

	lockSigned, err := proto.Marshal(lockEl)
	if err != nil {
		return nil, err
	}

	lockHash := sha256.Sum256(lockSigned)

	blockHash, err := hex.DecodeString(req.GetTokenId())
	if err != nil {
		return nil, err
	}

	invs := []*tokenstrike.Inv{
		{
			Parent:     blockHash,
			Type:       tokenstrike.TYPE_LOCK,
			EntityHash: lockHash[:],
		},
	}

	resp, err := s.issuerInvSlice[0].Inv(context.Background(), &tokenstrike.InvReq{
		Invs: invs,
	})

	if err != nil {
		return nil, err
	}

	if resp.Needed != nil {
		for _, need := range resp.Needed {
			if need {
				DataReq := &tokenstrike.Data{
					Data: &tokenstrike.Data_Lock{lockEl},
				}

				//send selected lock and NOW skip check of warning
				_, err := s.issuerInvSlice[0].PostData(context.TODO(), DataReq)
				if err != nil {
					return nil, err
				}
				s.db.LockToken(req.TokenId, lockEl)
			}
		}
	}

	id, err := s.inv.AwaitJustification(req.TokenId, lockHash[:])
	if err != nil {
		return nil, err
	}

	return &rpcservice.LockTokenResponse{LockId: *id}, nil
}
