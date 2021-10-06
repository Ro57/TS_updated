package repository

import (
	stdErrors "errors"
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

func (b *Bbolt) GetIssuerTokens() (replicatorrpc.IssuerTokens, error) {
	panic("implement me")
}

func (b *Bbolt) GetChainInfoDB(tokenId string) (*replicator.ChainInfo, error) {
	panic("implement me")
}

func (b *Bbolt) GetMerkleBlockDB(tokenId, hash string) ([]*replicator.MerkleBlock, error) {
	panic("implement me")
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
