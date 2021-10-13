package utils

type Key []byte
type Address interface {
	CheckSig(address string, signature []byte, data []byte) bool
	GenerateKey(randomSeed []byte) Key
	Sign(k Key, data []byte) []byte
}
