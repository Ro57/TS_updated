package utils

import (
	ed "crypto/ed25519"
	"errors"
	"fmt"
	"token-strike/internal/types"
)

func (a AddressScheme) ParseAddr(ps string) (types.Address, error) {
	bytePS := []byte(ps)
	if len(bytePS) != ed.PublicKeySize {
		return nil, errors.New(fmt.Sprintf("parse public address: bad length: %v", len(bytePS)))
	}
	address := Address{publicKey: bytePS}
	return address, nil
}
