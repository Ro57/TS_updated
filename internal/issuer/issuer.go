package issuer

import (
	"context"
	"errors"
	"flag"
	"net"
	"token-strike/internal/database"
	"token-strike/internal/database/repository"

	address2 "token-strike/internal/types/address"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var _ rpcservice.RPCServiceServer = &Issuer{}

type Issuer struct {
	config     *config.Config
	tokendb    *repository.Bbolt
	invServer  *tokenstrikemock.TokenStrikeMock
	subChannel tokenstrike.TokenStrike_SubscribeClient

	private address2.PrivateKey
	address address2.Address
	peers   []string
}

func NewServer(cfg *config.Config, tokendb *repository.Bbolt, pk address2.PrivateKey, target string) error {
	flag.Parse()

	lis, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	issuerImpl, err := CreateIssuer(cfg, tokendb, pk)
	if err != nil {
		return err
	}

	rpcservice.RegisterRPCServiceServer(grpcServer, issuerImpl)
	return grpcServer.Serve(lis)
}

func CreateIssuer(cfg *config.Config, tokendb database.DBRepository, pk address2.PrivateKey) (*Issuer, error) {
	invServer := tokenstrikemock.New(tokendb, address.NewSimpleAddress(pk.GetPublicKey()))
	issuer := &Issuer{
		private:   pk,
		address:   address.NewSimpleAddress(pk.GetPublicKey()),
		peers:     make([]string, 0),
		config:    cfg,
		invServer: invServer,
	}

	return issuer, nil
}
func (i *Issuer) AddPeer(ctx context.Context, request *rpcservice.PeerRequest) (*empty.Empty, error) {
	if request.Url != "" {
		i.peers = append(i.peers, request.Url)
		return nil, nil
	}
	return nil, errors.New("url cannot is empty")
}

func (i *Issuer) SendToken(ctx context.Context, req *rpcservice.TransferTokensRequest) (*rpcservice.TransferTokensResponse, error) {
	panic("implement me")
}
func (i *Issuer) LockToken(ctx context.Context, req *rpcservice.LockTokenRequest) (*rpcservice.LockTokenResponse, error) {
	panic("implement me")
}
func (i *Issuer) PostData(ctx context.Context, req *tokenstrike.Data) (*tokenstrike.PostDataResp, error) {
	return i.invServer.PostData(ctx, req)
}

func (i *Issuer) Inv(ctx context.Context, req *tokenstrike.InvReq) (*tokenstrike.InvResp, error) {
	return i.invServer.Inv(ctx, req)
}

func (i *Issuer) GetTokenStatus(ctx context.Context, req *tokenstrike.TokenStatusReq) (*tokenstrike.TokenStatus, error) {
	return i.invServer.GetTokenStatus(ctx, req)
}
