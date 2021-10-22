package utils

import (
	ed "crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"token-strike/internal/types"
)

func (a SimpleAddressScheme) ParseAddr(ps string) (types.Address, error) {
	bytePS, err := hex.DecodeString(ps)
	if err != nil {
		return nil, err
	}

	if len(bytePS) != ed.PublicKeySize {
		return nil, errors.New(fmt.Sprintf("parse public address: bad length: %v", len(bytePS)))
	}

	address := SimpleAddress{publicKey: bytePS}

	return address, nil
}
