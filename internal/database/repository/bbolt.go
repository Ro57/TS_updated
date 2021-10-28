package repository

import (
	"encoding/hex"
	"encoding/json"
	stdErrors "errors"
	"fmt"
	"time"
	"token-strike/internal/database"
	"token-strike/internal/errors"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/replicator"

	"github.com/golang/protobuf/proto"
	"go.etcd.io/bbolt"
)

func NewBbolt(db *database.TokenStrikeDB) *Bbolt {
	return &Bbolt{
		db: db,
	}
}

type Bbolt struct {
	db *database.TokenStrikeDB
}

func (b *Bbolt) GetTokenStatus(token string) (*DB.Block, *DB.State, error) {

	var (
		genesisBlock = &DB.Block{}
		genesisState = &DB.State{}
	)

	err := b.db.View(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return stdErrors.New("tokens do not exist")
		}

		tokenBucket := rootBucket.Bucket([]byte(token))
		if tokenBucket == nil {
			return errors.TokenNotFoundErr
		}

		blockData := tokenBucket.Get(database.GenesisBlockKey)
		if blockData == nil {
			return errors.InfoNotFoundErr
		}

		err := proto.Unmarshal(blockData, genesisBlock)
		if err != nil {
			return err
		}

		stateData := tokenBucket.Get(database.GenesisBlockKey)
		if stateData == nil {
			return errors.InfoNotFoundErr
		}

		err = proto.Unmarshal(stateData, genesisState)
		if err != nil {
			return err
		}

		return nil
	})

	return genesisBlock, genesisState, err
}

func (b *Bbolt) GetTokenList() ([]*replicator.Token, error) {
	var resultList []*replicator.Token

	err := b.db.View(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return stdErrors.New("tokens do not exist")
		}

		return rootBucket.ForEach(func(k, _ []byte) error {
			tokenBucket := rootBucket.Bucket(k)

			// skip useless buckets
			if tokenBucket == nil {
				return nil
			}

			var dbToken DB.Token
			err := proto.Unmarshal(tokenBucket.Get(database.InfoKey), &dbToken)
			if err != nil {
				return err
			}

			token := replicator.Token{
				Name:  string(k),
				Token: &dbToken,
				Root:  string(tokenBucket.Get(database.RootHashKey)),
			}

			resultList = append(resultList, &token)
			return nil
		})
	})

	return resultList, err
}

func (b *Bbolt) GetToken(name string) (replicator.Token, error) {
	var (
		token = replicator.Token{
			Name:  name,
			Token: &DB.Token{},
		}
	)

	err := b.db.View(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return errors.TokensDBNotFound
		}

		tokenBucket := rootBucket.Bucket([]byte(name))
		if tokenBucket == nil {
			return errors.TokenNotFoundErr
		}
		infoBytes := tokenBucket.Get(database.InfoKey)
		if infoBytes == nil {
			return errors.InfoNotFoundErr
		}

		err := proto.Unmarshal(infoBytes, token.Token)
		if err != nil {
			return err
		}

		token.Root = string(tokenBucket.Get(database.RootHashKey))
		return nil
	})

	return token, err
}

func (b *Bbolt) GetIssuerTokens() (tokens replicator.IssuerToken, err error) {
	err = b.db.View(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return errors.TokensDBNotFound
		}

		issuerTokensBytes := rootBucket.Get(database.IssuerTokens)
		if issuerTokensBytes == nil {
			tokens = replicator.IssuerToken{}
			return nil
		}

		err := json.Unmarshal(issuerTokensBytes, &tokens)
		if err != nil {
			return err
		}

		return nil
	})
	return
}

func (b *Bbolt) GetChainInfoDB(tokenId string) (*replicator.ChainInfo, error) {
	var (
		resp = &replicator.ChainInfo{
			State:  &DB.State{},
			Blocks: []*DB.Block{},
			Root:   "",
		}

		err     error
		dbstate DB.State
	)
	err = b.db.View(func(tx *bbolt.Tx) error {

		// getting chain buckets
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return errors.TokensDBNotFound
		}

		tokenBucket := rootBucket.Bucket([]byte(tokenId))
		if tokenBucket == nil {
			return errors.TokensDBNotFound
		}

		// unmarshal chain state
		dbStateByte := tokenBucket.Get(database.StateKey)
		err = proto.Unmarshal(dbStateByte, &dbstate)
		if err != nil {
			return err
		}

		// getting chain blocks
		var (
			rootHash    = tokenBucket.Get(database.RootHashKey)
			chainBucket = tokenBucket.Bucket(database.ChainKey)

			currentHash = rootHash
		)

		for {
			blockBytes := chainBucket.Get(currentHash)
			if blockBytes == nil {
				return fmt.Errorf(
					"block doesnot find by root hash=%v",
					currentHash,
				)
			}

			var block DB.Block
			err := proto.Unmarshal(blockBytes, &block)
			if err != nil {
				return err
			}

			resp.Blocks = append(resp.Blocks, &block)

			if block.PrevBlock == "" {
				break
			}

			currentHash = []byte(block.PrevBlock)

		}

		resp.State = &dbstate
		resp.Root = string(rootHash)

		return nil
	})

	return resp, err
}

func (b *Bbolt) GetMerkleBlockDB(tokenId, hash string) ([]*replicator.MerkleBlock, error) {
	var (
		blocks = []*replicator.MerkleBlock{}
	)

	err := b.db.Update(func(tx *bbolt.Tx) error {

		tokensBucket := tx.Bucket(database.TokensKey)
		tokenBucket := tokensBucket.Bucket([]byte(tokenId))
		if tokenBucket == nil {
			return errors.TokenNotFoundErr
		}

		infoBytes := tokenBucket.Get(database.InfoKey)
		if infoBytes == nil {
			return errors.InfoNotFoundErr
		}

		rootHash := tokenBucket.Get(database.RootHashKey)
		if rootHash == nil {
			return errors.RootHashNotFoundErr
		}

		if string(rootHash) != hash {

			var (
				currentHash = rootHash
				chainBucket = tokenBucket.Bucket(database.ChainKey)
			)

			for {
				blockBytes := chainBucket.Get(currentHash)
				if blockBytes == nil {
					return errors.BlockNotFoundErr
				}

				var block DB.Block
				err := proto.Unmarshal(blockBytes, &block)
				if err != nil {
					return err
				}

				if string(currentHash) == hash {
					break
				}

				merkleBlock := replicator.MerkleBlock{
					Hash:     string(currentHash),
					PrevHash: block.PrevBlock,
				}

				blocks = append(blocks, &merkleBlock)
				currentHash = []byte(merkleBlock.PrevHash)
			}
		}

		return nil
	})

	return blocks, err
}

func (b *Bbolt) SaveBlock(name string, block *DB.Block) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket := rootBucket.Bucket([]byte(name))

		if string(tokenBucket.Get(database.RootHashKey)) != block.PrevBlock {
			return fmt.Errorf(
				"invalid hash of the previous block want %s but get %s",
				tokenBucket.Get(database.RootHashKey),
				block.PrevBlock,
			)
		}

		blockSignatureBytes := []byte(block.GetSignature())

		err = tokenBucket.Put(database.RootHashKey, blockSignatureBytes)
		if err != nil {
			return err
		}

		blockBytes, errMarshal := proto.Marshal(block)
		if errMarshal != nil {
			return errMarshal
		}

		chainBucket, err := tokenBucket.CreateBucketIfNotExists(database.ChainKey)
		if err != nil {
			return err
		}

		return chainBucket.Put(blockSignatureBytes, blockBytes)
	})
}

func (b *Bbolt) SyncBlock(name string, blocks []*DB.Block) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			var err error
			rootBucket, err = tx.CreateBucketIfNotExists(database.TokensKey)
			if err != nil {
				return err
			}
		}

		tokenBucket, err := rootBucket.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		// below is the algorithm for searching and saving blocks in order
		var (
			numberCurrentBlock = 0
			quantityBlocks     = len(blocks)
		)

		// works until it passes through the entire number of blocks
		for quantityBlocks != numberCurrentBlock {

			// every time with a new pass we request the current signature
			// after saving it changes
			currentSignature := string(tokenBucket.Get(database.RootHashKey))

			// if the first block is incorrect
			if currentSignature != blocks[0].PrevBlock {
				// start searching the entire array
				for index, block := range blocks {
					// after finding the required block, save it
					if currentSignature == block.Signature {
						// save block
						err := b.SaveBlock(name, block)
						if err != nil {
							return err
						}
						blocks = append(blocks[:index], blocks[index+1:]...)
					}
				}
			} else { // if the blocks are in the correct order we save block
				err := b.SaveBlock(name, blocks[0])
				if err != nil {
					return err
				}
				blocks = append(blocks[:0], blocks[1:]...)
			}
			numberCurrentBlock++
		}

		return nil
	})
}

func (b *Bbolt) SaveIssuerTokenDB(name, issuer string) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokens := rootBucket.Get(database.IssuerTokens)
		if tokens == nil {
			tokens, _ = json.Marshal(replicator.IssuerToken{})
		}

		var issuerTokens replicator.IssuerToken
		errUnmarshal := json.Unmarshal(tokens, &issuerTokens)
		if errUnmarshal != nil {
			return errUnmarshal
		}

		issuerTokens.AddToken(issuer, name)

		issuerTokensBytes, errMarshal := json.Marshal(issuerTokens)
		if errMarshal != nil {
			return errMarshal
		}

		return rootBucket.Put(database.IssuerTokens, issuerTokensBytes)
	})
}

func (b *Bbolt) AssemblyBlock(name string, justifications []*DB.Justification) (*DB.Block, error) {
	state := DB.State{}

	block := &DB.Block{
		Justifications: justifications,
		Creation:       time.Now().Unix(),
	}

	// TODO: implement after created chain interface
	//currentBlockHash, currentBlockHeight, err := s.chain.GetBestBlock()
	//if err != nil {
	//	return nil, err
	//}
	//
	//block.PktBlockHash = currentBlockHash.String()
	//block.PktBlockHeight = currentBlockHeight

	err := b.db.Update(func(tx *bbolt.Tx) error {
		var lastBlock DB.Block

		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket := rootBucket.Bucket([]byte(name))

		lastHash := tokenBucket.Get(database.RootHashKey)
		if lastHash == nil {
			return errors.LastBlockNotFoundErr
		}

		chainBucket := tokenBucket.Bucket(database.ChainKey)
		if chainBucket == nil {
			return errors.ChainBucketNotFoundErr
		}

		jsonBlock := chainBucket.Get(lastHash)
		if jsonBlock == nil {
			return errors.LastBlockNotFoundErr
		}

		nativeErr := proto.Unmarshal(jsonBlock, &lastBlock)
		if nativeErr != nil {
			return fmt.Errorf("unmarshal block form json: %v", nativeErr)
		}

		jsonState := tokenBucket.Get(database.StateKey)
		if jsonState == nil {
			return errors.StateNotFoundErr
		}

		nativeErr = proto.Unmarshal(jsonState, &state)
		if nativeErr != nil {
			return fmt.Errorf("marshal new state: %v", nativeErr)

		}

		block.Height = lastBlock.Height + 1

		//hashState := encoder.CreateHash(jsonState) // TODO: rework with new signer
		hashState := []byte("implement me")
		block.State = hex.EncodeToString(hashState)

		// TODO: Change to signature generation
		block.Signature = block.GetState()
		block.PrevBlock = string(lastHash)

		newBlockBytes, nativeErr := proto.Marshal(block)
		if nativeErr != nil {
			return err
		}

		blockSignatureBytes := []byte(block.GetSignature())
		err = tokenBucket.Put(database.RootHashKey, blockSignatureBytes)
		if err != nil {
			return err
		}

		return chainBucket.Put(blockSignatureBytes, newBlockBytes)
	})

	if err != nil {
		return nil, err
	}

	return block, nil
}

func (b *Bbolt) IssueTokenDB(name string, offer *DB.Token, block *DB.Block, state *DB.State) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket, err := rootBucket.CreateBucket([]byte(name))
		if err != nil {
			return err
		}

		// if information about token did not exist then create
		if tokenBucket.Get(database.InfoKey) == nil {
			tokenBytes, err := proto.Marshal(offer)
			if err != nil {
				return err
			}

			err = tokenBucket.Put(database.InfoKey, tokenBytes)
			if err != nil {
				return err
			}
		}

		// if token state did not exist then create
		if tokenBucket.Get(database.StateKey) == nil {
			marshaledState, err := proto.Marshal(state)
			if err != nil {
				return err
			}
			errPut := tokenBucket.Put(database.StateKey, marshaledState)
			if errPut != nil {
				return errPut
			}

			errPut = tokenBucket.Put(database.GenesisStateKey, marshaledState)
			if errPut != nil {
				return errPut
			}
		}

		err = tokenBucket.Put(database.RootHashKey, []byte(""))
		if err != nil {
			return err
		}

		if string(tokenBucket.Get(database.RootHashKey)) != block.PrevBlock {
			return fmt.Errorf(
				"invalid hash of the previous block want %s but get %s",
				tokenBucket.Get(database.RootHashKey),
				block.PrevBlock,
			)
		}

		blockSignatureBytes := []byte(block.GetSignature())

		err = tokenBucket.Put(database.RootHashKey, blockSignatureBytes)
		if err != nil {
			return err
		}

		err = tokenBucket.Put(database.GenesisBlockKey, blockSignatureBytes)
		if err != nil {
			return err
		}

		blockBytes, errMarshal := proto.Marshal(block)
		if errMarshal != nil {
			return errMarshal
		}

		chainBucket, err := tokenBucket.CreateBucketIfNotExists(database.ChainKey)
		if err != nil {
			return err
		}
		return chainBucket.Put(blockSignatureBytes, blockBytes)
	})
}

func (b *Bbolt) TransferTokens(tokenID, lock string) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket := rootBucket.Bucket([]byte(tokenID))
		if err != nil {
			return err
		}

		var state DB.State
		stateBytes := tokenBucket.Get(database.StateKey)
		err = proto.Unmarshal(stateBytes, &state)
		if err != nil {
			return err
		}

		lockHashIndex := state.GetLockIndexByHash(lock, state.Locks)

		if lockHashIndex == nil {
			return fmt.Errorf("not found lock %v in state", lock)
		}

		lock := state.Locks[*lockHashIndex]

		recipientIndex := state.GetOwnerIndexByHolder(lock.Recipient, state.Owners)
		if recipientIndex == nil {
			index := len(state.Owners)
			recipientIndex = &index
			state.Owners = append(state.Owners, &DB.Owner{HolderWallet: lock.Recipient, Count: 0})
		}

		// Remove lock
		state.Locks = append(state.Locks[:*lockHashIndex], state.Locks[*lockHashIndex+1:]...)

		// Change balance
		state.Owners[*recipientIndex].Count = state.Owners[*recipientIndex].Count + lock.Count

		stateBytes, err = proto.Marshal(&state)
		if err != nil {
			return nil
		}

		return tokenBucket.Put(database.StateKey, stateBytes)
	})
}

func (b *Bbolt) LockToken(tokenID string, lock *lock.Lock) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket := rootBucket.Bucket([]byte(tokenID))
		if err != nil {
			return err
		}

		var state DB.State
		stateBytes := tokenBucket.Get(database.StateKey)
		err = proto.Unmarshal(stateBytes, &state)
		if err != nil {
			return err
		}

		state.Locks = append(state.Locks, lock)

		senderIndex := state.GetOwnerIndexByHolder(lock.Sender, state.Owners)
		if senderIndex == nil {
			return fmt.Errorf("holder with name %v not found in state", lock.Sender)
		}

		state.Owners[*senderIndex].Count = state.Owners[*senderIndex].Count - lock.Count

		stateBytes, err = proto.Marshal(&state)
		if err != nil {
			return nil
		}

		return tokenBucket.Put(database.StateKey, stateBytes)
	})
}

func (b *Bbolt) ApplyJustification(tokenID string, justification *DB.Justification) error {
	switch j := justification.Content.(type) {
	case *DB.Justification_Lock:
		return b.LockToken(tokenID, j.Lock.Lock)
	case *DB.Justification_Transfer:
		return b.TransferTokens(tokenID, j.Transfer.Lock)
	default:
		return nil
	}
}
