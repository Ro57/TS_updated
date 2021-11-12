package tokenstrikemock

import (
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"

	"google.golang.org/protobuf/runtime/protoiface"
)

type MempoolEntry struct {
	ParentHash string
	Expiration int64
	Type       uint32
	Message    protoiface.MessageV1
}

func (m *MempoolEntry) GetDataMsg() *tokenstrike.Data {
	switch m.Type {
	case tokenstrike.TYPE_LOCK:
		return &tokenstrike.Data{
			Data:  &tokenstrike.Data_Lock{Lock: m.Message.(*lock.Lock)},
			Token: m.ParentHash,
		}
	case tokenstrike.TYPE_BLOCK:
		return &tokenstrike.Data{
			Data:  &tokenstrike.Data_Block{Block: m.Message.(*DB.Block)},
			Token: m.ParentHash,
		}
	case tokenstrike.TYPE_TX:
		return &tokenstrike.Data{
			Data:  &tokenstrike.Data_Transfer{Transfer: m.Message.(*tokenstrike.TransferTokens)},
			Token: m.ParentHash,
		}
	case tokenstrike.TYPE_ISSUE:
		msg := m.Message.(*DB.Block)
		return &tokenstrike.Data{
			Data:  &tokenstrike.Data_Block{Block: msg},
			Token: msg.GetSignature(),
		}
	}
	return nil
}
