package replicatorrpc

import (
	"context"
	"log"
	"net"
	"testing"
	"token-strike/tsp2p/server/replicator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

var lis *bufconn.Listener
var errImp = status.Error(codes.Unimplemented, "GetBlockSequence not implimented")

func init() {
	lis = bufconn.Listen(1000)
	s := grpc.NewServer()

	replicator.RegisterReplicatorServer(s, &Server{
		domain: "https://test.com",
	})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetBlockSequence(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.GetBlockSequence(ctx, &replicator.GetBlockSequenceRequest{
		Name: "someToken",
	})
	if err.Error() != status.Error(codes.Unimplemented, "GetBlockSequence not implimented").Error() {
		t.Fatalf("Get error %v want %v", err, errImp)
	}
}
