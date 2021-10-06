package node

type UseCase interface {
	ConnectUseCase
	GetUseCase
	SaveUseCase
}

type ConnectUseCase interface {
	Connect() error
	Disconnect()
}

// GetUseCase describe how to work with node
type GetUseCase interface {
	// GetKey ... todo need to discuss and add params
	GetKey() error
	// GetUtxo ... todo need to discuss and add params
	GetUtxo() error
	// GetBlock get the selected block from node
	GetBlock() error
	// GetBestBlock get the last block in the valid chain
	GetBestBlock() error
	// GetBlockHash ...
	GetBlockHash() error
	// GetHtlcContract ...  todo need to discuss and add params
	GetHtlcContract() error
}

type SaveUseCase interface{}
