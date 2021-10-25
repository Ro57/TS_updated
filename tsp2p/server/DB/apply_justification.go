package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	lock "token-strike/tsp2p/server/lock"

	"github.com/golang/protobuf/proto"
)

func (s *State) ApplyJustification(justification isJustification_Content) error {
	switch j := justification.(type) {
	case *Justification_Lock:
		return s.lockToken(j)
	case *Justification_Transfer:
		return s.transferToken(j)
	default:
		return nil
	}
}

func (s *State) lockToken(justification *Justification_Lock) error {
	s.Locks = append(s.Locks, justification.Lock.Lock)

	OwnerIndex := getOwnerIndexByHolder(justification.Lock.Lock.Sender, s.Owners)
	if OwnerIndex == nil {
		return fmt.Errorf("holder with name %v not found in state", justification.Lock.Lock.Sender)
	}

	s.Owners[*OwnerIndex].Count = s.Owners[*OwnerIndex].Count - justification.Lock.Lock.Count
	return nil
}

func (s *State) transferToken(justification *Justification_Transfer) error {
	lockHashIndex := getLockIndex(justification.Transfer.Lock, s.Locks)

	if lockHashIndex == nil {
		return fmt.Errorf("not found lock %v in state", justification.Transfer.Lock)
	}

	s.Locks = append(s.Locks[:*lockHashIndex], s.Locks[*lockHashIndex+1:]...)

	return nil
}

// Get index of owners array by holder wallet. Returns nil if no holder is found
func getOwnerIndexByHolder(holder string, Owners []*Owner) *int {
	for index, owner := range Owners {
		if owner.HolderWallet == holder {
			return &index
		}
	}
	return nil
}

// Get index of owners array by holder wallet. Returns nil if no holder is found
func getLockIndex(lockHash string, locks []*lock.Lock) *int {
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
