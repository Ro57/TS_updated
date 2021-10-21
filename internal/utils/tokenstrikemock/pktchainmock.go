package tokenstrikemock

import (
	"crypto/sha256"
	"token-strike/internal/types"
)

type SimplePktChainMock struct {
}

var _ types.PktChain = &SimplePktChainMock{}

func (s SimplePktChainMock) CurrentHeight() int32 {
	return 10
}

func (s SimplePktChainMock) BlockHashAtHeight(i int32) []byte {
	var result []byte
	if i < s.CurrentHeight() {
		sha := sha256.Sum256([]byte(string(i)))
		result = sha[:]
	}
	return result
}

func (s SimplePktChainMock) AnnounceData(bytes []byte) chan types.AnnProof {
	panic("implement me")
}

func (s SimplePktChainMock) VerifyProof(proof types.AnnProof) int32 {
	panic("implement me")
}


