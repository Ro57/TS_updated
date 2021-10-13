package utils

<<<<<<< HEAD
<<<<<<< HEAD
import "token-strike/internal/types"

func (e *Address) GenerateKey(randomSeed []byte) types.Key {
	return types.Key{}
=======
=======
>>>>>>> 1819400 (feat(sprint1): add GenerateKey method)
import (
	ed "crypto/ed25519"

	"token-strike/internal/types"
)

func (e *Address) GenerateKey(randomSeed []byte) types.Key {
	return ed.NewKeyFromSeed(randomSeed)
<<<<<<< HEAD
>>>>>>> 1819400 (feat(sprint1): add GenerateKey method)
=======
>>>>>>> 1819400 (feat(sprint1): add GenerateKey method)
}
