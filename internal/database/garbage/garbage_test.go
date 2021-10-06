package garbage

import (
	"os"
	"testing"
	"token-strike/internal/database"

	"github.com/stretchr/testify/require"
)

const (
	path = "/.lnd/data/chain/pkt"
	name = "/garbage_test.db"
)

func TestGarbage(t *testing.T) {
	home, err := os.UserHomeDir()
	require.NoError(t, err)

	db, err := database.Connect(home + path + name)
	require.NoError(t, err)
	defer db.Close()

	Generate(db)

}
