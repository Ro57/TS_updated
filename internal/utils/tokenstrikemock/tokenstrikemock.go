package tokenstrikemock

import (
	"token-strike/tsp2p/server/tokenstrike"
)

type TokenStrikeMock struct {

}

var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}