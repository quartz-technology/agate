package storage_manager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDefaultDatabaseMigrator(t *testing.T) {
	t.Parallel()

	migrator := NewDefaultDatabaseMigrator()

	require.NotNil(t, migrator)
	require.Nil(t, migrator.clt)
}

func TestDefaultDatabaseMigrator_Init(t *testing.T) {
	// TODO.
	t.Parallel()
}

func TestDefaultDatabaseMigrator_Migrate(t *testing.T) {
	// TODO.
	t.Parallel()
}
