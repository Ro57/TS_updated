package utils_test

import (
	"encoding/hex"
	"fmt"
	"math"
	"testing"
	"time"
	address2 "token-strike/internal/types/address"

	"token-strike/internal/database"
	"token-strike/internal/database/repository"
	"token-strike/internal/types/pkt"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/issuer"
	"token-strike/internal/utils/pktchain"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/internal/utils/wallet"
	"token-strike/tsp2p/server/DB"
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
	http                                       = "0.0.0.0:3333"
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

	go func() {
		err = tokenstrikemock.NewServer(tokendb, isaacAddress, httpIsaac)
		if err != nil {
			t.Error(err)
		}
	}()

	go func() {
		err = tokenstrikemock.NewServer(tokendb, isaacAddress, httpAlice)
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(1 * time.Second)

	cfg := &config.Config{
		DB:     tokendb,
		Chain:  activePktChain,
		Scheme: activeAddressScheme,
	}

	issuer, err := issuer.CreateIssuer(cfg, isaacPrivateKey, httpIsaac)
	if err != nil {
		t.Error(err)
	}

	alice, err := wallet.CreateWallet(*cfg, alicePrivateKey, httpAlice, []string{httpIsaac})
	if err != nil {
		t.Error(err)
	}

	tokenID, err := issuer.IssueToken(
		[]*DB.Owner{
			{
				HolderWallet: aliceAddress.String(),
				Count:        6,
			},
			{
				HolderWallet: bobAddress.String(),
				Count:        4,
			},
		},
		math.MaxInt32,
	)
	if err != nil {
		t.Error(err)
	}

	lockID, err := alice.LockTokens(config.LockArgs{
		TokenId:    tokenID,
		Amount:     3,
		Recipient:  bobAddress.String(),
		SecretHash: "",
	})
	if err != nil {
		t.Error(err)
	}

	transferHash, err := alice.SendTokens(tokenID, lockID, []byte(""))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hex.EncodeToString(transferHash))
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
