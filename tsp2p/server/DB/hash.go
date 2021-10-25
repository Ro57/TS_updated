package DB

import (
	"crypto/sha256"
	"github.com/golang/protobuf/proto"
)

func (m Block) GetHash() ([]byte, error) {
	data, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	res := sha256.Sum256(data)
	return res[:], nil
}

func (m State) GetHash() ([]byte, error) {
	stateBytes, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	res := sha256.Sum256(stateBytes)
	return res[:], nil
}
