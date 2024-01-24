package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDefaultManager(t *testing.T) {
	t.Parallel()

	manager := NewDefaultManager()

	require.NotNil(t, manager)
	require.Nil(t, manager.store)
}

func TestDefaultManager_Init(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultManager_StoreRelays(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultManager_StoreAggregatedRelayData(t *testing.T) {
	// TODO.
	t.Parallel()
}
