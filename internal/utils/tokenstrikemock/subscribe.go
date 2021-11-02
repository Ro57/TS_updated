package tokenstrikemock

import (
	"fmt"
	"token-strike/tsp2p/server/tokenstrike"
)

//
//import (
//	"fmt"
//	"token-strike/tsp2p/server/tokenstrike"
//
//	empty "github.com/golang/protobuf/ptypes/empty"
//)
//
//func (t *TokenStrikeMock) Subscribe(empty *empty.Empty, server tokenstrike.TokenStrike_SubscribeServer) error {
//	if t.peers == nil {
//		t.peers = []tokenstrike.TokenStrike_SubscribeServer{}
//	}
//	t.peers = append(t.peers, server)
//
//	<-server.Context().Done()
//
//	return nil
//}

func (t *TokenStrikeMock) sendDataToSubscribers(msg *tokenstrike.Data) error {
	var genError = []error{}

	//for i, peer := range t.peers {
	//	err := peer.Send(msg)
	//	if err != nil {
	//		genError = append(genError, fmt.Errorf("%v : %v /n", i, err))
	//	}
	//}

	if len(genError) > 0 {
		err := genError[0]

		for i, e := range genError[1:] {
			err = fmt.Errorf("%v/n %v : %v", err, i, e)
		}

		return err
	}

	return nil
}
