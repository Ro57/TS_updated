package wallet

import (
	"token-strike/internal/types"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/tokenstrikemock"
)

type SimpleWallet struct {
	invServer tokenstrikemock.TokenStrikeMock
}

var _ types.Wallet = &SimpleWallet{}

func CreateWallet(cfg config.Config, pk types.PrivateKey, http string) SimpleWallet
