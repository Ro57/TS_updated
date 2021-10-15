package utils

import "token-strike/internal/types"

func (p *PktChain) VerifyProof(annProof types.AnnProof) int32 {
	return annProof.Num
}
