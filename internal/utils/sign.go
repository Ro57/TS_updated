package utils

import (
	"crypto/ed25519"
	"token-strike/internal/types"
)

func (a *Address) Sign(key types.Key, data []byte) []byte {
	return ed25519.Sign(key, data)
}
