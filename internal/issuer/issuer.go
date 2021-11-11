package issuer

import (
	"context"
	"net"
	"token-strike/internal/database"
	"token-strike/internal/database/repository"

	"token-strike/internal/types/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var _ rpcservice.RPCServiceServer = &Issuer{}

type Issuer struct {
	config    *config.Config
	tokendb   database.DBRepository
	invServer *tokenstrikemock.TokenStrikeMock

	private address.PrivateKey
	address address.Address
}

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

func NewServer(cfg *config.Config, tokendb *repository.Bbolt, pk address.PrivateKey, target string) error {
	lis, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	issuerImpl, err := CreateIssuer(cfg, tokendb, pk)
	if err != nil {
		return err
	}

	issuerImpl.bootBlockGenerator()

	rpcservice.RegisterRPCServiceServer(grpcServer, issuerImpl)
	return grpcServer.Serve(lis)
}

func CreateIssuer(cfg *config.Config, tokendb database.DBRepository, pk address.PrivateKey) (*Issuer, error) {
	invServer := tokenstrikemock.New(tokendb, pk.Address())
	issuer := &Issuer{
		private:   pk,
		address:   pk.Address(),
		config:    cfg,
		invServer: invServer,
		tokendb:   tokendb,
	}

	return issuer, nil
}
func (i *Issuer) AddPeer(ctx context.Context, request *rpcservice.PeerRequest) (*empty.Empty, error) {
	return &empty.Empty{}, i.invServer.AddPeer(request.Url)
}

func (i *Issuer) SendToken(ctx context.Context, req *rpcservice.TransferTokensRequest) (*rpcservice.TransferTokensResponse, error) {
	return &rpcservice.TransferTokensResponse{}, nil
}
func (i *Issuer) LockToken(ctx context.Context, req *rpcservice.LockTokenRequest) (*rpcservice.LockTokenResponse, error) {
	return &rpcservice.LockTokenResponse{}, nil
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
