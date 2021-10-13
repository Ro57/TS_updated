package utils

import (
	ed "crypto/ed25519"

	"token-strike/internal/types"
)

func (e *Address) GenerateKey(randomSeed []byte) types.Key {
	return ed.NewKeyFromSeed(randomSeed)
}
