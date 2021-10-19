package utils

import (
	ed "crypto/ed25519"
)

//TODO: it will not work with public interface or we should verify it by our hands
func (e SimpleAddress) CheckSig(data []byte, signature []byte) bool {
	return ed.Verify(e.publicKey, data, signature)
}
