package wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

func (s SimpleWallet) SendTokens(tokenId string, lockId []byte, secret string) ([]byte, error) {
	secretHex, err := hex.DecodeString(secret)
	if err != nil {
		return nil, err
	}
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
