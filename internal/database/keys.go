package database

var (
	GenesisStateKey = []byte("genesis_state")
	GenesisBlockKey = []byte("genesis_block")

	IssuerTokens = []byte("issuer_tokens")

	InfoKey   = []byte("info")
	StateKey  = []byte("state")
	ChainKey  = []byte("chain")
	TokensKey = []byte("tokens")
	// rootHash is a hash of last block in chain
	TipBlockHashKey = []byte("rootHash")
	// Replication is a information about replication server configuration
	Replication = []byte("replication")
	// Issuer is a collection with key issuer pubKey
	Issuers = []byte("issuers")
)
