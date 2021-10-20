package types

type AddressScheme interface {
	GenerateKey(seed [32]byte) PrivateKey
	ParseAddr(ps string) (Address, error)
}

type PrivateKey interface {
	Address() Address
	Sign(data []byte) []byte
	Equal(PrivateKey) bool
}

type Address interface {
	CheckSig(data []byte, signature []byte) bool
	String() string
}
