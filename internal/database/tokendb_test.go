package database

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"
	"time"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.etcd.io/bbolt"
)

const (
	path = "~/.lnd/data/chain/pkt"
	name = "./test.db"
)

const (
	TypeUpdate = "update"
	TypeView   = "view"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	db *TokenStrikeDB
}

func (suite *TestSuite) SetupTest() {
	db, err := Connect(name)
	if err != nil {
		suite.T().Fatalf("Connection refused %v", err)
	}

	suite.db = db
}

func (suite *TestSuite) AfterTest(suiteName, testName string) {
	err := suite.db.Clear()
	if err != nil {
		suite.T().Fatalf("Clear test file exception %v", err)
	}
	suite.db.Close()
}

func (suite *TestSuite) selectType(txType string) func(func(*bbolt.Tx) error) error {
	switch txType {
	case TypeUpdate:
		return suite.db.Update
	case TypeView:
		return suite.db.View
	}
	return nil
}

func (suite *TestSuite) TestConectDB() {
	suite.T().Run("create DB", func(t *testing.T) {
		err := suite.db.Ping()
		require.NoError(t, err, "db not connected: ")
	})

	suite.T().Run("close DB ", func(t *testing.T) {
		suite.db.Close()
		err := suite.db.Ping()
		require.Error(t, err, "ping after close connection")
	})
}

func (suite *TestSuite) TestEmployee() {
	testTransactions := []struct {
		name   string
		txType string
		tx     func(*bbolt.Tx) error
	}{
		{
			name:   "create business",
			txType: TypeUpdate,
			tx: func(tx *bbolt.Tx) error {
				_, err := tx.CreateBucket([]byte("Business"))

				if err != nil {
					return fmt.Errorf("create bucket: %s", err)
				}

				return nil
			},
		},
		{
			name:   "create employee bucket",
			txType: TypeUpdate,
			tx: func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("Business"))

				_, err := b.CreateBucket([]byte("Employee"))
				if err != nil {
					return fmt.Errorf("create bucket: %s", err)
				}

				return nil

			},
		},
		{
			name:   "create emp1",
			txType: TypeUpdate,
			tx: func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("Business"))
				emp := b.Bucket([]byte("Employee"))

				err := b.Put([]byte("Emp1"), []byte("Number 100"))

				if err != nil {
					return err
				}

				err = emp.Put([]byte("Time"), []byte("180h"))

				return err

			},
		},
		{
			name:   "get emp1 from db",
			txType: TypeView,
			tx: func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("Business"))
				emp := b.Bucket([]byte("Employee"))

				empNum := b.Get([]byte("Emp1"))
				if string(empNum) != "Number 100" {
					return fmt.Errorf("want string 'Number 100' but get '%s'", string(empNum))
				}

				empTime := emp.Get([]byte("Time"))
				if string(empTime) != "180h" {
					return fmt.Errorf("want string '180h' but get '%s'", string(empTime))
				}

				return nil
			},
		},
	}

	for _, tt := range testTransactions {
		suite.T().Run(tt.name, func(t *testing.T) {
			err := suite.selectType(tt.txType)(tt.tx)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func (suite *TestSuite) TestTokenBlock() {
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

	testTransactions := []struct {
		name   string
		txType string
		tx     func(*bbolt.Tx) error
	}{
		{
			name:   "create blockchain",
			txType: TypeUpdate,
			tx: func(tx *bbolt.Tx) error {
				b, err := tx.CreateBucket([]byte(tokenName))
				if err != nil {
					return fmt.Errorf("create top level bucket: %s", err)
				}

				chain, err := b.CreateBucket([]byte("chain"))
				if err != nil {
					return fmt.Errorf("create chain bucket: %s", err)
				}
				tokenByte, nativeErr := json.Marshal(wantToken)
				if nativeErr != nil {
					return fmt.Errorf("(update) marshal token structure: %s", nativeErr)
				}

				err = b.Put([]byte("Info"), tokenByte)
				if err != nil {
					return fmt.Errorf("put token %s", err)
				}

				blockByte, nativeErr := json.Marshal(wantBlock)
				if nativeErr != nil {
					return fmt.Errorf("(update) marshal block structure: %s", nativeErr)
				}

				lastBlock = sha256.Sum256(blockByte)

				err = chain.Put(lastBlock[:], blockByte)
				if err != nil {
					return fmt.Errorf("put block %s", err)
				}

				return nil
			},
		},
		{
			name:   "get data from blockchain",
			txType: TypeView,
			tx: func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte(tokenName))
				if b == nil {
					return fmt.Errorf("bucket %s not found", tokenName)
				}

				chain := b.Bucket([]byte("chain"))
				if chain == nil {
					return fmt.Errorf("bucket chain not found")
				}

				marshalToken, nativeErr := json.Marshal(wantToken)
				if nativeErr != nil {
					return fmt.Errorf("(view) marshal token structure: %s", nativeErr)
				}

				tokenByte := b.Get([]byte("Info"))
				if tokenByte == nil {
					return fmt.Errorf("token info with name %s not found", tokenByte)
				}
				if string(marshalToken) != string(tokenByte) {
					return fmt.Errorf("want token structure %s but get %s", marshalToken, tokenByte)
				}

				marshalBlock, nativeErr := json.Marshal(wantBlock)
				if nativeErr != nil {
					return fmt.Errorf("(view) marshal token structure: %s", nativeErr)
				}

				blockByte := chain.Get(lastBlock[:])
				if blockByte == nil {
					return fmt.Errorf("block with hash %s not found", string(lastBlock[:]))
				}
				if string(marshalBlock) != string(blockByte) {
					return fmt.Errorf("want block structure %s but get %s", marshalBlock, blockByte)
				}

				return nil
			},
		},
	}

	for _, tt := range testTransactions {
		suite.T().Run(tt.name, func(t *testing.T) {
			err := suite.selectType(tt.txType)(tt.tx)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
