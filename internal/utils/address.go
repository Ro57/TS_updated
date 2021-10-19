package utils

import (
	ed "crypto/ed25519"
	"token-strike/internal/types"
)

type SimpleAddress struct {
	publicKey ed.PublicKey
}

var _ types.Address = &SimpleAddress{}
