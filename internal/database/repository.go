package database

import (
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/replicator"
)

type DBRepository interface {
	GetRepository
	SaveRepository
}

type SaveRepository interface {
	SaveBlock(name string, block *DB.Block) error
	SaveIssuerTokenDB(name string, offer *DB.Token) error
	AssemblyBlock(name string, justifications []*DB.Justification) (*DB.Block, error)
	IssueTokenDB(name string, offer *DB.Token, block *DB.Block, recipient []*DB.Owner) error
}

type GetRepository interface {
	GetTokenList() ([]*replicator.Token, error)
	GetToken(name string) (replicator.Token, error)
	GetIssuerTokens() (replicator.IssuerTokens, error)
	GetChainInfoDB(tokenId string) (*replicator.ChainInfo, error)
	GetMerkleBlockDB(tokenId, hash string) ([]*replicator.MerkleBlock, error)
}
