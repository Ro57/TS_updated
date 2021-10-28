package config

import (
	"token-strike/internal/database"
	"token-strike/internal/types"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/pktchain"
)

type Config struct {
	Scheme types.AddressScheme
	Chain  types.PktChain
	DB     database.DBRepository
}

func DefaultSimpleConfig() Config {
	return Config{
		Scheme: &address.SimpleAddressScheme{},
		Chain:  &pktchain.SimplePktChain{},
	}
}

type LockArgs struct {
	TokenId    string
	Amount     uint64
	Recipient  string
	SecretHash string
}

func (l LockArgs) GetTokenId() string {
	return l.TokenId
}

func (l LockArgs) GetAmount() uint64 {
	return l.Amount
}

func (l LockArgs) GetRecipient() string {
	return l.Recipient
}

func (l LockArgs) GetSecretHash() string {
	return l.SecretHash
}
