package replicatorrpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"log"
	"net"
	"net/http"
	"token-strike/config"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	rpcPort  string
	restPort string
	domain   string
}

var _ replicator.ReplicatorServer = (*Server)(nil)

func New(cfg *config.Config) (*Server, error) {
	var serv = &Server{
		rpcPort:  cfg.RpcPort,
		restPort: cfg.HttpPort,
		domain:   cfg.Domain,
	}

	return serv, nil
}

func (s *Server) RunRestServer() error {
	var opts []runtime.ServeMuxOption
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := replicator.RegisterReplicatorHandlerFromEndpoint(ctx, mux, s.rpcPort, dialOpts)
	if err != nil {
		return err
	}
	log.Printf("Replication REST server start on port: %v \n", s.restPort)
	err = http.ListenAndServe(s.restPort, mux)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) RunGRPCServer() error {
	listener, err := net.Listen("tcp", s.rpcPort)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	root := grpc.NewServer(opts...)
	replicator.RegisterReplicatorServer(root, s)

	log.Printf("Replication GRPC server start on port: %v \n", s.rpcPort)
	err = root.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) SyncChain(stream replicator.Replicator_SyncChainServer) error {
	return status.Error(codes.Unimplemented, "SyncChain not implimented")
}

func (s *Server) GetBlockSequence(ctx context.Context, req *replicator.GetBlockSequenceRequest) (*replicator.GetUrlTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetBlockSequence not implimented")
}

func (s *Server) GetIssuerTokens(ctx context.Context, req *replicator.GetIssuerTokensRequest) (*replicator.GetIssuerTokensResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetIssuerTokens not implimented")
}

func (s *Server) GetHeaders(ctx context.Context, req *replicator.GetHeadersRequest) (*replicator.GetHeadersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetHeaders not implimented")
}

func (s *Server) IssueToken(ctx context.Context, req *replicator.IssueTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "IssueToken not implimented")
}

func (s *Server) GetToken(ctx context.Context, req *replicator.GetTokenRequest) (*replicator.GetTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetToken not implimented")
}

func (s *Server) GetTokenList(ctx context.Context, req *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetTokenList not implimented")
}

func (s *Server) GenerateURL(ctx context.Context, req *replicator.GenerateURLRequest) (*replicator.GenerateURLResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GenerateURL not implimented")
}

func (s *Server) GetUrlSequence(ctx context.Context, req *replicator.GetUrlSequenceRequest) (*replicator.GetUrlSequenceResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetUrlSequence not implimented")
}
