package idgen

import (
	"fmt"
	"strconv"
	"strings"
)

const separator = ":"

func Encode(blockHash string, number int) string {
	return fmt.Sprint(blockHash, separator, number)
}

func Decode(ID string) (*string, *int, error) {
	tuple := strings.Split(ID, separator)

	blockHash := tuple[0]

	justificationNumber, err := strconv.Atoi(tuple[1])
	if err != nil {
		return nil, nil, err
	}

	return &blockHash, &justificationNumber, nil
}
