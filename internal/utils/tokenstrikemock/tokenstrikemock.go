package tokenstrikemock

import (
	"token-strike/internal/database"
	"token-strike/internal/types"
	"token-strike/internal/utils"
	"token-strike/tsp2p/server/tokenstrike"
)

type TokenStrikeMock struct {
	bboltDB       database.DBRepository
	issuer        types.Address
	pktChain      utils.SimplePktChain
	addressScheme utils.SimpleAddressScheme
	invCache      map[string]tokenstrike.Inv
}

var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}

func New(db database.DBRepository, issuer types.Address) *TokenStrikeMock {
	return &TokenStrikeMock{
		bboltDB:       db,
		issuer:        issuer,
		pktChain:      utils.SimplePktChain{},
		addressScheme: utils.SimpleAddressScheme{},
		invCache:      make(map[string]tokenstrike.Inv),
	}
}
