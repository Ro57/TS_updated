package tokenstrikemock

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type Mempool interface {
	AddPeer(url string) error
	List() map[string]*MempoolEntry
	Remove(id string) bool
	Insert(entry MempoolEntry) string
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

func (t *TokenStrikeMock) Insert(entry MempoolEntry) string {
	t.mempoolEntries[entry.ParentHash] = &entry

	go t.sendingMessages(entry.ParentHash)

	return entry.ParentHash
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
		err := t.sendMessageToPeer(hash, peer, index)
		genError = fmt.Errorf("%v : %s /n %s", index, err, genError)
	}

	if genError != nil {
		fmt.Println(genError)
	}
}

func (t *TokenStrikeMock) sendMessageToPeer(hash string, peer string, index int) error {
	var genError error

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
				Parent:     []byte(t.mempoolEntries[hash].ParentHash),
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
	return genError
}

func (t *TokenStrikeMock) timerSendingMessages() {
	for {
		var (
			countPeers = len(t.peers) - 1
		)

		if len(t.mempoolEntries) > 0 && countPeers > 0 {
			rand.Seed(time.Now().UnixNano())

			hash := t.randomHash()
			if hash == "" {
				fmt.Printf("mempool %v is empty?", t.mempoolEntries)
			}

			err := t.sendMessageToPeer(hash, t.peers[rand.Intn(countPeers)], 1)
			if err != nil {
				fmt.Println(err)
			}
		}

		time.Sleep(time.Second * NumberSecondsWaitTime)
	}
}

func (t *TokenStrikeMock) randomHash() string {
	for k := range t.mempoolEntries {
		return k
	}

	return ""
}
