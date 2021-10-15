package utils

import "crypto/sha256"

func (p PktChain) BlockHashAtHeight(i int32) []byte {
	var result []byte
	if i < p.CurrentHeight() {
		sha := sha256.Sum256([]byte(string(i)))
		result = sha[:]
	}
	return result
}
