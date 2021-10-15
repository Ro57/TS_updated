package utils

import "time"

func (p PktChain) CurrentHeight() int32 {
	//todo think how represents that number as const
	height := (time.Now().Unix() - 1566269808) / 60
	return int32(height)
}
