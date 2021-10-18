package utils

import (
	"token-strike/internal/types"
)

type AddressScheme struct{}

var _ types.AddressScheme = &AddressScheme{}
