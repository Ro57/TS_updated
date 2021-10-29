package pktchain

import (
	"crypto/sha256"
	"math/rand"
	"time"
	"token-strike/internal/types/pkt"
)

type SimplePktChain struct {
}

var _ pkt.PktChain = (*SimplePktChain)(nil)

func (p SimplePktChain) BlockHashAtHeight(i int32) []byte {
	var result []byte
	if i < p.CurrentHeight() {
		sha := sha256.Sum256([]byte(string(i)))
		result = sha[:]
	}
	return result
}

func (p SimplePktChain) CurrentHeight() int32 {
	//TODO: think how represents that number as const
	height := (time.Now().Unix() - 1566269808) / 60
	return int32(height)
}

func (p *SimplePktChain) VerifyProof(annProof pkt.AnnProof) int32 {
	return annProof.Num
}

func (p *SimplePktChain) AnnounceData(data []byte) chan pkt.AnnProof {
	annProof := make(chan pkt.AnnProof)

	// payload simulations
	go func() {
		n := randomSleep(500, 2000)
		time.Sleep(time.Duration(n) * time.Millisecond)
		annProof <- pkt.AnnProof{
			Num: p.CurrentHeight(),
		}
	}()

	return annProof
}

func randomSleep(min, max int) int {
	return rand.Intn(max-min) + min
}
