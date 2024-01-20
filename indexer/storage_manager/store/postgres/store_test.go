package postgres

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDefaultStore(t *testing.T) {
	store := NewDefaultStore()

	require.NotNil(t, store)
	require.Nil(t, store.connectionPool)
}

func TestDefaultStore_Init(t *testing.T) {
	// TODO.
}

func TestDefaultStore_ListRelays(t *testing.T) {
	// TODO.
}

func TestDefaultStore_CreateRelays(t *testing.T) {
	// TODO.
}

func TestDefaultStore_CreateBids(t *testing.T) {
	// TODO.
}

func TestDefaultStore_CreateSubmissions(t *testing.T) {
	// TODO.
}

func TestDefaultStore_ExecInTx(t *testing.T) {
	// TODO.
}

func TestDefaultStore_Close(t *testing.T) {
	// TODO.
}
