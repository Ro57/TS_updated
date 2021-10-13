package types

import (
	ed "crypto/ed25519"
)

type Key = ed.PrivateKey

type Address interface {
	CheckSig(address string, signature []byte, data []byte) bool
	GenerateKey(randomSeed []byte) Key
	Sign(k Key, data []byte) []byte
}
