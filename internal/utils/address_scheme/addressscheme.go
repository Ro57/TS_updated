package address_scheme

import (
	ed "crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"token-strike/internal/types/address"
	simpleAddress "token-strike/internal/utils/address"

	"token-strike/internal/utils/privkey"
)

type SimpleAddressScheme struct{}

var _ address.AddressScheme = &SimpleAddressScheme{}

func (s SimpleAddressScheme) ParseAddr(ps string) (address.Address, error) {

	bytePS, err := hex.DecodeString(ps)
	if err != nil {
		return nil, err
	}

	if len(bytePS) != ed.PublicKeySize {
		return nil, errors.New(fmt.Sprintf("parse public address: bad length: %v", len(bytePS)))
	}

	address := simpleAddress.NewSimpleAddress(bytePS)

	return address, nil

}
func (s *SimpleAddressScheme) GenerateKey(randomSeed [32]byte) address.PrivateKey {
	key := ed.NewKeyFromSeed(randomSeed[:])
	return privkey.SimplePrivateKey{Key: key}
}
