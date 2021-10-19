package utils

import (
	"token-strike/internal/types"
)

type SimpleAddressScheme struct{}

var _ types.AddressScheme = &SimpleAddressScheme{}
