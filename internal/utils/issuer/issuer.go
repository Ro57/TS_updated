package issuer

import (
	"token-strike/internal/types"
	"token-strike/internal/utils/config"
)

type SimpleIssuer struct{}

var _ types.Issuer = &SimpleIssuer{}

func CreateIssuer(cfg config.Config, pk types.PrivateKey) SimpleIssuer
