package types

type OwnerCollection map[string]uint64

type Wallet interface {
	DiscoverToken(tokenID string, issuerUrlHints []string)
	LockTokens(tokenId string, amount uint64, recipient string, secretHash string) (string, error)
	SendTokens(tokenId string, lockId string, secretHex string) (string, error)
}

type Issuer interface {
	issueToken(owners OwnerCollection, expiration int32) string
}
