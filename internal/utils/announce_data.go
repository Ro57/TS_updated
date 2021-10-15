package utils

import (
	"math/rand"
	"time"
	"token-strike/internal/types"
)

func (p *PktChain) AnnounceData(data []byte) chan types.AnnProof {
	annProof := make(chan types.AnnProof)

	// payload simulations
	go func() {
		n := randomSleep(500, 2000)
		time.Sleep(time.Duration(n) * time.Millisecond)
		annProof <- types.AnnProof{
			Num: p.CurrentHeight(),
		}
	}()

	return annProof
}

func randomSleep(min, max int) int {
	return rand.Intn(max-min) + min
}
