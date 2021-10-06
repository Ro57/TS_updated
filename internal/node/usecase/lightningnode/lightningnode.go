package lightningnode

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"token-strike/internal/node"
	"token-strike/tsp2p/server/lnrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ node.UseCase = &Node{}

// Node data struct for lnd node client
type Node struct {
	host            string
	certificatePath string
	closeConnection func()
	client          lnrpc.LightningClient
}

// New return exemplar of Node struct
func New(host, certificatePath string) *Node {
	return &Node{
		host:            host,
		certificatePath: certificatePath,
	}
}

// Connect get and save lnd client in Node struct
func (n *Node) Connect() error {
	tlsCredentials, err := loadTLSCredentials(n.certificatePath)
	if err != nil {
		return err
	}

	connection, err := grpc.Dial(n.host, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		return err
	}

	n.closeConnection = func() {
		connection.Close()
	}

	n.client = lnrpc.NewLightningClient(connection)

	return nil
}

// loadTLSCredentials gets tls cert for auth
func loadTLSCredentials(path string) (credentials.TransportCredentials, error) {
	cert, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

// Disconnect close dial connection for grpc client
func (n Node) Disconnect() {
	n.closeConnection()
}

func (n Node) GetKey() error {
	return fmt.Errorf("implement me")
}

func (n Node) GetUtxo() error {
	return fmt.Errorf("implement me")
}

func (n Node) GetBlock() error {
	return fmt.Errorf("implement me")
}

func (n Node) GetBestBlock() error {
	return fmt.Errorf("implement me")
}

func (n Node) GetBlockHash() error {
	return fmt.Errorf("implement me")
}

func (n Node) GetHtlcContract() error {
	return fmt.Errorf("implement me")
}
