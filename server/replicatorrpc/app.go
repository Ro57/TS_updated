package replicatorrpc

import (
	"context"
	"fmt"
	"net"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	domain string
}

var _ replicator.ReplicatorServer = (*Server)(nil)

func New(host, domain string) (*Server, error) {
	var serv = &Server{
		domain: domain,
	}

	return serv, nil
}

func (s *Server) RunGRPCServer(host string) error {
	opts := []grpc.ServerOption{}

	root := grpc.NewServer(opts...)
	replicator.RegisterReplicatorServer(root, s)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	fmt.Printf("Replcication sever start on port: %v \n", host)

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
