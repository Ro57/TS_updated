package types

// AnnProof — lock expiration time in PKT blocks
type AnnProof int32

type PktChain interface {
	CurrentHeight() int32
	BlockHashAtHeight(int32) []byte
	AnnounceData([]byte) chan AnnProof

	// VerifyProof yields the block height when the proof was mades
	VerifyProof(AnnProof) int32
}
