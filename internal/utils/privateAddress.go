package utils

import (
	ed "crypto/ed25519"
	"token-strike/internal/types"
)

func (p SimplePrivateKey) Address() types.Address {
	publicKey := make([]byte, ed.PublicKeySize)

	copy(publicKey, p.key[32:])

	return SimpleAddress{publicKey}
}
