package utils_test

import (
	"encoding/hex"
	"math"
	"testing"
	"time"
	"token-strike/internal/database"
	"token-strike/internal/database/repository"
	"token-strike/internal/types"
	"token-strike/internal/utils"
	"token-strike/tsp2p/server/DB"

	"github.com/golang/protobuf/proto"
)

const (
	aliceIndex = iota
	bobIndex
	christyIndex
)

func randomSeed(l, offset int) [32]byte {
	bytes := [32]byte{}
	for i := 0; i < l; i++ {
		bytes[i] = byte(i + offset)
	}
	return bytes
}

func TestAllFunctions(t *testing.T) {
	var activeAddressScheme types.AddressScheme = &utils.AddressScheme{}
	var activePktChain types.PktChain = &utils.PktChain{}
	seedSlice := [][32]byte{randomSeed(32, 0), randomSeed(32, 32), randomSeed(32, 64)}
	privKeySlice := []types.PrivateKey{}
	addressSlice := []types.Address{}

	db, err := database.Connect("./test.db")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			t.Error(err)
		}

		err = db.Clear()
		if err != nil {
			t.Error(err)
		}

	}()
	tokendb := repository.NewBbolt(db)

	for _, s := range seedSlice {
		privKeySlice = append(privKeySlice, activeAddressScheme.GenerateKey(s))
	}

	for _, k := range privKeySlice {

		address, err := activeAddressScheme.ParseAddr(k.Public())
		if err != nil {
			t.Fatal(err)
		}

		addressSlice = append(addressSlice, address)
	}

	issuerSeed := randomSeed(32, 96)
	issuerPubKey := hex.EncodeToString(issuerSeed[:])

	token := DB.Token{
		Count:        10,
		Expiration:   math.MaxInt32,
		Creation:     time.Now().Unix(),
		IssuerPubkey: issuerPubKey,
		Urls: []string{
			"http://localhost:3333/token1",
		},
	}

	state := &DB.State{
		Token: &token,
		Owners: []*DB.Owner{
			{
				HolderWallet: addressSlice[aliceIndex].String(),
				Count:        6,
			},
			{
				HolderWallet: addressSlice[bobIndex].String(),
				Count:        4,
			},
		},
		Locks: nil,
	}

	stateBytes, err := proto.Marshal(state)
	if err != nil {
		t.Error(err)
	}

	stateHash := hex.EncodeToString(stateBytes)

	block := &DB.Block{
		PrevBlock:      "0000000000000000000000000000000000000000000000000000000000000000",
		Justifications: nil,
		Creation:       time.Now().Unix(),
		State:          stateHash,
		PktBlockHash:   string(activePktChain.BlockHashAtHeight(activePktChain.CurrentHeight())),
		PktBlockHeight: activePktChain.CurrentHeight(),
		Height:         0,
	}

	bs0, err := proto.Marshal(block)
	if err != nil {
		t.Error(err)
	}

	sig := privKeySlice[aliceIndex].Sign(bs0)
	block.Signature = string(sig)

	tokenID := hex.EncodeToString(bs0)

	tokendb.SaveIssuerTokenDB(tokenID, issuerPubKey)

	tokendb.IssueTokenDB(tokenID, &token, block, []*DB.Owner{})
}
