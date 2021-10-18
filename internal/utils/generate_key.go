package utils

import (
	ed "crypto/ed25519"

	"token-strike/internal/types"
)

func (e *AddressScheme) GenerateKey(randomSeed [32]byte) types.PrivateKey {
	key := ed.NewKeyFromSeed(randomSeed[:])
	return PrivateKey{key}
}
