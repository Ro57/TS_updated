package DB

import (
	"crypto/sha256"
	"github.com/golang/protobuf/proto"
)

func (m Block) GetBlockHash() ([]byte, error) {
	data, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	res := sha256.Sum256(data)
	return res[:], nil
}

func (m State) GetStateHash() ([]byte, error) {
	stateBytes, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	return stateBytes, nil
}
