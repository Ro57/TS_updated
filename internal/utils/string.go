package utils

import "encoding/hex"

func (a SimpleAddress) String() string {
	return hex.EncodeToString(a.publicKey)
}
