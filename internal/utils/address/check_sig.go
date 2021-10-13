package utils

import (
	ed "crypto/ed25519"
)

func (e *Address) CheckSig(address string, signature []byte, data []byte) bool {
	public := ed.PublicKey(address)
	return ed.Verify(public, data, signature)
}
