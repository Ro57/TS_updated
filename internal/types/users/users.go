package users

import "token-strike/tsp2p/server/DB"

type LockArgs interface {
	GetTokenId() string
	GetAmount() uint64
	GetRecipient() string
	GetSecretHash() string
}

type Wallet interface {
	DiscoverToken(tokenID string) error
	LockTokens(args LockArgs) ([]byte, error)
	SendTokens(tokenId string, lockId []byte, secretHex []byte) ([]byte, error)
}

type Issuer interface {
	IssueToken(owners []*DB.Owner, expiration int32) (string, error)
}
