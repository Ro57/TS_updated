package wallet

import (
	"crypto/sha256"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

func (s SimpleWallet) SendTokens(tokenId string, lockId []byte, secretHex []byte) ([]byte, error) {
	transferTokens := &tokenstrike.TransferTokens{
		Htlc: secretHex[:],
		Lock: lockId[:],
	}

	transferTokensB, err := proto.Marshal(transferTokens)
	if err != nil {
		return nil, err
	}

	transferTokensHash := sha256.Sum256(transferTokensB)

	return transferTokensHash[:], nil
}
