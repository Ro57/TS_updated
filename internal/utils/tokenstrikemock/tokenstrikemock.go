package tokenstrikemock

import (
	"token-strike/internal/database"
	address2 "token-strike/internal/types/address"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/pktchain"
	"token-strike/tsp2p/server/tokenstrike"
)

const (
	NeedData     = true
	DontNeedData = false
)

type TokenStrikeMock struct {
	bboltDB       database.DBRepository
	issuer        address2.Address
	pktChain      pktchain.SimplePktChain
	addressScheme address.SimpleAddressScheme
	invCache      map[string]tokenstrike.Inv
}

var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}

func New(db database.DBRepository, issuer address2.Address) *TokenStrikeMock {
	return &TokenStrikeMock{
		bboltDB:       db,
		issuer:        issuer,
		pktChain:      pktchain.SimplePktChain{},
		addressScheme: address.SimpleAddressScheme{},
		invCache:      make(map[string]tokenstrike.Inv),
	}
}
