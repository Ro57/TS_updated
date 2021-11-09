package tokenstrikemock

import (
	"token-strike/internal/database"
	"token-strike/internal/types/address"
	"token-strike/internal/types/pkt"
	"token-strike/internal/utils/pktchain"
	addressScheme "token-strike/internal/utils/simple"
	"token-strike/tsp2p/server/tokenstrike"
)

const (
	NeedData     = true
	DontNeedData = false
)

type TokenStrikeMock struct {
	address         address.Address
	bboltDB         database.DBRepository
	pktChain        pkt.PktChain
	addressScheme   addressScheme.SimpleAddressScheme
	invCache        map[string]tokenstrike.Inv
	blockDispatcher chan *tokenstrike.Data_Block
	lockDispatchers []chan *LockForBlock
	txDispatchers   []chan *TxForBlock
}

//var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}

func New(db database.DBRepository, simpleAddress address.Address) *TokenStrikeMock {
	return &TokenStrikeMock{
		bboltDB:         db,
		pktChain:        &pktchain.SimplePktChain{},
		addressScheme:   addressScheme.SimpleAddressScheme{},
		invCache:        make(map[string]tokenstrike.Inv),
		address:         simpleAddress,
		blockDispatcher: make(chan *tokenstrike.Data_Block),
		lockDispatchers: []chan *LockForBlock{},
		txDispatchers:   []chan *TxForBlock{},
	}
}
