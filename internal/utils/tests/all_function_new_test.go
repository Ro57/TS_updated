package utils_test

import (
	"math"
	"testing"

	"token-strike/internal/database"
	"token-strike/internal/database/repository"
	"token-strike/internal/types"
	"token-strike/internal/utils/address"
	"token-strike/internal/utils/config"
	"token-strike/internal/utils/issuer"
	"token-strike/internal/utils/pktchain"
	"token-strike/tsp2p/server/DB"
)

// creating keys
var (
	isaacPrivateKey = (&address.SimpleAddressScheme{}).GenerateKey(randomSeed(32, 0))

	alicePrivateKey = (&address.SimpleAddressScheme{}).GenerateKey(randomSeed(32, 32))
	aliceAddress    = address.NewSimpleAddress(alicePrivateKey.GetPublicKey())

	bobPrivateKey = (&address.SimpleAddressScheme{}).GenerateKey(randomSeed(32, 32))
	bobAddress    = address.NewSimpleAddress(bobPrivateKey.GetPublicKey())
)

// creating additional variables
var (
	activePktChain types.PktChain = &pktchain.SimplePktChain{}
)

func TestAllFunctionsNew(t *testing.T) {
	// initialization the database
	db, err := database.Connect("./test.db")
	if err != nil {
		panic(err)
	}
	defer closeDB(db, t)

	tokendb := repository.NewBbolt(db)

	cfg := &config.Config{
		DB: tokendb,
	}

	issuer, err := issuer.CreateIssuer(cfg, isaacPrivateKey, "tcp://0.0.0.0:3333")
	if err != nil {
		t.Error(err)
	}

	issuer.IssueToken(
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
