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

	errWant := status.Error(codes.Unimplemented, "GetBlockSequence not implimented")

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.GetBlockSequence(ctx, &replicator.GetBlockSequenceRequest{
		Name: "someToken",
	})
	if err.Error() != errWant.Error() {
		t.Fatalf("Get error %v want %v", err, errWant)
	}
}

func TestServer_GetIssuerTokens(t *testing.T) {
	ctx := context.Background()

	errWant := status.Error(codes.Unimplemented, "GetIssuerTokens not implimented")

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.GetIssuerTokens(ctx, &replicator.GetIssuerTokensRequest{})
	if err.Error() != errWant.Error() {
		t.Fatalf("Get error %v want %v", err, errWant)
	}
}

func TestServer_GetHeaders(t *testing.T) {
	ctx := context.Background()

	errWant := status.Error(codes.Unimplemented, "GetHeaders not implimented")

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.GetHeaders(ctx, &replicator.GetHeadersRequest{})
	if err.Error() != errWant.Error() {
		t.Fatalf("Get error %v want %v", err, errWant)
	}
}

func TestServer_IssueToken(t *testing.T) {
	ctx := context.Background()

	errWant := status.Error(codes.Unimplemented, "IssueToken not implimented")

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.IssueToken(ctx, &replicator.IssueTokenRequest{})
	if err.Error() != errWant.Error() {
		t.Fatalf("Get error %v want %v", err, errWant)
	}
}

func TestServer_GetToken(t *testing.T) {
	ctx := context.Background()

	errWant := status.Error(codes.Unimplemented, "GetToken not implimented")

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.GetToken(ctx, &replicator.GetTokenRequest{})
	if err.Error() != errWant.Error() {
		t.Fatalf("Get error %v want %v", err, errWant)
	}
}

func TestServer_GetTokenList(t *testing.T) {
	ctx := context.Background()

	errWant := status.Error(codes.Unimplemented, "GetTokenList not implimented")

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.GetTokenList(ctx, &replicator.GetTokenListRequest{})
	if err.Error() != errWant.Error() {
		t.Fatalf("Get error %v want %v", err, errWant)
	}
}

func TestServer_GenerateURL(t *testing.T) {
	ctx := context.Background()

	errWant := status.Error(codes.Unimplemented, "GenerateURL not implimented")

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.GenerateURL(ctx, &replicator.GenerateURLRequest{})
	if err.Error() != errWant.Error() {
		t.Fatalf("Get error %v want %v", err, errWant)
	}
}

func TestServer_GetUrlSequence(t *testing.T) {
	ctx := context.Background()

	errWant := status.Error(codes.Unimplemented, "GetUrlSequence not implimented")

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := replicator.NewReplicatorClient(conn)

	_, err = client.GetUrlSequence(ctx, &replicator.GetUrlSequenceRequest{})
	if err.Error() != errWant.Error() {
		t.Fatalf("Get error %v want %v", err, errWant)
	}
}