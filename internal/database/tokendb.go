package database

import (
	"fmt"
	"os"

	"go.etcd.io/bbolt"
)

const (
	permitions = 0666
)

type TokenStrikeDB struct {
	client bbolt.DB
	path   string
}

func Connect(path string) (*TokenStrikeDB, error) {
	db, err := bbolt.Open(path, permitions, nil)
	if err != nil {
		return nil, err
	}

	tsDB := &TokenStrikeDB{
		client: *db,
		path:   path,
	}

	return tsDB, nil
}

func (t *TokenStrikeDB) GetClient() *bbolt.DB {
	return &t.client
}

func (t *TokenStrikeDB) Close() error {
	return t.client.Close()
}

func (t *TokenStrikeDB) Ping() error {
	return t.client.View(func(tx *bbolt.Tx) error {
		return nil
	})
}

// Clear all database structure with buckets and their content
func (t *TokenStrikeDB) Clear() error {
	if err := os.Truncate(t.path, 0); err != nil {
		return fmt.Errorf("failed to truncate: %v", err)
	}

	t = nil
	return nil
}

func (t *TokenStrikeDB) Update(f func(tx *bbolt.Tx) error) error {
	return t.client.Update(f)
}

func (t *TokenStrikeDB) View(f func(tx *bbolt.Tx) error) error {
	return t.client.View(f)
}
