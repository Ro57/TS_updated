package replicatorrpc

import (
	"fmt"
	"net"

	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc"
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
