package storage_manager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDefaultStorageManager(t *testing.T) {
	t.Parallel()

	manager := NewDefaultStorageManager()

	require.NotNil(t, manager)
	require.Nil(t, manager.store)
}

func TestDefaultStorageManager_Init(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultStorageManager_StoreRelays(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultStorageManager_StoreAggregatedRelayData(t *testing.T) {
	// TODO.
	t.Parallel()
}
