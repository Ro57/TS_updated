package blockchainnode

// BlockchainNode describe how to work with node
type BlockchainNode interface{
	// Ping gets the status of node
	Ping()
	// GetKey ... todo need to discuss and add params
	GetKey()
	// GetUtxo ... todo need to discuss and add params
	GetUtxo()
	// GetBlock get the selected block from node
	GetBlock()
	// GetBestBlock get the last block in the valid chain
	GetBestBlock()
	// GetBlockHash ...
	GetBlockHash()
	// GetHtlsContract ...  todo need to discuss and add params
	GetHtlsContract()
}
