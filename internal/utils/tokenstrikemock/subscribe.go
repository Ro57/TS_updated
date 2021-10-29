package tokenstrikemock

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"token-strike/tsp2p/server/tokenstrike"
)

func (t *TokenStrikeMock) Subscribe(empty *empty.Empty, server tokenstrike.TokenStrike_SubscribeServer) error {
	if t.peers == nil {
		t.peers = []tokenstrike.TokenStrike_SubscribeServer{}
	}
	t.peers = append(t.peers, server)
	return nil
}

func (t *TokenStrikeMock) sendDataToSubscribers(msg *tokenstrike.Data) error {
	var genError error

	for i, peer := range t.peers {
		err := peer.Send(msg)
		if err != nil {
			genError = fmt.Errorf("%v : %s /n %s", i, err, genError)
		}
	}

	return genError
}
