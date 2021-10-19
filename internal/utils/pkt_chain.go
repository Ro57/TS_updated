package utils

import "token-strike/internal/types"

type SimplePktChain struct {
}

var _ types.PktChain = (*SimplePktChain)(nil)
