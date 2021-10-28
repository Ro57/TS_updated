package types

import "token-strike/tsp2p/server/DB"

type OwnerCollection map[string]uint64

type Wallet interface {
	DiscoverToken(tokenID string, issuerUrlHints []string)
	LockTokens(tokenId string, amount uint64, recipient string, secretHash string) (string, error)
	SendTokens(tokenId string, lockId string, secretHex string) (string, error)
}

type Issuer interface {
	IssueToken(owners []*DB.Owner, expiration int32) (string, error)
}
