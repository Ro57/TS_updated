package repository

import (
	"encoding/json"
	stdErrors "errors"
	"fmt"
	"token-strike/internal/errors"

	"token-strike/internal/database"
	"token-strike/server/replicatorrpc"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/replicator"

	"github.com/golang/protobuf/proto"
	"go.etcd.io/bbolt"
)

func NewBbolt(db database.TokenStrikeDB) *Bbolt {
	return &Bbolt{
		db: db,
	}
}

type Bbolt struct {
	db database.TokenStrikeDB
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

func (b *Bbolt) GetIssuerTokens() (tokens replicatorrpc.IssuerTokens, err error) {
	err = b.db.View(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return errors.TokensDBNotFound
		}

		issuerTokensBytes := rootBucket.Get(database.IssuerTokens)
		if issuerTokensBytes == nil {
			tokens = replicatorrpc.IssuerTokens{}
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
	err = b.db.Update(func(tx *bbolt.Tx) error {

		// getting chain buckets
		rootBucket := tx.Bucket(database.TokensKey)
		tokenBucket := rootBucket.Bucket([]byte(tokenId))
		dbStateByte := tokenBucket.Get(database.StateKey)

		// unmarshal chain state
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
				return stdErrors.New(fmt.Sprintf(
					"block doesnot find by root hash=%v",
					currentHash,
				))
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
	panic("implement me")
}

func (b *Bbolt) SyncBlock(name string, blocks []*DB.Block) error {
	panic("implement me")
}

func (b *Bbolt) SaveIssuerTokenDB(name string, offer *DB.Token) error {
	panic("implement me")
}

func (b *Bbolt) AssemblyBlock(name string, justifications []*DB.Justification) (*DB.Block, error) {
	panic("implement me")
}

func (b *Bbolt) IssueTokenDB(name string, offer *DB.Token, block *DB.Block, recipient []*DB.Owner) error {
	panic("implement me")
}
