package DB

import (
	"crypto/sha256"
	"github.com/golang/protobuf/proto"
)

func (m Block) GetBlockHash() []byte {
	data, _ := proto.Marshal(&m)
	res := sha256.Sum256(data)
	return res[:]
}
