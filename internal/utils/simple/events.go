package simple

import (
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
)

type lockEvent struct {
	content lock.Lock
}

type txEvent struct {
	content justifications.TranferToken
}

type blockEvent struct {
	content DB.Block
}

// TODO: implement interface
type issueEvent struct {
	pubKey  string
	content DB.Block
}

func (s *blockEvent) GetContent() DB.Block {
	return s.content
}

func (s *lockEvent) GetContent() lock.Lock {
	return s.content
}

func (s *txEvent) GetContent() justifications.TranferToken {
	return s.content
}
