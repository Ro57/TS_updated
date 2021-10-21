package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
)

func (m Block) GetBlockHash() []byte {
	data, _ := proto.Marshal(&m)
	res := sha256.Sum256(data)
	return res[:]
}

func (m State) GetStateHash() string {
	stateBytes, _ := proto.Marshal(&m)
	return hex.EncodeToString(stateBytes)
}
