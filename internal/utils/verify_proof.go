package utils

import "token-strike/internal/types"

func (p *SimplePktChain) VerifyProof(annProof types.AnnProof) int32 {
	return annProof.Num
}
