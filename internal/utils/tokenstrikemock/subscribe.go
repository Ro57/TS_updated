package tokenstrikemock

import (
	"encoding/hex"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"
)

type LockForBlock struct {
	TokenID string
	Content lock.Lock
}

type TxForBlock struct {
	TokenID string
	Content justifications.TranferToken
}

func (t *TokenStrikeMock) CreateNewLockChannel() chan *LockForBlock {
	newLockChannel := make(chan *LockForBlock)

	t.lockDispatchers = append(t.lockDispatchers, newLockChannel)

	return newLockChannel
}

func (t *TokenStrikeMock) CreateNewTxChannel() chan *TxForBlock {
	newTxChannel := make(chan *TxForBlock)

	t.txDispatchers = append(t.txDispatchers, newTxChannel)

	return newTxChannel
}

func (t *TokenStrikeMock) dispatch(msg *tokenstrike.Data) {
	go func() {
		switch data := msg.Data.(type) {
		case *tokenstrike.Data_Block:
			t.blockDispatcher <- data
		case *tokenstrike.Data_Lock:
			for _, lockDispatcher := range t.lockDispatchers {
				lockDispatcher <- &LockForBlock{
					Content: *data.Lock,
					TokenID: msg.Token,
				}
			}
		case *tokenstrike.Data_Transfer:
			for _, txDispatcher := range t.txDispatchers {
				txDispatcher <- &TxForBlock{
					Content: justifications.TranferToken{
						HtlcSecret: hex.EncodeToString(data.Transfer.Htlc),
						Lock:       data.Transfer.LockId,
					},
					TokenID: msg.Token,
				}
			}
		}
	}()
}
