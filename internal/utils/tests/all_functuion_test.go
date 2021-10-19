package utils_test

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"math"
	"math/rand"
	"testing"
	"time"
	"token-strike/internal/database"
	"token-strike/internal/database/repository"
	"token-strike/internal/types"
	"token-strike/internal/utils"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/lock"
)

const (
	aliceIndex = iota
	bobIndex
	christyIndex
	isaacIndex
	lastIndex
)

func randomSeed(l, offset int) [32]byte {
	bytes := [32]byte{}
	for i := 0; i < l; i++ {
		bytes[i] = byte(i + offset)
	}
	return bytes
}

func TestAllFunctions(t *testing.T) {
	var activeAddressScheme types.AddressScheme = &utils.SimpleAddressScheme{}
	var activePktChain types.PktChain = &utils.SimplePktChain{}
	seedSlice := [][32]byte{randomSeed(32, 0), randomSeed(32, 32), randomSeed(32, 64), randomSeed(32, 96)}
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

	token := DB.Token{
		Count:        10,
		Expiration:   math.MaxInt32,
		Creation:     time.Now().Unix(),
		IssuerPubkey: addressSlice[isaacIndex].String(),
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

	tokendb.SaveIssuerTokenDB(tokenID, addressSlice[isaacIndex].String())

	tokendb.IssueTokenDB(tokenID, &token, block, []*DB.Owner{})

	// n3
	// generate random secret 32 byte
	randomSecret := make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	rand.Read(randomSecret)
	htlcFL := sha256.Sum256(randomSecret)
	htlcSL := sha256.Sum256(htlcFL[:])
	//populate lock with special data todo check the addresses
	lockEl := &lock.Lock{
		Count:          3,
		Recipient:      addressSlice[isaacIndex].String(),
		Sender:         addressSlice[aliceIndex].String(),
		HtlcSecretHash: hex.EncodeToString(htlcSL[:]),
		ProofCount:     block.PktBlockHeight + 60,
		PktBlockHash:   activePktChain.BlockHashAtHeight(activePktChain.CurrentHeight()),
		PktBlockHeight: uint32(activePktChain.CurrentHeight()),
		Signature:      "",
	}

	bs0, err = proto.Marshal(lockEl)
	if err != nil {
		t.Error(err)
	}

	sig = privKeySlice[isaacIndex].Sign(bs0) //todo is it right index for signing?
	lockEl.Signature = hex.EncodeToString(sig)

	//prprd slice of Inv todo think maybe we dont need in extra init, cause each elem have default values
	var Invs = make([]tokenstrikemock.TokenStrikeMock, lastIndex, lastIndex)
	Invs[aliceIndex] = tokenstrikemock.TokenStrikeMock{}
	Invs[bobIndex] = tokenstrikemock.TokenStrikeMock{}
	Invs[christyIndex] = tokenstrikemock.TokenStrikeMock{}
	Invs[isaacIndex] = tokenstrikemock.TokenStrikeMock{}

}
