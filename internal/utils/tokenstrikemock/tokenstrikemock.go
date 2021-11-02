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
	address       address.SimpleAddress
	bboltDB       database.DBRepository
	pktChain      pktchain.SimplePktChain
	addressScheme address.SimpleAddressScheme
	invCache      map[string]tokenstrike.Inv
}

var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}

func New(db database.DBRepository, simpleAddress address.SimpleAddress) *TokenStrikeMock {
	return &TokenStrikeMock{
		bboltDB:       db,
		pktChain:      pktchain.SimplePktChain{},
		addressScheme: address.SimpleAddressScheme{},
		invCache:      make(map[string]tokenstrike.Inv),
		address:       simpleAddress,
	}
}
