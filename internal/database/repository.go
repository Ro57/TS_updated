package database

import (
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/replicator"
)

type DBRepository interface {
	GetRepository
	SaveRepository
}

type SaveRepository interface {
	SaveBlock(name string, block *DB.Block) error
	SyncBlock(name string, blocks []*DB.Block) error
	// Add new token to local DB
	SaveIssuerTokenDB(name string, issuer string) error
	AssemblyBlock(name string, justifications []*DB.Justification) (*DB.Block, error)
	// Issue new token
	IssueTokenDB(name string, offer *DB.Token, block *DB.Block, state *DB.State) error
	TransferTokens(htlc, lock string) error
	LockToken(lock *lock.Lock) error
}

type GetRepository interface {
	GetTokenList() ([]*replicator.Token, error)
	GetToken(name string) (replicator.Token, error)
	GetIssuerTokens() (replicator.IssuerToken, error)
	GetChainInfoDB(tokenId string) (*replicator.ChainInfo, error)
	GetMerkleBlockDB(tokenId, hash string) ([]*replicator.MerkleBlock, error)
}
