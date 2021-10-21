package tokenstrikemock

import (
	ed "crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"token-strike/internal/types"
)

// SimpleAddressMock ...
type SimpleAddressMock struct {
	publicKey ed.PublicKey
}

var _ types.Address = &SimpleAddressMock{}

func (s SimpleAddressMock) CheckSig(data []byte, signature []byte) bool {
	return true
}

func (s SimpleAddressMock) String() string {
	return hex.EncodeToString(s.publicKey)
}

// SimpleAddressSchemeMock ...
type SimpleAddressSchemeMock struct{}

var _ types.AddressScheme = &SimpleAddressSchemeMock{}

func (s SimpleAddressSchemeMock) GenerateKey(randomSeed [32]byte) types.PrivateKey {
	seed := []byte("1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa")
	key := ed.NewKeyFromSeed(seed[:32])
	return SimplePrivateKeyMock{key}
}

func (s SimpleAddressSchemeMock) ParseAddr(ps string) (types.Address, error) {
	bytePS := []byte(ps)
	if len(bytePS) != ed.PublicKeySize {
		return nil, errors.New(fmt.Sprintf("parse public address: bad length: %v", len(bytePS)))
	}
	address := SimpleAddressMock{publicKey: bytePS}
	return address, nil
}

// SimplePrivateKeyMock ...
type SimplePrivateKeyMock struct {
	key ed.PrivateKey
}

var _ types.PrivateKey = &SimplePrivateKeyMock{}

// TODO: rework implementations
func (s SimplePrivateKeyMock) Address() types.Address {
	publicKey := make([]byte, ed.PublicKeySize)
	copy(publicKey, s.key[32:])
	return nil
}

func (s SimplePrivateKeyMock) Sign(data []byte) []byte {
	panic("implement me")
}
