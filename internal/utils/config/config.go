package config

import (
	"token-strike/internal/database/repository"
	"token-strike/internal/types"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/pktchain"
)

type Config struct {
	Scheme types.AddressScheme
	Chain  types.PktChain
	DB     repository.Bbolt
}

func DefaultSimpleConfig() Config {
	return Config{
		Scheme: &address.SimpleAddressScheme{},
		Chain:  &pktchain.SimplePktChain{},
	}
}
