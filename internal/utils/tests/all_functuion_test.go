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
	"token-strike/tsp2p/server/justifications"
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

	stateBytes, err := state.GetStateHash()
	if err != nil {
		t.Error(err)
	}

	block := &DB.Block{
		PrevBlock:      "0000000000000000000000000000000000000000000000000000000000000000",
		Justifications: nil,
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
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

	tokendb.IssueTokenDB(tokenID, &token, block, state)

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
	AliceTokenStrikeServer := tokenstrikemock.TokenStrikeMock{}

	lockSigned, err := proto.Marshal(lockEl)
	if err != nil {
		t.Error(err)
	}

	lockHash := sha256.Sum256(lockSigned)

	//saving all locks to map for gets it later
	locksPost := make(map[string]*lock.Lock, 0)
	locksPost[lockEl.GetSignature()] = lockEl

	invs := []*tokenstrike.Inv{
		{
			Parent:     blockHash[:],
			Type:       tokenstrike.TYPE_LOCK,
			EntityHash: lockHash[:],
		},
	}

	resp, err := IsaacTokenStrikeServer.Inv(context.TODO(), &tokenstrike.InvReq{
		Invs: invs,
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Needed != nil {
		for idx, need := range resp.Needed {
			if need {
				//take need elem from list by idx of resp
				lockdata := locksPost[string(invs[idx].EntityHash)]
				DataReq := &tokenstrike.Data{
					Data: &tokenstrike.Data_Lock{lockdata},
				}

				//send selected lock and NOW skip check of warning
				_, err := IsaacTokenStrikeServer.PostData(context.TODO(), DataReq)
				if err != nil {
					t.Error(err)
				}

				// Update state: add lock to state and reduce number of tokens owned by A by 3
				state.Locks = []*lock.Lock{locksPost[string(invs[idx].EntityHash)]}
				state.Owners[0].Count = state.Owners[0].Count - lockEl.Count
			}
		}
	}

	stateBytes, err = state.GetStateHash()
	if err != nil {
		t.Error(err)
	}

	blockIsaac := &DB.Block{
		PrevBlock: "0000000000000000000000000000000000000000000000000000000000000001",
		Justifications: []*DB.Justification{
			{
				Content: &DB.Justification_Transfer{
					Transfer: &justifications.TranferToken{
						HtlcSecret: lockEl.HtlcSecretHash,
						Lock:       lockEl.Signature,
					},
				},
			},
		},
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
		PktBlockHash:   string(activePktChain.BlockHashAtHeight(activePktChain.CurrentHeight())),
		PktBlockHeight: activePktChain.CurrentHeight(),
		Height:         1,
	}

	bs0Isaac, err := proto.Marshal(blockIsaac)
	if err != nil {
		t.Error(err)
	}

	sigIsaac := privKeySlice[isaacIndex].Sign(bs0Isaac)
	blockIsaac.Signature = hex.EncodeToString(sigIsaac)

	blockIsaacSigned, err := proto.Marshal(blockIsaac)
	if err != nil {
		t.Error(err)
	}

	blockIsaacHash := sha256.Sum256(blockIsaacSigned)

	invsIsaac := []*tokenstrike.Inv{
		{
			Parent:     blockHash[:],
			Type:       tokenstrike.TYPE_BLOCK,
			EntityHash: blockIsaacHash[:],
		},
	}

	resp, err = AliceTokenStrikeServer.Inv(context.TODO(), &tokenstrike.InvReq{
		Invs: invsIsaac,
	})
	if err != nil {
		t.Error(err)
	}

	if resp.Needed != nil {
		for _, need := range resp.Needed {
			if need {
				DataReq := &tokenstrike.Data{
					Data: &tokenstrike.Data_Block{Block: blockIsaac},
				}

				//send selected lock and NOW skip check of warning
				_, err := IsaacTokenStrikeServer.PostData(context.TODO(), DataReq)
				if err != nil {
					t.Error(err)
				}

			}
		}
	}

}
