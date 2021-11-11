package utils

import (
	"io/ioutil"
	"os"
	"testing"
	"token-strike/internal/database"
)

func InitTempDatabase(t *testing.T) (*database.TokenStrikeDB, string) {
	path := tempfile()
	//defer os.RemoveAll(path)

	db, err := database.Connect(path)
	if err != nil {
		t.Fatal(err)
	} else if db == nil {
		t.Fatal("expected db")
	}

	if s := db.GetClient().Path(); s != path {
		t.Fatalf("unexpected path: %s", s)
	}

	return db, path
}

// tempfile returns a temporary file path.
func tempfile() string {
	f, err := ioutil.TempFile("", "bolt-")
	if err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
	return f.Name()
}

func CloseDB(db *database.TokenStrikeDB, t *testing.T) {
	err := db.Close()
	if err != nil {
		t.Error(err)
	}

	err = db.Clear()
	if err != nil {
		t.Error(err)
	}

}
