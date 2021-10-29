package wallet

import (
	"crypto/sha256"
	"token-strike/internal/types/users"
	"token-strike/tsp2p/server/lock"

	"github.com/golang/protobuf/proto"
)

func (s SimpleWallet) LockTokens(args users.LockArgs) ([]byte, error) {
	//populate lock with special data todo check the addresses
	lockEl := &lock.Lock{
		Count:          3,
		Recipient:      args.GetRecipient(),
		Sender:         s.address.String(),
		HtlcSecretHash: args.GetSecretHash(),
		ProofCount:     s.pkt.CurrentHeight() + 60,
		PktBlockHash:   s.pkt.BlockHashAtHeight(s.pkt.CurrentHeight()),
		PktBlockHeight: uint32(s.pkt.CurrentHeight()),
		Signature:      "",
	}

	err := lockEl.Sing(s.privateKey)
	if err != nil {
		return nil, err
	}

	lockSigned, err := proto.Marshal(lockEl)
	if err != nil {
		return nil, err
	}

	lockHash := sha256.Sum256(lockSigned)

	return lockHash[:], nil
}
