package utils_test

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
	"token-strike/internal/utils/tokenstrikemock"

	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/rpcservice"
	"token-strike/tsp2p/server/tokenstrike"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoiface"
)

func (suite *TestSuite) TestInsert() {

	type args struct {
		hash       string
		message    protoiface.MessageV1
		msgType    uint32
		expiration int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid test",
			args: args{
				hash:       "test Hash",
				message:    &DB.State{},
				expiration: 12312312,
			},
			want: "dGVzdCBIYXNoL1wvXDEyMzEyMzEy",
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {

			if got := suite.tokenStrike.Insert(
				tokenstrikemock.MempoolEntry{
					Hash:       tt.args.hash,
					Message:    tt.args.message,
					Type:       tt.args.msgType,
					Expiration: tt.args.expiration,
				},
			); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *TestSuite) TestSendingMessages() {
	suite.preRunSendingMessage()

	tk := tokenstrikemock.New(nil, nil)

	err := tk.AddPeer(":5555")
	if err != nil {
		suite.T().Fatal(err)
	}

	insertHash := tk.Insert(
		tokenstrikemock.MempoolEntry{
			Hash:       hex.EncodeToString(suite.sendingMessageTestData.hash),
			Message:    suite.sendingMessageTestData.block,
			Type:       2,
			Expiration: 123,
		},
	)
	fmt.Printf("insertHash: %s \n", insertHash)

	// waiting while goroutine will process data
	time.Sleep(time.Second * 2)

	if !strings.EqualFold(
		suite.sendingMessageTestData.block.String(),
		suite.sendingMessageTestData.wallet.data.GetBlock().String(),
	) {
		suite.T().Fatal("blocks not equal")
	}
}

func (suite *TestSuite) preRunSendingMessage() {
	// create new block
	suite.sendingMessageTestData.block = &DB.Block{
		PrevBlock:      "test",
		Justifications: nil,
		Creation:       0,
		State:          "test",
		PktBlockHash:   "test",
		PktBlockHeight: 0,
		Height:         0,
		Signature:      "test",
	}
	suite.sendingMessageTestData.hash, _ = suite.sendingMessageTestData.block.GetHash()

	// create inv for new block
	suite.sendingMessageTestData.invs = []*tokenstrike.Inv{
		{
			Parent:     suite.sendingMessageTestData.hash[:],
			Type:       tokenstrike.TYPE_BLOCK,
			EntityHash: suite.sendingMessageTestData.hash[:],
		},
	}

	// strat wallet server
	go suite.createWalletMockServer()
	time.Sleep(time.Second * 1)
}

func (suite *TestSuite) createWalletMockServer() {
	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		suite.T().Fatal(err)
	}

	grpcServer := grpc.NewServer()

	suite.sendingMessageTestData.wallet = &walletMockServer{}

	rpcservice.RegisterRPCServiceServer(grpcServer, suite.sendingMessageTestData.wallet)
	err = grpcServer.Serve(lis)
	if err != nil {
		suite.T().Fatal(err)
	}
}
