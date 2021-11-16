package simple

import (
	"encoding/hex"
	"token-strike/internal/types/dispatcher"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"
)

type SimpleTokenDispatcher struct {
	lock              chan dispatcher.LockEvent
	block             chan dispatcher.BlockEvent
	tx                chan dispatcher.TxEvent
	lockActionsSlice  []func(lock.Lock)
	blockActionsSlice []func(DB.Block)
	txActionsSlice    []func(justifications.TranferToken)
}

func NewTokenDispatcher() dispatcher.TokenDispatcher {
	return &SimpleTokenDispatcher{
		lock:              make(chan dispatcher.LockEvent),
		block:             make(chan dispatcher.BlockEvent),
		tx:                make(chan dispatcher.TxEvent),
		lockActionsSlice:  []func(lock.Lock){},
		blockActionsSlice: []func(DB.Block){},
		txActionsSlice:    []func(justifications.TranferToken){},
	}
}

func (s *SimpleTokenDispatcher) Observe() error {
	go func() {
		for {
			select {
			case lock := <-s.lock:
				s.executeLockActions(lock.GetContent())
			case block := <-s.block:
				s.executeBlockActions(block.GetContent())
			case tx := <-s.tx:
				s.executeTXActions(tx.GetContent())
			}
		}
	}()
	return nil
}

func (s *SimpleTokenDispatcher) SendLock(eventLock lock.Lock) {
	s.lock <- &lockEvent{
		content: eventLock,
	}
}

func (s *SimpleTokenDispatcher) SendBlock(block DB.Block) {
	s.block <- &blockEvent{
		content: block,
	}
}

func (s *SimpleTokenDispatcher) SendTx(tx tokenstrike.TransferTokens) {
	s.tx <- &txEvent{
		content: justifications.TranferToken{
			HtlcSecret: hex.EncodeToString(tx.GetHtlc()),
			Lock:       tx.GetLockId(),
		},
	}
}

func (s *SimpleTokenDispatcher) WaitLockAction(action dispatcher.LockAction) chan dispatcher.IdResult {
	wait := make(chan dispatcher.IdResult)
	s.lockActionsSlice = append(s.lockActionsSlice, func(eventLock lock.Lock) {
		id, err := action(eventLock)
		wait <- dispatcher.IdResult{
			ID:  id,
			Err: err,
		}
	})

	return wait
}

func (s *SimpleTokenDispatcher) WaitBlockAction(action dispatcher.BlockAction) chan dispatcher.IdResult {
	wait := make(chan dispatcher.IdResult)
	s.blockActionsSlice = append(s.blockActionsSlice, func(b DB.Block) {
		id, err := action(b)
		wait <- dispatcher.IdResult{
			ID:  id,
			Err: err,
		}
	})

	return wait
}

func (s *SimpleTokenDispatcher) WaitTxAction(action dispatcher.TXAction) chan dispatcher.IdResult {
	wait := make(chan dispatcher.IdResult)
	s.txActionsSlice = append(s.txActionsSlice, func(tx justifications.TranferToken) {
		id, err := action(tx)
		wait <- dispatcher.IdResult{
			ID:  id,
			Err: err,
		}
	})

	return wait
}

func (s *SimpleTokenDispatcher) executeLockActions(eventLock lock.Lock) {
	for _, waiter := range s.lockActionsSlice {
		waiter(eventLock)
	}

	s.lockActionsSlice = []func(lock.Lock){}
}

func (s *SimpleTokenDispatcher) executeBlockActions(block DB.Block) {
	for _, waiter := range s.blockActionsSlice {
		waiter(block)
	}

	s.blockActionsSlice = []func(DB.Block){}
}

func (s *SimpleTokenDispatcher) executeTXActions(tx justifications.TranferToken) {
	for _, waiter := range s.txActionsSlice {
		waiter(tx)
	}

	s.txActionsSlice = []func(justifications.TranferToken){}
}
