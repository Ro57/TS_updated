package utils

import ed "crypto/ed25519"

func (p PrivateKey) Public() string {
	publicKey := make([]byte, ed.PublicKeySize)
	copy(publicKey, p.key[32:])
	return string(publicKey)
}
