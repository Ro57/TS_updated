package lock

import (
	"crypto/sha256"
	"github.com/golang/protobuf/proto"
)

func (m Lock) GetHash() ([]byte, error) {
	bytes, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	res := sha256.Sum256(bytes)
	return res[:], nil
}
