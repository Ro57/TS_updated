package utils

import (
	ed "crypto/ed25519"
	"token-strike/internal/types"
)

type PrivateKey struct {
	key ed.PrivateKey
}

var _ types.PrivateKey = &PrivateKey{}
