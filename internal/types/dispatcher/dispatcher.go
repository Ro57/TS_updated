package dispatcher

import (
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"
)

type LockAction func(lock.Lock) (string, error)
type BlockAction func(DB.Block) (string, error)
type TXAction func(justifications.TranferToken) (string, error)

type IdResult struct {
	ID  string
	Err error
}

type BlockEvent interface {
	GetContent() DB.Block
}

type LockEvent interface {
	GetContent() lock.Lock
}

type TxEvent interface {
	GetContent() justifications.TranferToken
}

type Dispatcher interface {
	Observe() error
}

type TokenDispatcher interface {
	Dispatcher
	SendLock(lock.Lock)
	SendBlock(DB.Block)
	SendTx(tokenstrike.TransferTokens)
	WaitLockAction(LockAction) chan IdResult
	WaitBlockAction([]byte, BlockAction) chan IdResult
	WaitTxAction(TXAction) chan IdResult
}

// TODO: describe interface for subscribe on issuer events
type IssueDispatcher interface {
	Dispatcher
}
