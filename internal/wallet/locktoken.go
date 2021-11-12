package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"token-strike/internal/utils/idgen"
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

	dispatcher := s.inv.Subscribe(req.TokenId)
	// Skip genesis block from pool
	<-dispatcher.Block

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
					Data:  &tokenstrike.Data_Lock{lockEl},
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

	lockBlock := <-dispatcher.Block

	number, err := idgen.EntityIndex(lockBlock.Content, lockHash[:])
	if err != nil {
		return nil, err
	}

	lockBlockBytes, err := lockBlock.Content.GetHash()
	if err != nil {
		return nil, err
	}

	lockBlockHash := hex.EncodeToString(lockBlockBytes)

	id := idgen.Encode(lockBlockHash, number)
	return &rpcservice.LockTokenResponse{LockId: id}, nil
}
