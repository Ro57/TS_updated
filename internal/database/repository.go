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
	// Add new token to local DB
	SaveBlock(name string, block *DB.Block) error
	TransferTokens(tokenID, lock string) error
	SyncBlock(name string, blocks []*DB.Block) error
	LockToken(tokenID string, lock *lock.Lock) error
	SaveIssuerTokenDB(name string, issuer string) error
	AssemblyBlock(name string, justifications []*DB.Justification) (*DB.Block, error)
	IssueTokenDB(name string, offer *DB.Token, block *DB.Block, state *DB.State) error

	// TODO: rework this method. Get justification array and apply all
	ApplyJustification(tokenID string, justification *DB.Justification) error
}

type GetRepository interface {
	GetTokenList() ([]*replicator.Token, error)
	GetToken(name string) (replicator.Token, error)
	GetIssuerTokens() (replicator.IssuerToken, error)
	GetTokenStatus(token string) (*DB.Block, *DB.State, error)
	GetChainInfoDB(tokenId string) (*replicator.ChainInfo, error)
	GetMerkleBlockDB(tokenId, hash string) ([]*replicator.MerkleBlock, error)
}
