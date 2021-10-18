package types

import (
	ed "crypto/ed25519"
)

type Key = ed.PrivateKey

type AddressScheme interface {
	GenerateKey(seed [32]byte) PrivateKey
	ParseAddr(ps string) (Address, error)
}

type PrivateKey interface {
	Public() string
	Sign(data []byte) []byte
}

type Address interface {
	CheckSig(data []byte, signature []byte) bool
	String() string
}
