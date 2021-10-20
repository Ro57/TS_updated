package utils

import (
	ed "crypto/ed25519"

	"token-strike/internal/types"
)

func (e *SimpleAddressScheme) GenerateKey(randomSeed [32]byte) types.PrivateKey {
	key := ed.NewKeyFromSeed(randomSeed[:])
	return SimplePrivateKey{Key: key}
}
