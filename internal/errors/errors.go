package errors

import "errors"

var (
	EmptyAddressErr = errors.New("empty address")

	TokensDBNotFound       = errors.New("tokens DB not created")
	InfoNotFoundErr        = errors.New("token info not found")
	StateNotFoundErr       = errors.New("state not found")
	TokenNotFoundErr       = errors.New("token does not found")
	BlockNotFoundErr       = errors.New("block info not found")
	RootHashNotFoundErr    = errors.New("root hash not found")
	LastBlockNotFoundErr   = errors.New("last block not found")
	ChainBucketNotFoundErr = errors.New("chain bucket not found")
)
