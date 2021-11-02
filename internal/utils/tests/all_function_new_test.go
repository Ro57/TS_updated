package utils_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
	"token-strike/internal/database"
	"token-strike/internal/database/repository"
	issuerNew "token-strike/internal/issuer"
	address2 "token-strike/internal/types/address"
	"token-strike/internal/types/pkt"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/pktchain"
	"token-strike/internal/wallet"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/rpcservice"
)

// creating keys
var (
	isaacPrivateKey = (&address.SimpleAddressScheme{}).GenerateKey(randomSeed(32, 0))
	isaacAddress    = address.NewSimpleAddress(isaacPrivateKey.GetPublicKey())

	alicePrivateKey = (&address.SimpleAddressScheme{}).GenerateKey(randomSeed(32, 32))
	aliceAddress    = address.NewSimpleAddress(alicePrivateKey.GetPublicKey())

	bobPrivateKey = (&address.SimpleAddressScheme{}).GenerateKey(randomSeed(32, 32))
	bobAddress    = address.NewSimpleAddress(bobPrivateKey.GetPublicKey())
)

const (
	httpIsaac = "0.0.0.0:3333"
	httpAlice = "0.0.0.0:3334"
)

// creating additional variables
var (
	activePktChain      pkt.PktChain           = &pktchain.SimplePktChain{}
	activeAddressScheme address2.AddressScheme = &address.SimpleAddressScheme{}
)

func TestAllFunctionsNew(t *testing.T) {
	// initialization the database
	db, err := database.Connect("./test.db")
	if err != nil {
		panic(err)
	}
	defer closeDB(db, t)

	tokendb := repository.NewBbolt(db)

	randomSecret := make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	rand.Read(randomSecret)
	htlcFL := sha256.Sum256(randomSecret)
	htlcSL := sha256.Sum256(htlcFL[:])

	cfg := &config.Config{
		Chain:  activePktChain,
		Scheme: activeAddressScheme,
	}

	go func() {
		err = issuerNew.NewServer(cfg, tokendb, isaacPrivateKey, httpIsaac)
		if err != nil {
			t.Error(err)
		}
	}()

	go func() {
		err = wallet.NewServer(tokendb, alicePrivateKey, httpAlice, []string{httpIsaac})
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(1 * time.Second)

	issuer, err := issuerNew.CreateClient(httpIsaac, "asd")
	if err != nil {
		t.Error(err)
	}

	alice, err := wallet.CreateClient(httpAlice, httpIsaac)
	if err != nil {
		t.Error(err)
	}

	tokenID, err := issuer.IssueToken(
		context.Background(),
		&rpcservice.IssueTokenRequest{
			Owners: []*DB.Owner{
				{
					HolderWallet: aliceAddress.String(),
					Count:        6,
				},
				{
					HolderWallet: bobAddress.String(),
					Count:        4,
				},
			},

			Expiration: math.MaxInt32,
		},
	)
	if err != nil {
		t.Error(err)
	}

	lockID, err := alice.LockToken(
		context.Background(),
		&rpcservice.LockTokenRequest{
			TokenId:    tokenID.TokenId,
			Amount:     3,
			Recipient:  bobAddress.String(),
			SecretHash: hex.EncodeToString(htlcSL[:]),
		})
	if err != nil {
		t.Error(err)
	}

	transferHash, err := alice.SendToken(
		context.Background(),
		&rpcservice.TransferTokensRequest{
			Token: tokenID.TokenId,
			Lock:  lockID.LockId,
			Htlc:  randomSecret,
		})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hex.EncodeToString(transferHash.Txid))
}

func closeDB(db *database.TokenStrikeDB, t *testing.T) {
	err := db.Close()
	if err != nil {
		t.Error(err)
	}

	err = db.Clear()
	if err != nil {
		t.Error(err)
	}

}
