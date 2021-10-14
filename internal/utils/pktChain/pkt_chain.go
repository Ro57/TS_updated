package utils

import "token-strike/internal/types"

type pktChain struct {
}

var _ types.PktChain = (*pktChain)(nil)
