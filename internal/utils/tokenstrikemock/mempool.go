package tokenstrikemock

import (
	"encoding/base64"
	"errors"
	"fmt"
	"google.golang.org/protobuf/runtime/protoiface"
)

type Mempool interface {
	AddPeer(url string) error
	List() map[string]*MempoolEntry
	Remove(id string) bool
	Insert(hash string, msg protoiface.MessageV1, expiration int64) string
}

type MempoolEntry struct {
	Hash       string
	Expiration int64
	Message    protoiface.MessageV1
}

// implementation

var _ Mempool = &TokenStrikeMock{}

func (t *TokenStrikeMock) List() map[string]*MempoolEntry {
	return t.mempoolEntries
}

func (t *TokenStrikeMock) Remove(id string) bool {
	_, ok := t.mempoolEntries[id]
	if ok {
		delete(t.mempoolEntries, id)
		return true
	}
	return false
}

func (t *TokenStrikeMock) Insert(hash string, message protoiface.MessageV1, expiration int64) string {
	key := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s/\\%s/\\%d", hash, message, expiration)))
	t.mempoolEntries[key] = &MempoolEntry{
		Hash:       hash,
		Expiration: expiration,
		Message:    message,
	}

	go t.sendingMessages(key)

	return key
}

func (t *TokenStrikeMock) AddPeer(url string) error {
	if url != "" {
		t.peers = append(t.peers, url)
		return nil
	}
	return errors.New("url cannot is empty")
}
