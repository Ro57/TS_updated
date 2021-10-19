package utils

import (
	ed "crypto/ed25519"
	"token-strike/internal/types"
)

type SimplePrivateKey struct {
	key ed.PrivateKey
}

var _ types.PrivateKey = &SimplePrivateKey{}
