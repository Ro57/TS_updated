package utils

import (
	"bytes"
	"token-strike/internal/types"
)

func (p SimplePrivateKey) Equal(private types.PrivateKey) bool {
	simple, ok := private.(SimplePrivateKey)
	if !ok {
		return false
	}

	return bytes.Equal(simple.Key, p.Key)
}
