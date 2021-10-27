package types

type OwnerCollection map[string]uint64

type Wallet interface {
	DiscoverToken(tokenID string, issuerUrlHints []string)
}

type Issuer interface {
	MintToken(owners OwnerCollection, expiration int32) string
}
