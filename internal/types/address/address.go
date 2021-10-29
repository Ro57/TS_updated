package address

type AddressScheme interface {
	GenerateKey(seed [32]byte) PrivateKey
	ParseAddr(ps string) (Address, error)
}

type PrivateKey interface {
	GetPublicKey() []byte
	Sign(data []byte) []byte
	Equal(PrivateKey) bool
}

type Address interface {
	CheckSig(data []byte, signature []byte) bool
	String() string
}
