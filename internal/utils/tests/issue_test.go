package utils_test

import "token-strike/tsp2p/server/DB"

type db interface {
	store(block DB.Block, state DB.State) error
}
