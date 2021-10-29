package address

import (
	ed "crypto/ed25519"
	"encoding/hex"
	"token-strike/internal/types/address"
)

type SimpleAddress struct {
	publicKey ed.PublicKey
}

var _ address.Address = &SimpleAddress{}

func NewSimpleAddress(pb ed.PublicKey) SimpleAddress {
	return SimpleAddress{publicKey: pb}
}

//TODO: it will not work with public interface or we should verify it by our hands
func (e SimpleAddress) CheckSig(data []byte, signature []byte) bool {
	return ed.Verify(e.publicKey, data, signature)
}

func (e SimpleAddress) String() string {
	return hex.EncodeToString(e.publicKey)
}
