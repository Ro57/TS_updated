package repository

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"testing"
	"time"

	"token-strike/internal/database"
	"token-strike/internal/errors"
	"token-strike/tsp2p/server/DB"

	"github.com/stretchr/testify/require"
)

func TestBbolt_GetChainInfoDB(t *testing.T) {
	path := tempfile()
	defer os.RemoveAll(path)

	db, err := database.Connect(path)
	if err != nil {
		t.Fatalf(err.Error())
	}

	type args struct {
		tokenId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		db      *database.TokenStrikeDB
		initDB  func(database.DBRepository)
	}{
		{
			name: "ValidTest",
			args: args{
				tokenId: "test",
			},
			wantErr: nil,
			db:      db,
			initDB:  validInitDB,
		},
		{
			name: "TokenNotFound",
			args: args{
				tokenId: "tokenNotFound",
			},
			wantErr: errors.TokensDBNotFound,
			db:      db,
			initDB:  withoutInitDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bbolt{
				db: tt.db,
			}
			tt.initDB(b)
			res, err := b.GetChainInfoDB(tt.args.tokenId)

			require.Equal(
				t,
				tt.wantErr,
				err,
				fmt.Sprintf("GetChainInfoDB() error = %v, wantErr %v", err, tt.wantErr),
			)
			fmt.Println(res)
		})
	}
}

func withoutInitDB(db database.DBRepository) {

}

func validInitDB(db database.DBRepository) {
	token := DB.Token{
		Count:        10,
		Expiration:   math.MaxInt32,
		Creation:     time.Now().Unix(),
		IssuerPubkey: "issuer_pub_key",
		Urls: []string{
			"http://localhost:3333/token1",
		},
	}

	state := &DB.State{
		Token: &token,
		Owners: []*DB.Owner{
			{
				HolderWallet: "alice",
				Count:        6,
			},
			{
				HolderWallet: "bob",
				Count:        4,
			},
		},
		Locks: nil,
	}

	stateBytes, err := state.GetHash()
	if err != nil {
		panic(err)
	}

	block := &DB.Block{
		PrevBlock:      "",
		Justifications: nil,
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
		PktBlockHash:   "some hash",
		PktBlockHeight: 10000,
		Height:         10010,
		Signature:      "some signature",
	}

	err = db.IssueTokenDB("test", &token, block, state)
	if err != nil {
		panic(err.Error())
	}
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
