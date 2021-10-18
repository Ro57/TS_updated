package utils_test

import (
	"crypto/ed25519"
	"crypto/sha256"
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

var ()

func init() {

}

func randomSeed(l, offset int) []byte {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(i + offset)
	}
	return bytes
}

func TestAllFunctions(t *testing.T) {
	tokenName := "smt"
	address := &utils.Address{}
	PktChain := &utils.PktChain{}
	seedSlice := [][]byte{randomSeed(32, 0), randomSeed(32, 32), randomSeed(32, 64)}
	privKeySlice := []types.Key{}
	addressSlice := []string{}

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
		privKeySlice = append(privKeySlice, address.GenerateKey(s))
	}

	for _, k := range privKeySlice {
		h := sha256.New()
		h.Write(k.Public().(ed25519.PublicKey))

		address := hex.EncodeToString(h.Sum(nil))
		addressSlice = append(addressSlice, address)
	}

	issuerPubKey := addressSlice[aliceIndex]

	tokendb.SaveIssuerTokenDB(tokenName, issuerPubKey)

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
				HolderWallet: addressSlice[aliceIndex],
				Count:        6,
			},
			{
				HolderWallet: addressSlice[bobIndex],
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
		PktBlockHash:   string(PktChain.BlockHashAtHeight(PktChain.CurrentHeight())),
		PktBlockHeight: PktChain.CurrentHeight(),
		Height:         0,
	}

	bs0, err := proto.Marshal(block)
	if err != nil {
		t.Error(err)
	}

	sig := address.Sign(privKeySlice[aliceIndex], bs0)
	block.Signature = string(sig)

	tokendb.IssueTokenDB(tokenName, &token, block, []*DB.Owner{})
}
