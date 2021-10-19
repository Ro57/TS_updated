package utils

import (
	"crypto/ed25519"
)

func (p SimplePrivateKey) Sign(data []byte) []byte {
	return ed25519.Sign(p.key, data)
}
