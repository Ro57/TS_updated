package utils

import "time"

func (p SimplePktChain) CurrentHeight() int32 {
	//TODO: think how represents that number as const
	height := (time.Now().Unix() - 1566269808) / 60
	return int32(height)
}
