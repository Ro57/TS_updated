package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"token-strike/internal/types"
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

func (m State) GetHash() ([]byte, error) {
	stateBytes, err := proto.Marshal(&m)
	if err != nil {
		return nil, err
	}
	res := sha256.Sum256(stateBytes)
	return res[:], nil
}
