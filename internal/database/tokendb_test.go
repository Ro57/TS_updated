package database

import (
	"crypto/sha256"
	"encoding/json"
	"os"
	"testing"
	"time"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"

	"github.com/stretchr/testify/suite"
	"go.etcd.io/bbolt"
)

const (
	path = "/.lnd/data/chain/pkt"
	name = "/test.db"
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
	home, err := os.UserHomeDir()
	suite.NoError(err)

	db, err := Connect(home + path + name)
	suite.NoError(err)

	suite.db = db
}

func (suite *TestSuite) AfterTest(suiteName, testName string) {
	err := suite.db.Clear()
	suite.NoError(err, "Clear test file exception")

	suite.db.Close()
}

// selectType return function by type of transaction
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
		suite.NoError(err, "db not connected: ")
	})

	suite.T().Run("close DB ", func(t *testing.T) {
		suite.db.Close()
		err := suite.db.Ping()
		suite.Error(err, "ping after close connection")
	})
}

func (suite *TestSuite) TestEmployee() {
	wantTime := "180h"
	wantNumber := "Number 100"

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
				suite.NoError(err, "create bucket")

				return nil
			},
		},
		{
			name:   "create employee bucket",
			txType: TypeUpdate,
			tx: func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("Business"))

				_, err := b.CreateBucket([]byte("Employee"))
				suite.NoError(err, "create bucket: %s", err)

				return nil

			},
		},
		{
			name:   "create emp1",
			txType: TypeUpdate,
			tx: func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("Business"))
				emp := b.Bucket([]byte("Employee"))

				err := b.Put([]byte("Emp1"), []byte(wantNumber))
				suite.NoError(err)

				err = emp.Put([]byte("Time"), []byte(wantTime))

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
				suite.Equal(string(empNum), wantNumber, "want number '%s' but get '%s'", wantNumber, empNum)

				empTime := emp.Get([]byte("Time"))
				suite.Equal(string(empTime), wantTime, "want time '%s' but get '%s'", wantTime, empTime)

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
				suite.NoError(err, "create top level bucket")

				chain, err := b.CreateBucket([]byte("chain"))
				suite.NoError(err, "create chain bucket")

				tokenByte, err := json.Marshal(wantToken)
				suite.NoError(err, "(update) marshal token structure")

				err = b.Put([]byte("Info"), tokenByte)
				suite.NoError(err, "put token")

				blockByte, err := json.Marshal(wantBlock)
				suite.NoError(err, "(update) marshal block structure")

				lastBlock = sha256.Sum256(blockByte)

				err = chain.Put(lastBlock[:], blockByte)
				suite.NoError(err, "put block ")

				return nil
			},
		},
		{
			name:   "get data from blockchain",
			txType: TypeView,
			tx: func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte(tokenName))
				suite.NotNil(b, "bucket %s not found", tokenName)

				chain := b.Bucket([]byte("chain"))
				suite.NotNil(b, "bucket chain not found")

				marshalToken, err := json.Marshal(wantToken)
				suite.NoError(err, "(view) marshal token structure")

				tokenByte := b.Get([]byte("Info"))
				suite.NotNil(tokenByte, "token info with name %s not found", tokenName)

				suite.Equal(
					string(marshalToken),
					string(tokenByte),
					"want token structure %s but get %s",
					marshalToken,
					tokenByte,
				)

				marshalBlock, err := json.Marshal(wantBlock)
				suite.NoError(err, "(view) marshal token structure")

				blockByte := chain.Get(lastBlock[:])
				suite.NotNil(blockByte, "block with hash %s not found", string(lastBlock[:]))

				suite.Equal(
					string(marshalBlock),
					string(blockByte),
					"want block structure %s but get %s",
					marshalBlock,
					blockByte,
				)

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
