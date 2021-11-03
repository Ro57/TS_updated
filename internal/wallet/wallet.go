package wallet

import (
	"context"
	"errors"
	"net"
	addressScheme "token-strike/internal/utils/simple"

	"token-strike/internal/database"
	addressTypes "token-strike/internal/types/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/pktchain"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/rpcservice"

	"google.golang.org/grpc"
)

type Server struct {
	address        addressScheme.SimpleAddress
	privateKey     addressTypes.PrivateKey
	pkt            pktchain.SimplePktChain
	scheme         addressScheme.SimpleAddressScheme
	db             database.DBRepository
	issuerInvSlice []rpcservice.RPCServiceClient
	inv            *tokenstrikemock.TokenStrikeMock
}

var _ rpcservice.RPCServiceServer = &Server{}

func CreateClient(target, selfAddr string) (rpcservice.RPCServiceClient, error) {
	conn, err := grpc.DialContext(
		context.TODO(),
		target,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	client := rpcservice.NewRPCServiceClient(conn)
	return client, nil
}

func NewServer(db database.DBRepository, pk addressTypes.PrivateKey, target string, issuerUrlHints []string) error {
	lis, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	walletImpl, err := CreateWallet(db, target, pk, issuerUrlHints)
	if err != nil {
		return err
	}

	rpcservice.RegisterRPCServiceServer(grpcServer, walletImpl)
	return grpcServer.Serve(lis)
}

func CreateWallet(db database.DBRepository, peerUrl string, pk addressTypes.PrivateKey, issuerUrlHints []string) (*Server, error) {
	if issuerUrlHints == nil {
		return nil, errors.New("issuer url collection empty")
	}

	cfg := config.DefaultSimpleConfig()

	walletAddress := addressScheme.NewSimpleAddress(pk.GetPublicKey())

	pkt := cfg.Chain.(*pktchain.SimplePktChain)
	scheme := cfg.Scheme.(*addressScheme.SimpleAddressScheme)

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
		inv:            tokenstrikemock.New(db, walletAddress),
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

		_, err = client.AddPeer(context.Background(), &rpcservice.PeerRequest{Url: peerUrl})
		if err != nil {
			return nil, err
		}

		issuerSlice = append(issuerSlice, client)
	}
	return
}
