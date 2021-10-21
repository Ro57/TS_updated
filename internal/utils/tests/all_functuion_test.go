package utils_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
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
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

const (
	aliceIndex = iota
	bobIndex
	christyIndex

	// this is issuer
	isaacIndex
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
		address := k.Address()

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

	sig := privKeySlice[isaacIndex].Sign(bs0)
	block.Signature = hex.EncodeToString(sig)

	blockSigned, err := proto.Marshal(block)
	if err != nil {
		t.Error(err)
	}

	blockHash := sha256.Sum256(blockSigned)

	tokenID := hex.EncodeToString(blockHash[:])

	tokendb.SaveIssuerTokenDB(tokenID, addressSlice[isaacIndex].String())

	tokendb.IssueTokenDB(tokenID, &token, block, []*DB.Owner{})

	// save block hash for next inv logic
	block0Hash := sha256.Sum256(bs0)

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

	//make isaac inv mock
	IsaacTokenStrikeServer := tokenstrikemock.TokenStrikeMock{}

	//prepare object with one token data
	lock3Hash := sha256.Sum256(bs0)

	//saving all locks to map for gets it later
	locksPost := make(map[string]*lock.Lock, 0)
	locksPost[string(lock3Hash[:])] = lockEl

	invs := []*tokenstrike.Inv{
		{
			Parent:     block0Hash[:],
			Type:       tokenstrike.TYPE_LOCK,
			EntityHash: lock3Hash[:], //todo is it right data?
		},
	}

	var InvReq = &tokenstrike.InvReq{
		Invs: invs,
	}

	resp, err := IsaacTokenStrikeServer.Inv(context.TODO(), InvReq)
	if err != nil {
		t.Error(err)
	}

	needed := resp.Needed

	if needed != nil {
		for idx, need := range needed {
			if need {
				//take need elem from list by idx of resp
				lockdata := locksPost[string(invs[idx].EntityHash)]
				DataReq := &tokenstrike.Data{
					Data: &tokenstrike.Data_Lock{lockdata},
				}

				//send selected lock and NOW skip check of warning
				_, err := IsaacTokenStrikeServer.PostData(context.TODO(), DataReq)
				if err != nil {
					t.Error(err) //todo if it dont pass we should send block, not lock
				}

			}
		}

	}

}
