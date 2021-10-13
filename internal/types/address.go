package types

import (
	ed "crypto/ed25519"
)

<<<<<<< HEAD:internal/types/address.go
type Key ed.PrivateKey
=======
type Key = ed.PrivateKey
>>>>>>> 7fa4619 (feat(sprint1): add utils package):internal/utils/address.go

type Address interface {
	CheckSig(address string, signature []byte, data []byte) bool
	GenerateKey(randomSeed []byte) Key
	Sign(k Key, data []byte) []byte
}
