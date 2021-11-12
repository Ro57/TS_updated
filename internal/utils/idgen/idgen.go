package idgen

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

const separator = ":"

func EntityIndex(Block DB.Block, hashOfEntity []byte) (int, error) {
	for i, justification := range Block.Justifications {
		if isEntity(justification, hashOfEntity) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("Justification with %v content hash, not found", hashOfEntity)
}

func Encode(blockHash string, number int) string {
	return fmt.Sprint(blockHash, separator, number)
}

func Decode(ID string) (*string, *int, error) {
	tuple := strings.Split(ID, separator)

	blockHash := tuple[0]

	justificationNumber, err := strconv.Atoi(tuple[1])
	if err != nil {
		return nil, nil, err
	}

	return &blockHash, &justificationNumber, nil
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

func isTransferEqual(justificationTX *justifications.TranferToken, entity []byte) bool {
	htlc, err := hex.DecodeString(justificationTX.HtlcSecret)
	if err != nil {
		return false
	}

	tx := &tokenstrike.TransferTokens{
		Htlc:   htlc,
		LockId: justificationTX.Lock,
	}

	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return false
	}

	hash := sha256.Sum256(txBytes)

	return bytes.Equal(hash[:], entity)
}
