package repository

import (
	"encoding/hex"
	"fmt"
	"time"

	"token-strike/internal/database"
	"token-strike/internal/errors"
	"token-strike/tsp2p/server/DB"

	"github.com/golang/protobuf/proto"
	"go.etcd.io/bbolt"
)

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
		if tokenBucket == nil {
			return errors.RootBucketNotFoundErr
		}

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

func (b *Bbolt) SyncBlock(name string, blocks []*DB.Block) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return errors.RootBucketNotFoundErr
		}

		tokenBucket := rootBucket.Bucket([]byte(name))
		if tokenBucket == nil {
			return errors.TokenNotFoundErr
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

func (b *Bbolt) SaveBlock(name string, block *DB.Block) error {
	err := b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return errors.RootBucketNotFoundErr
		}

		tokenBucket := rootBucket.Bucket([]byte(name))
		if tokenBucket == nil {
			return errors.TokenNotFoundErr
		}

		if hex.EncodeToString(tokenBucket.Get(database.RootHashKey)) != block.PrevBlock {
			return fmt.Errorf(
				"invalid hash of the previous block want %s but get %s",
				hex.EncodeToString(tokenBucket.Get(database.RootHashKey)),
				block.PrevBlock,
			)
		}

		blockHash, err := block.GetHash()
		if err != nil {
			return err
		}

		err = tokenBucket.Put(database.RootHashKey, blockHash)
		if err != nil {
			return err
		}

		blockBytes, errMarshal := proto.Marshal(block)
		if errMarshal != nil {
			return errMarshal
		}

		chainBucket := tokenBucket.Bucket(database.ChainKey)
		if chainBucket == nil {
			return errors.ChainBucketNotFoundErr
		}

		return chainBucket.Put(blockHash, blockBytes)
	})
	if err != nil {
		return err
	}

	for _, justification := range block.GetJustifications() {
		err := b.ApplyJustification(name, justification)
		if err != nil {
			return err
		}
	}

	return nil
}
