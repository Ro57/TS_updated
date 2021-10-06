package replicatorrpc

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"testing"
	"token-strike/config"
	"token-strike/internal/database"
	"token-strike/internal/database/repository"

	"token-strike/server/replicatorrpc"
	"token-strike/tsp2p/server/replicator"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var (
	domain = "http://localhost"

	pktMock      interface{}
	httpMock     interface{}
	databaseMock database.DBRepository
	hashFuncMock func() string
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

// TODO: implement initRestClient function

func (suite *TestSuite) initListener() error {
	suite.listener = bufconn.Listen(1000)
	s := grpc.NewServer()

	server, err := replicatorrpc.New(
		&config.Config{
			RpcPort:  "",
			HttpPort: "",
			Domain:   domain,
		},
		repository.NewBbolt(suite.initTempDatabase()),
	)
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

func (suite *TestSuite) initTempDatabase() *database.TokenStrikeDB {
	path := tempfile()
	defer os.RemoveAll(path)

	db, err := database.Connect(path)
	if err != nil {
		suite.T().Fatal(err)
	} else if db == nil {
		suite.T().Fatal("expected db")
	}

	if s := db.GetClient().Path(); s != path {
		suite.T().Fatalf("unexpected path: %s", s)
	}

	return db
}

// tempfile returns a temporary file path.
func tempfile() string {
	f, err := ioutil.TempFile("", "bolt-")
	if err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
	return f.Name()
}
