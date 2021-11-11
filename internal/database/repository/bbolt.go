package repository

import (
	"token-strike/internal/database"
)

func NewBbolt(db *database.TokenStrikeDB) *Bbolt {
	return &Bbolt{
		db: db,
	}
}

type Bbolt struct {
	db *database.TokenStrikeDB
}
