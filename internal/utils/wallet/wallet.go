package wallet

import (
	"context"
	"errors"
	"io"
	"token-strike/internal/database"
	address2 "token-strike/internal/types/address"
	"token-strike/internal/types/users"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/pktchain"
	"token-strike/tsp2p/server/tokenstrike"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SimpleWallet struct {
	address        address.SimpleAddress
	privateKey     address2.PrivateKey
	pkt            pktchain.SimplePktChain
	scheme         address.SimpleAddressScheme
	db             database.DBRepository
	invEvents      tokenstrike.TokenStrike_SubscribeClient
	invClient      tokenstrike.TokenStrikeClient
	issuerInvSlice []tokenstrike.TokenStrikeClient
	subChannel     tokenstrike.TokenStrike_SubscribeClient
}

var _ users.Wallet = &SimpleWallet{}

func CreateWallet(cfg config.Config, pk address2.PrivateKey, http string, issuerUrlHints []string) (*SimpleWallet, error) {
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

	subChannel, err := invClient.Subscribe(context.TODO(), &empty.Empty{})
	if err != nil {
		return nil, err
	}
	walletAddress := address.NewSimpleAddress(pk.GetPublicKey())

	events, err := invClient.Subscribe(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

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
		invEvents:      events,
		invClient:      invClient,
		subChannel:     subChannel,
		issuerInvSlice: issuerClients,
	}, nil

}

func (s SimpleWallet) selectEvent(events tokenstrike.TokenStrike_SubscribeClient) error {
	for {
		data, err := events.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		switch d := data.Data.(type) {
		case *tokenstrike.Data_Block:
			// TODO: set token name to save
			s.db.SaveBlock("", d.Block)
		}
	}
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
