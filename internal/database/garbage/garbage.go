package garbage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"token-strike/internal/database"
	"token-strike/internal/database/repository"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
)

const (
	tokenCount = 2
	blockCount = 5
)

type garbageToken struct {
	name  string
	token *DB.Token
	block []*DB.Block
}

func createHash(sid string) string {
	return hex.EncodeToString(createHashBytes([]byte(sid)))
}

func createHashBytes(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)

	return hasher.Sum(nil)
}

func Generate(db *database.TokenStrikeDB) error {
	repos := repository.NewBbolt(*db)

	for i := 0; i < tokenCount; i++ {
		// TODO: change on real algoritm of hash generator
		pubKey := hex.EncodeToString(createHashBytes([]byte{byte(i)}))

		name := fmt.Sprintf("token%v", i+1)
		blockchain := generateBlockChain(i)
		repos.IssueTokenDB(name,
			&DB.Token{
				Count:        int64(1000 * i),
				Expiration:   int32(time.Hour * time.Duration(i)),
				Creation:     time.Now().Unix(),
				IssuerPubkey: pubKey,
				Urls: []string{
					"https://replication.com/" + name,
				},
			},
			blockchain[0],
			[]*DB.Owner{},
		)

		for _, block := range blockchain {
			repos.SaveBlock(name, block)
		}
	}

	return nil
}

func generateBlockChain(index int) []*DB.Block {
	genesisBlock := &DB.Block{
		PrevBlock: "",
		Justifications: []*DB.Justification{
			&DB.Justification{
				Content: &DB.Justification_Genesis{
					Genesis: &justifications.Genesis{
						Token: fmt.Sprintf("token%v", index+1),
					},
				},
			},
		},
	}

	blockSlice := []*DB.Block{
		genesisBlock,
	}

	for i := 1; i < blockCount; i++ {
		htlcTransfer := createHash(fmt.Sprintf("htlc%v%v", i, index))
		htlcLock := createHash(fmt.Sprintf("htlcLock%v%v", i, index))
		lockID := createHash(fmt.Sprintf("lock%v%v", i, index))
		recipient := createHash(fmt.Sprintf("recipient%v%v", i, index))
		sender := createHash(fmt.Sprintf("sender%v%v", i, index))
		signature := createHash(fmt.Sprintf("signature%v%v", i, index))
		state := createHash(fmt.Sprintf("state%v%v", i, index))
		pktHash := createHash(fmt.Sprintf("pktHash%v%v", i, index))

		blockSlice = append(blockSlice, &DB.Block{
			PrevBlock: blockSlice[i-1].State,
			Justifications: []*DB.Justification{
				&DB.Justification{
					Content: &DB.Justification_Transfer{
						Transfer: &justifications.TranferToken{
							HtlcSecret: htlcTransfer,
							Lock:       lockID,
						},
					},
				},
				&DB.Justification{
					Content: &DB.Justification_Lock{
						Lock: &justifications.LockToken{
							Lock: &lock.Lock{
								Count:          int64(10 * i),
								Recipient:      recipient,
								Sender:         sender,
								HtlcSecretHash: htlcLock,
								ProofCount:     int32(2 * i),
								CreationHeight: uint64(i),
								Signature:      signature,
							},
						},
					},
				},
			},

			Creation:       time.Now().Unix(),
			State:          state,
			PktBlockHash:   pktHash,
			PktBlockHeight: int32(2 * i),
			Height:         uint64(i),
			Signature:      signature,
		})
	}

	return blockSlice
}
