package tokenstrikemock

import (
	"flag"
	"net"

	"token-strike/internal/database"
	address2 "token-strike/internal/types/address"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/pktchain"
	"token-strike/tsp2p/server/tokenstrike"

	"google.golang.org/grpc"
)

const (
	NeedData     = true
	DontNeedData = false
)

type TokenStrikeMock struct {
	bboltDB       database.DBRepository
	issuer        address2.Address
	pktChain      pktchain.SimplePktChain
	addressScheme address.SimpleAddressScheme
	invCache      map[string]tokenstrike.Inv
	peers         []tokenstrike.TokenStrike_SubscribeServer
}

var _ tokenstrike.TokenStrikeServer = &TokenStrikeMock{}

func New(db database.DBRepository, issuer address2.Address) *TokenStrikeMock {
	return &TokenStrikeMock{
		bboltDB:       db,
		issuer:        issuer,
		pktChain:      pktchain.SimplePktChain{},
		addressScheme: address.SimpleAddressScheme{},
		invCache:      make(map[string]tokenstrike.Inv),
	}
}
func NewServer(db database.DBRepository, issuer address2.Address, target string) error {
	flag.Parse()

	lis, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	tokenstrike.RegisterTokenStrikeServer(grpcServer, New(db, issuer))
	return grpcServer.Serve(lis)
}
