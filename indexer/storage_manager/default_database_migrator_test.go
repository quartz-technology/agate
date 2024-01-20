package storage_manager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDefaultDatabaseMigrator(t *testing.T) {
	migrator := NewDefaultDatabaseMigrator()

	require.NotNil(t, migrator)
	require.Nil(t, migrator.clt)
}

func TestDefaultDatabaseMigrator_Init(t *testing.T) {
	// TODO.
}

func TestDefaultDatabaseMigrator_Migrate(t *testing.T) {
	// TODO.
}
