package tokenstrikemock

import (
	"token-strike/internal/database"
	"token-strike/internal/types/address"
	"token-strike/internal/types/pkt"
	addressScheme "token-strike/internal/utils/address_scheme"
	"token-strike/internal/utils/pktchain"
	"token-strike/tsp2p/server/tokenstrike"
)

const (
	NeedData     = true
	DontNeedData = false
)

type TokenStrikeMock struct {
	address       address.Address
	bboltDB       database.DBRepository
	pktChain      pkt.PktChain
	addressScheme addressScheme.SimpleAddressScheme
	invCache      map[string]tokenstrike.Inv
}

//var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}

func New(db database.DBRepository, simpleAddress address.Address) *TokenStrikeMock {
	return &TokenStrikeMock{
		bboltDB:       db,
		pktChain:      &pktchain.SimplePktChain{},
		addressScheme: addressScheme.SimpleAddressScheme{},
		invCache:      make(map[string]tokenstrike.Inv),
		address:       simpleAddress,
	}
}
