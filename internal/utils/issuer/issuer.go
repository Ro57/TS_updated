package issuer

import (
	"token-strike/internal/types"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/tokenstrikemock"
)

type SimpleIssuer struct {
	invServer tokenstrikemock.TokenStrikeMock
}

var _ types.Issuer = &SimpleIssuer{}

func CreateIssuer(cfg config.Config, pk types.PrivateKey, http string) SimpleIssuer
