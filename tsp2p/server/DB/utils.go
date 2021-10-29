package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"token-strike/internal/types/address"

	"token-strike/tsp2p/server/lock"

	"github.com/golang/protobuf/proto"
)

// --------------- Block -------------------

// Sing signed block received private key and stored signature
func (m *Block) Sing(key address.PrivateKey) error {
	bytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	m.Signature = hex.EncodeToString(key.Sign(bytes))
	return nil
}

func (m Block) GetHash() ([]byte, error) {
	data, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	res := sha256.Sum256(data)
	return res[:], nil
}

// --------------- State -------------------

func (m State) GetHash() ([]byte, error) {
	stateBytes, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	res := sha256.Sum256(stateBytes)
	return res[:], nil
}

func (m State) OwnerExist(addr string) bool {
	for _, owner := range m.Owners {
		if owner.HolderWallet == addr {
			return true
		}
	}
	return false
}

func (m State) GetOwner(addr string) *Owner {
	for _, owner := range m.Owners {
		if owner.HolderWallet == addr {
			return owner
		}
	}
	return nil
}

func (m State) GetLock(sender, recipient string) *lock.Lock {
	for _, lock := range m.Locks {
		if lock.Sender == sender &&
			lock.Recipient == recipient {
			return lock
		}
	}
	return nil
}

func (m State) GetLockIndexByHash(lockHash string, locks []*lock.Lock) *int {
	for index, lock := range locks {
		lockByte, err := proto.Marshal(lock)
		if err != nil {
			return nil
		}

		curLockHash := sha256.Sum256(lockByte)

		if hex.EncodeToString(curLockHash[:]) == lockHash {
			return &index
		}

	}
	return nil
}

func (m State) GetOwnerIndexByHolder(holder string, Owners []*Owner) *int {
	for index, owner := range Owners {
		if owner.HolderWallet == holder {
			return &index
		}
	}
	return nil
}
