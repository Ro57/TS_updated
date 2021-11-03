package simple

import (
	"bytes"
	ed "crypto/ed25519"

	"token-strike/internal/types/address"
)

func NewSimplePrivateKey(key ed.PrivateKey) SimplePrivateKey {
	return SimplePrivateKey{key: key}
}

type SimplePrivateKey struct {
	key ed.PrivateKey
}

var _ address.PrivateKey = &SimplePrivateKey{}

func (p SimplePrivateKey) GetPublicKey() []byte {
	publicKey := make([]byte, ed.PublicKeySize)
	copy(publicKey, p.key[32:])
	return publicKey
}

func (p SimplePrivateKey) Equal(private address.PrivateKey) bool {
	simple, ok := private.(SimplePrivateKey)
	if !ok {
		return false
	}
	return bytes.Equal(simple.key, p.key)
}

func (p SimplePrivateKey) Sign(data []byte) []byte {
	return ed.Sign(p.key, data)
}

func (p SimplePrivateKey) Address() address.Address {
	publicKey := make([]byte, ed.PublicKeySize)

	copy(publicKey, p.key[32:])

	return SimpleAddress{publicKey}
}
