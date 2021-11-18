package repository

import (
	"encoding/hex"
	"encoding/json"
	stdErrors "errors"
	"fmt"

	"token-strike/internal/database"
	"token-strike/internal/errors"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/replicator"

	"github.com/golang/protobuf/proto"
	"go.etcd.io/bbolt"
)

func (b *Bbolt) GetChainInfoDB(tokenId string) (*replicator.ChainInfo, error) {
	var (
		resp = &replicator.ChainInfo{
			State: &DB.State{},

			// Blocks are reversed
			Blocks: []*DB.Block{},
			Root:   "",
		}

		err   error
		state DB.State
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

		chainBucket := tokenBucket.Bucket(database.ChainKey)
		if chainBucket == nil {
			return errors.ChainBucketNotFoundErr
		}

		// unmarshal chain state
		dbStateByte := tokenBucket.Get(database.StateKey)
		err = proto.Unmarshal(dbStateByte, &state)
		if err != nil {
			return err
		}

		// getting chain blocks
		var (
			rootHash    = tokenBucket.Get(database.TipBlockHashKey)
			currentHash = rootHash
		)

		for {
			blockBytes := chainBucket.Get(currentHash)
			if blockBytes == nil {
				return fmt.Errorf(
					"block not found by root hash=%v",
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

			currentHash, err = hex.DecodeString(block.PrevBlock)
			if err != nil {
				return err
			}
		}

		resp.State = &state
		resp.Root = string(rootHash)

		return nil
	})

	return resp, err
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

		rootHash := tokenBucket.Get(database.TipBlockHashKey)
		if rootHash == nil {
			return errors.RootHashNotFoundErr
		}

		if string(rootHash) != hash {

			chainBucket := tokenBucket.Bucket(database.ChainKey)
			if chainBucket == nil {
				return errors.ChainBucketNotFoundErr
			}

			var currentHash = rootHash
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

				merkleBlock := replicator.MerkleBlock{
					Hash:     string(currentHash),
					PrevHash: block.PrevBlock,
				}

				blocks = append(blocks, &merkleBlock)

				if string(currentHash) == hash {
					break
				}

				currentHash = []byte(merkleBlock.PrevHash)
			}
		}

		return nil
	})

	return blocks, err
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
			if tokenBucket == nil { // skip useless buckets
				return nil
			}

			var tokenInfo DB.Token
			err := proto.Unmarshal(tokenBucket.Get(database.InfoKey), &tokenInfo)
			if err != nil {
				return err
			}

			token := replicator.Token{
				Name:  string(k),
				Token: &tokenInfo,
				Root:  string(tokenBucket.Get(database.TipBlockHashKey)),
			}

			resultList = append(resultList, &token)
			return nil
		})
	})

	return resultList, err
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

		stateData := tokenBucket.Get(database.GenesisStateKey)
		if stateData == nil {
			return errors.InfoNotFoundErr
		}

		err := proto.Unmarshal(stateData, genesisState)
		if err != nil {
			return err
		}

		blockData := tokenBucket.Get(database.GenesisBlockKey)
		if blockData == nil {
			return errors.InfoNotFoundErr
		}

		err = proto.Unmarshal(blockData, genesisBlock)
		if err != nil {
			return err
		}

		return nil
	})

	return genesisBlock, genesisState, err
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

		token.Root = string(tokenBucket.Get(database.TipBlockHashKey))
		return nil
	})

	return token, err
}
