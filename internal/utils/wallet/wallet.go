package wallet

import (
	"context"
	"errors"
	"token-strike/internal/database"
	"token-strike/internal/types"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/pktchain"
	"token-strike/internal/utils/privkey"
	"token-strike/tsp2p/server/tokenstrike"

	"google.golang.org/grpc"
)

type SimpleWallet struct {
	address        address.SimpleAddress
	privateKey     privkey.SimplePrivateKey
	pkt            pktchain.SimplePktChain
	scheme         address.SimpleAddressScheme
	db             database.DBRepository
	invClient      tokenstrike.TokenStrikeClient
	issuerInvSlice []tokenstrike.TokenStrikeClient
}

var _ types.Wallet = &SimpleWallet{}

func CreateWallet(cfg config.Config, pk privkey.SimplePrivateKey, http string, issuerUrlHints []string) (*SimpleWallet, error) {
	if issuerUrlHints == nil {
		return nil, errors.New("issuer url collection empty")
	}

	conn, err := grpc.DialContext(
		context.TODO(),
		http,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	invClient := tokenstrike.NewTokenStrikeClient(conn)
	walletAddress := address.NewSimpleAddress(pk.GetPublicKey())

	pkt, ok := cfg.Chain.(*pktchain.SimplePktChain)
	if !ok {
		return nil, errors.New("pkt type incorrect")
	}

	scheme, ok := cfg.Scheme.(*address.SimpleAddressScheme)
	if !ok {
		return nil, errors.New("address scheme type incorrect")
	}

	issuerClients, err := getClients(issuerUrlHints)
	if err != nil {
		return nil, err
	}

	return &SimpleWallet{
		address:        walletAddress,
		privateKey:     pk,
		pkt:            *pkt,
		scheme:         *scheme,
		db:             cfg.DB,
		invClient:      invClient,
		issuerInvSlice: issuerClients,
	}, nil

}

func getClients(issuerUrls []string) (issuerSlice []tokenstrike.TokenStrikeClient, err error) {
	for _, url := range issuerUrls {
		conn, err := grpc.DialContext(
			context.TODO(),
			url,
			grpc.WithInsecure(),
		)
		if err != nil {
			return nil, err
		}

		client := tokenstrike.NewTokenStrikeClient(conn)

		issuerSlice = append(issuerSlice, client)
	}
	return
}
