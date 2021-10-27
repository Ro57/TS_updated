package wallet

import (
	"token-strike/internal/types"
	"token-strike/internal/utils/config"
)

type SimpleWallet struct{}

var _ types.Wallet = &SimpleWallet{}

func CreateWallet(cfg config.Config, pk types.PrivateKey) SimpleWallet
