package tokenstrikemock

import (
	"token-strike/internal/database"
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
	pktChain      pktchain.SimplePktChain
	addressScheme address.SimpleAddressScheme
	invCache      map[string]tokenstrike.Inv
}

var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}

func New(db database.DBRepository) *TokenStrikeMock {
	return &TokenStrikeMock{
		bboltDB:       db,
		pktChain:      pktchain.SimplePktChain{},
		addressScheme: address.SimpleAddressScheme{},
		invCache:      make(map[string]tokenstrike.Inv),
	}
}
