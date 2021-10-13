package utils

<<<<<<< HEAD
import "token-strike/internal/types"

func (e *Address) GenerateKey(randomSeed []byte) types.Key {
	return types.Key{}
=======
import (
	ed "crypto/ed25519"

	"token-strike/internal/types"
)

func (e *Address) GenerateKey(randomSeed []byte) types.Key {
	return ed.NewKeyFromSeed(randomSeed)
>>>>>>> 1819400 (feat(sprint1): add GenerateKey method)
}
