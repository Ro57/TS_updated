package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"token-strike/internal/utils/idgen"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/DB"
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

	// Skip genesis block from pool
	// <-s.dispatcher.Block

	err := lockEl.Sing(s.privateKey)
	if err != nil {
		return nil, err
	}

	lockSigned, err := proto.Marshal(lockEl)
	if err != nil {
		return nil, err
	}

	lockHash := sha256.Sum256(lockSigned)

	_ = s.inv.Insert(tokenstrikemock.MempoolEntry{
		ParentHash: req.TokenId,
		Expiration: 123,
		Type:       tokenstrike.TYPE_LOCK,
		Message:    lockEl,
	})

	wait := s.dispatcher.WaitBlockAction(func(b DB.Block) (string, error) {
		number, err := idgen.EntityIndex(b, lockHash[:])
		if err != nil {
			return "", err
		}

		lockBlockBytes, err := b.GetHash()
		if err != nil {
			return "", err
		}

		lockBlockHash := hex.EncodeToString(lockBlockBytes)

		id := idgen.Encode(lockBlockHash, number)

		return id, nil
	})

	lockID := <-wait

	if lockID.Err != nil {
		return &rpcservice.LockTokenResponse{}, lockID.Err
	}

	return &rpcservice.LockTokenResponse{LockId: lockID.ID}, nil
}
