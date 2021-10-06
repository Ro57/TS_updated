package database

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"
	"time"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"

	"go.etcd.io/bbolt"
)

const (
	path = "~/.lnd/data/chain/pkt"
	name = "./test.db"
)

func TestCreateDB(t *testing.T) {
	ts, err := Connect(name)
	if err != nil {
		t.Fatalf("Connection refused %v", err)
	}

	defer ts.Close()

	err = ts.Ping()
	if err != nil {
		t.Fatal("db not connected: ", err)
	}
}

func TestPing(t *testing.T) {
	ts, err := Connect(name)
	if err != nil {
		t.Fatalf("Connection refused %v", err)
	}

	ts.Close()
	err = ts.Ping()

	if err == nil {
		t.Fatal("ping after close connection")
	}
}

func TestEmployeeUpdateView(t *testing.T) {
	ts, err := Connect(name)
	if err != nil {
		t.Fatalf("Connection refused %v", err)
	}

	defer ts.Close()

	defer func() {
		err = ts.Clear()
		if err != nil {
			t.Fatal("Clear failed", err)
		}
	}()

	err = ts.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Business"))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return nil
	})
	if err != nil {
		t.Fatal("Create top level busket structures failed: ", err)
	}

	err = ts.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Business"))

		_, err = b.CreateBucket([]byte("Employee"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return nil

	})
	if err != nil {
		t.Fatal("Create busket structures failed: ", err)
	}

	err = ts.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Business"))
		emp := b.Bucket([]byte("Employee"))

		err := b.Put([]byte("Emp1"), []byte("Number 100"))

		if err != nil {
			return err
		}

		err = emp.Put([]byte("Time"), []byte("180h"))

		return err
	})
	if err != nil {
		t.Fatal("Put information into bucket: ", err)
	}

	err = ts.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Business"))
		emp := b.Bucket([]byte("Employee"))

		empNum := b.Get([]byte("Emp1"))
		if string(empNum) != "Number 100" {
			t.Fatalf("want string 'Number 100' but get '%s'", string(empNum))
		}

		empTime := emp.Get([]byte("Time"))
		if string(empTime) != "180h" {
			t.Fatalf("want string '180h' but get '%s'", string(empTime))
		}

		return nil
	})
	if err != nil {
		t.Fatal("Get information from bucket: ", err)
	}
}

func TestTokenBlock(t *testing.T) {
	var wantExpBlockNumber int32 = 2000
	wantCreateTime := time.Now()
	tokenName := "smt"
	wantToken := DB.Token{
		Count:        200,
		Expiration:   wantExpBlockNumber,
		Creation:     wantCreateTime.UnixNano(),
		Urls:         []string{"https://some.token"},
		IssuerPubkey: "issuerPublicKey",
	}

	justificationPool := []*DB.Justification{}

	justificationPool = append(justificationPool,
		&DB.Justification{
			Content: &DB.Justification_Transfer{
				Transfer: &justifications.TranferToken{
					HtlcSecret: "some",
					Lock:       "some",
				},
			},
		},
	)

	wantBlock := DB.Block{
		Justifications: justificationPool,
		Signature:      "someSig",
		PrevBlock:      "hashPrevBlock",
		Creation:       time.Now().Unix(),
		State:          "hashOfState",
		PktBlockHash:   "hashFromPkt",
		PktBlockHeight: 1000,
		Height:         10,
	}

	var lastBlock [sha256.Size]byte

	ts, err := Connect(name)
	if err != nil {
		t.Fatalf("Connection refused %v", err)
	}

	defer ts.Close()

	defer func() {
		err := ts.Clear()
		if err != nil {
			t.Fatal("clear failed", err)
		}
	}()

	err = ts.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucket([]byte(tokenName))
		if err != nil {
			t.Fatalf("create top level bucket: %s", err)
		}

		chain, err := b.CreateBucket([]byte("chain"))
		if err != nil {
			t.Fatalf("create chain bucket: %s", err)
		}
		tokenByte, nativeErr := json.Marshal(wantToken)
		if nativeErr != nil {
			t.Fatalf("(update) marshal token structure: %s", nativeErr)
		}

		err = b.Put([]byte("Info"), tokenByte)
		if err != nil {
			t.Fatalf("put token %s", err)
		}

		blockByte, nativeErr := json.Marshal(wantBlock)
		if nativeErr != nil {
			t.Fatalf("(update) marshal block structure: %s", nativeErr)
		}

		lastBlock = sha256.Sum256(blockByte)

		err = chain.Put(lastBlock[:], blockByte)
		if err != nil {
			t.Fatalf("put block %s", err)
		}

		return nil
	})
	if err != nil {
		t.Fatal("generate DB structure: ", err)
	}

	err = ts.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(tokenName))
		if b == nil {
			t.Fatalf("bucket %s not found", tokenName)
		}

		chain := b.Bucket([]byte("chain"))
		if chain == nil {
			t.Fatal("bucket chain not found")
		}

		marshalToken, nativeErr := json.Marshal(wantToken)
		if nativeErr != nil {
			t.Fatalf("(view) marshal token structure: %s", nativeErr)
		}

		tokenByte := b.Get([]byte("Info"))
		if tokenByte == nil {
			t.Fatalf("token info with name %s not found", tokenByte)
		}
		if string(marshalToken) != string(tokenByte) {
			t.Fatalf("want token structure %s but get %s", marshalToken, tokenByte)
		}

		marshalBlock, nativeErr := json.Marshal(wantBlock)
		if nativeErr != nil {
			t.Fatalf("(view) marshal token structure: %s", nativeErr)
		}

		blockByte := chain.Get(lastBlock[:])
		if blockByte == nil {
			t.Fatalf("block with hash %s not found", string(lastBlock[:]))
		}
		if string(marshalBlock) != string(blockByte) {
			t.Fatalf("want block structure %s but get %s", marshalBlock, blockByte)
		}

		return nil
	})
	if err != nil {
		t.Fatal("read data from DB: ", err)
	}

}
