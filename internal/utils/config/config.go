package config

import (
	"token-strike/internal/types/address"
	"token-strike/internal/types/pkt"
	addressScheme "token-strike/internal/utils/address_scheme"
	"token-strike/internal/utils/pktchain"
)

type Config struct {
	Scheme address.AddressScheme
	Chain  pkt.PktChain
}

func DefaultSimpleConfig() Config {
	return Config{
		Scheme: &addressScheme.SimpleAddressScheme{},
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
