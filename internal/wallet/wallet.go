package wallet

import (
	"context"
	"errors"
	"token-strike/internal/database"
	addressTypes "token-strike/internal/types/address"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/pktchain"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/rpcservice"

	"google.golang.org/grpc"
)

type Server struct {
	address        address.SimpleAddress
	privateKey     addressTypes.PrivateKey
	pkt            pktchain.SimplePktChain
	scheme         address.SimpleAddressScheme
	db             database.DBRepository
	issuerInvSlice []rpcservice.RPCServiceClient
	inv            *tokenstrikemock.TokenStrikeMock
}

var _ rpcservice.RPCServiceServer = &Server{}

func CreateWallet(db database.DBRepository, pk addressTypes.PrivateKey, peerUrl string, issuerUrlHints []string) (*Server, error) {
	if issuerUrlHints == nil {
		return nil, errors.New("issuer url collection empty")
	}

	cfg := config.DefaultSimpleConfig()

	walletAddress := address.NewSimpleAddress(pk.GetPublicKey())

	pkt := cfg.Chain.(*pktchain.SimplePktChain)
	scheme := cfg.Scheme.(*address.SimpleAddressScheme)

	issuerClients, err := getClients(peerUrl, issuerUrlHints)
	if err != nil {
		return nil, err
	}

	wallet := &Server{
		address:        walletAddress,
		privateKey:     pk,
		pkt:            *pkt,
		scheme:         *scheme,
		db:             db,
		issuerInvSlice: issuerClients,
		inv:            tokenstrikemock.New(db),
	}

	return wallet, nil
}

func getClients(peerUrl string, issuerUrls []string) (issuerSlice []rpcservice.RPCServiceClient, err error) {
	for _, url := range issuerUrls {
		conn, err := grpc.DialContext(
			context.TODO(),
			url,
			grpc.WithInsecure(),
		)
		if err != nil {
			return nil, err
		}

		client := rpcservice.NewRPCServiceClient(conn)
		client.AddPeer(context.Background(), &rpcservice.PeerRequest{Url: peerUrl})

		issuerSlice = append(issuerSlice, client)
	}
	return
}
