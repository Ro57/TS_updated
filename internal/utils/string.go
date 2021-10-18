package utils

import "encoding/hex"

func (a Address) String() string {
	return hex.EncodeToString(a.publicKey)
}
