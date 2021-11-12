package tokenstrikemock

import (
	"encoding/hex"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"
)

type LockEvent struct {
	TokenID string
	Content lock.Lock
}

type TxEvent struct {
	TokenID string
	Content justifications.TranferToken
}

type BlockEvent struct {
	TokenID string
	Content DB.Block
}

type Dispatcher struct {
	Lock  chan *LockEvent
	TX    chan *TxEvent
	Block chan *BlockEvent
}

func (t *TokenStrikeMock) Subscribe(PartenHash string) Dispatcher {
	dispatcher := &Dispatcher{
		Lock:  make(chan *LockEvent),
		TX:    make(chan *TxEvent),
		Block: make(chan *BlockEvent),
	}

	t.dispatchers[PartenHash] = append(t.dispatchers[PartenHash], dispatcher)

	return *dispatcher
}

func (t *TokenStrikeMock) dispatch(msg *tokenstrike.Data) {
	go func() {
		switch data := msg.Data.(type) {
		case *tokenstrike.Data_Block:
			t.dispatchBlock(msg.Token, data)
		case *tokenstrike.Data_Lock:
			t.dispatchLock(msg.Token, data)
		case *tokenstrike.Data_Transfer:
			t.dispatchTx(msg.Token, data)

		}
	}()
}

func (t *TokenStrikeMock) dispatchBlock(ParentHash string, data *tokenstrike.Data_Block) {
	for _, dispathcer := range t.dispatchers[ParentHash] {
		dispathcer.Block <- &BlockEvent{
			TokenID: ParentHash,
			Content: *data.Block,
		}
	}
}

func (t *TokenStrikeMock) dispatchLock(ParentHash string, data *tokenstrike.Data_Lock) {
	for _, dispathcer := range t.dispatchers[ParentHash] {
		dispathcer.Lock <- &LockEvent{
			TokenID: ParentHash,
			Content: *data.Lock,
		}
	}
}

func (t *TokenStrikeMock) dispatchTx(ParentHash string, data *tokenstrike.Data_Transfer) {
	for _, dispathcer := range t.dispatchers[ParentHash] {
		dispathcer.TX <- &TxEvent{
			TokenID: ParentHash,
			Content: justifications.TranferToken{
				HtlcSecret: hex.EncodeToString(data.Transfer.GetHtlc()),
				Lock:       data.Transfer.GetLockId(),
			},
		}
	}
}
