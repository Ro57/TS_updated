package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"token-strike/internal/errors"
	"token-strike/internal/types"
	"token-strike/tsp2p/server/lock"
)

// --------------- Block -------------------

// Sing signed block received private key and stored signature
func (m *Block) Sing(key types.PrivateKey) error {
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

func (m *State) TransferTokens(sender, recipient string) error {
	transferLock := m.GetLock(sender, recipient)

	if transferLock == nil {

		sendOwner := m.GetOwner(sender)
		if sendOwner == nil {
			return errors.OwnerNoFoundErr
		}

		recipOwner := m.GetOwner(recipient)
		if recipOwner == nil {
			recipOwner = &Owner{
				HolderWallet: recipient,
				Count:        0,
			}
		}

		sendOwner.Count = sendOwner.Count - transferLock.Count
		recipOwner.Count = recipOwner.Count + transferLock.Count

		return m.RemoveLock(transferLock)
	}

	return errors.LockNotFoundErr
}

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

func (m *State) RemoveLock(incomingLock *lock.Lock) error {
	for index, lock := range m.Locks {
		if lock == incomingLock {
			m.Locks = append(m.Locks[:index], m.Locks[index+1:]...)
			return nil
		}
	}
	return errors.LockNotFoundErr
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
