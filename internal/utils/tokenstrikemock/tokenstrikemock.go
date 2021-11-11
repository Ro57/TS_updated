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

	NumberSecondsWaitTime = 60
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
	mempoolEntries  map[string]*MempoolEntry
	peers           []string
}

//var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}

func New(db database.DBRepository, simpleAddress address.Address) (res *TokenStrikeMock) {
	defer func() {
		go res.timerSendingMessages()
	}()

	return &TokenStrikeMock{
		bboltDB:         db,
		pktChain:        &pktchain.SimplePktChain{},
		addressScheme:   addressScheme.SimpleAddressScheme{},
		invCache:        make(map[string]tokenstrike.Inv),
		address:         simpleAddress,
		mempoolEntries:  make(map[string]*MempoolEntry, 0),
		peers:           make([]string, 0),
		blockDispatcher: make(chan *tokenstrike.Data_Block),
		lockDispatchers: []chan *LockForBlock{},
		txDispatchers:   []chan *TxForBlock{},
	}
}
