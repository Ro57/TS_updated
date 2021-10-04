package replicatorrpc

import (
	"context"
	"log"
	"net"
	"testing"

	"token-strike/server/replicatorrpc"
	"token-strike/tsp2p/server/replicator"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var (
	domain = "http://localhost"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	listener   *bufconn.Listener
	grpcClient replicator.ReplicatorClient
	conn       *grpc.ClientConn
}

func (suite *TestSuite) SetupTest() {
	err := suite.initListener()
	if err != nil {
		suite.T().Fatalf(err.Error())
	}

	suite.initGrpcClient()
}

func (suite *TestSuite) initGrpcClient() {
	ctx := context.Background()

	var err error
	suite.conn, err = grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(suite.bufDialer), grpc.WithInsecure())
	if err != nil {
		suite.T().Fatalf("Failed to dial bufnet: %v", err)
	}

	suite.grpcClient = replicator.NewReplicatorClient(suite.conn)
}

func (suite *TestSuite) initListener() error {
	suite.listener = bufconn.Listen(1000)
	s := grpc.NewServer()

	server, err := replicatorrpc.New("", domain)
	if err != nil {
		return err
	}

	replicator.RegisterReplicatorServer(s, server)

	go func() {
		if err := s.Serve(suite.listener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
	return nil
}

func (suite *TestSuite) bufDialer(context.Context, string) (net.Conn, error) {
	return suite.listener.Dial()
}
