package wallet

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"token-strike/internal/utils/idgen"
	"token-strike/internal/utils/tokenstrikemock"
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

	_ = s.inv.Insert(tokenstrikemock.MempoolEntry{
		ParentHash: req.TokenId,
		Expiration: 123,
		Type:       tokenstrike.TYPE_LOCK,
		Message:    lockEl,
	})

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
