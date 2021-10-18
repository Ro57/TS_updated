package utils

func (a Address) String() string {
	return string(a.publicKey)
}
