package privkey

import (
	"bytes"
	ed "crypto/ed25519"
	"token-strike/internal/types/address"
)

type SimplePrivateKey struct {
	Key ed.PrivateKey
}

var _ address.PrivateKey = &SimplePrivateKey{}

func (p SimplePrivateKey) GetPublicKey() []byte {
	publicKey := make([]byte, ed.PublicKeySize)
	copy(publicKey, p.Key[32:])
	return publicKey
}

func (p SimplePrivateKey) Equal(private address.PrivateKey) bool {
	simple, ok := private.(SimplePrivateKey)
	if !ok {
		return false
	}
	return bytes.Equal(simple.Key, p.Key)
}

func (p SimplePrivateKey) Sign(data []byte) []byte {
	return ed.Sign(p.Key, data)
}
