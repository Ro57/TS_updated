package tokenstrikemock

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoiface"
)

type Mempool interface {
	AddPeer(url string) error
	List() map[string]*MempoolEntry
	Remove(id string) bool
	Insert(hash string, messageType uint32, msg protoiface.MessageV1, expiration int64) string
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

func (t *TokenStrikeMock) Insert(hash string, messageType uint32, message protoiface.MessageV1, expiration int64) string {
	key := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s/\\%s/\\%d", hash, message, expiration)))
	t.mempoolEntries[key] = &MempoolEntry{
		Type:       messageType,
		Expiration: expiration,
		Hash:       hash,
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

func (t *TokenStrikeMock) sendingMessages(hash string) {
	var genError error

	for index, peer := range t.peers {
		conn, err := grpc.DialContext(
			context.TODO(),
			peer,
			grpc.WithInsecure(),
		)
		if err != nil {
			genError = fmt.Errorf("%v : %s /n %s", index, err, genError)
		}

		client := rpcservice.NewRPCServiceClient(conn)

		blockHash, err := proto.Marshal(t.mempoolEntries[hash].Message)
		if err != nil {
			genError = fmt.Errorf("%v : %s /n %s", index, err, genError)
		}

		resp, err := client.Inv(
			context.Background(),
			&tokenstrike.InvReq{Invs: []*tokenstrike.Inv{
				{
					Parent:     []byte(t.mempoolEntries[hash].Hash),
					Type:       t.mempoolEntries[hash].Type,
					EntityHash: blockHash[:],
				},
			}},
		)
		if err != nil {
			genError = fmt.Errorf("%v : %s /n %s", index, err, genError)
		}

		if resp.Needed != nil {
			for _, need := range resp.Needed {
				if need {
					//send selected lock and NOW skip check of warning
					_, err := client.PostData(context.TODO(), t.mempoolEntries[hash].GetDataMsg())
					if err != nil {
						genError = fmt.Errorf("%v : %s /n %s", index, err, genError)
					}
				}
			}
		}
	}
	fmt.Println(genError)
}
