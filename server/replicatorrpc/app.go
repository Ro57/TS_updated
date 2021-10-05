package replicatorrpc

import (
	"context"
	"log"
	"net"
	"net/http"

	"token-strike/config"
	"token-strike/internal/database"
	"token-strike/tsp2p/server/replicator"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	rpcPort  string
	restPort string
	domain   string

	db database.DBRepository
}

var _ replicator.ReplicatorServer = (*Server)(nil)

func New(cfg *config.Config, dbRep database.DBRepository) (*Server, error) {
	var serv = &Server{
		rpcPort:  cfg.RpcPort,
		restPort: cfg.HttpPort,
		domain:   cfg.Domain,
		db:       dbRep,
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
