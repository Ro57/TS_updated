package tokenstrikemock

import (
	"token-strike/internal/types/dispatcher"
	"token-strike/internal/utils/simple"
	"token-strike/tsp2p/server/tokenstrike"
)

func (t *TokenStrikeMock) Subscribe(PartenHash string) dispatcher.Dispatcher {

	// if dispather not exist, create it
	if t.dispatchers[PartenHash] == nil {
		t.dispatchers[PartenHash] = simple.NewTokenDispatcher()
	}

	return t.dispatchers[PartenHash]
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
	if parentDispatcher, ok := t.dispatchers[ParentHash]; ok {
		parentDispatcher.SendBlock(*data.Block)
	}

}

func (t *TokenStrikeMock) dispatchLock(ParentHash string, data *tokenstrike.Data_Lock) {
	if parentDispatcher, ok := t.dispatchers[ParentHash]; ok {
		parentDispatcher.SendLock(*data.Lock)
	}
}

func (t *TokenStrikeMock) dispatchTx(ParentHash string, data *tokenstrike.Data_Transfer) {
	if parentDispatcher, ok := t.dispatchers[ParentHash]; ok {
		parentDispatcher.SendTx(*data.Transfer)
	}
}
