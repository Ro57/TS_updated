package utils

import (
	ed "crypto/ed25519"
	"token-strike/internal/types"
)

type Address struct {
	publicKey ed.PublicKey
}

var _ types.Address = &Address{}
