package utils

import "token-strike/internal/types"

func (p *pktChain) VerifyProof(annProof types.AnnProof) int32 {
	return annProof.Num
}
