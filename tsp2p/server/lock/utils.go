package lock

import (
	"crypto/sha256"
	"encoding/hex"
	"token-strike/internal/types/address"

	"github.com/golang/protobuf/proto"
)

func (m *Lock) Sing(key address.PrivateKey) error {
	bytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	m.Signature = hex.EncodeToString(key.Sign(bytes))
	return nil
}

func (m Lock) GetHash() ([]byte, error) {
	bytes, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	res := sha256.Sum256(bytes)
	return res[:], nil
}
