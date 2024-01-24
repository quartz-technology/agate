package postgres

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDefaultStore(t *testing.T) {
	t.Parallel()

	store := NewDefaultStore()

	require.NotNil(t, store)
	require.Nil(t, store.connectionPool)
}

func TestDefaultStore_Init(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultStore_ListRelays(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultStore_CreateRelays(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultStore_CreateBids(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultStore_CreateSubmissions(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultStore_ExecInTx(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultStore_Close(t *testing.T) {
	// TODO.
	t.Parallel()
}
