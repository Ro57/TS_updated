package tokenstrikemock

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"token-strike/internal/utils/idgen"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"

	"github.com/golang/protobuf/proto"
)

func (t TokenStrikeMock) AwaitJustification(tokenId string, hashOfEntity []byte) (*string, error) {
	// Obtained after the method dispatch is called
	respBlock := <-t.blockDispatcher

	blockBytes, err := respBlock.Block.GetHash()
	if err != nil {
		return nil, err
	}

	blockHash := hex.EncodeToString(blockBytes)

	number, err := getJustificationIndex(respBlock.Block.Justifications, hashOfEntity)
	if err != nil {
		return nil, err
	}

	id := idgen.Encode(blockHash, number)

	return &id, nil
}

func getJustificationIndex(justifications []*DB.Justification, hashOfEntity []byte) (int, error) {
	for i, justification := range justifications {
		if isEntity(justification, hashOfEntity) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("Justification with %v content hash, not found", hashOfEntity)
}

func isEntity(justification *DB.Justification, hashOfEntity []byte) bool {
	switch data := justification.Content.(type) {
	case *DB.Justification_Lock:
		return isLockEqual(data.Lock.Lock, hashOfEntity)
	case *DB.Justification_Transfer:
		return isTransferEqual(data.Transfer, hashOfEntity)
	}

	return false
}

func isLockEqual(lock *lock.Lock, entity []byte) bool {
	hash, err := lock.GetHash()
	if err != nil {
		return false
	}

	return bytes.Equal(hash, entity)
}

func isTransferEqual(tx *justifications.TranferToken, entity []byte) bool {

	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return false
	}

	hash := sha256.Sum256(txBytes)

	return bytes.Equal(hash[:], entity)
}
