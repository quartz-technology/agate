package storage_manager

import "fmt"

// NewDefaultDatabaseMigratorInitializationError is raised if the DefaultDatabaseMigrator encounters
// an error during the initialization process.
func NewDefaultDatabaseMigratorInitializationError(err error) error {
	return fmt.Errorf("failed to initialize agate database migrator: %w", err)
}

// NewDefaultDatabaseMigratorMigrationError is raised when the DefaultDatabaseMigrator fails to
// apply the database migrations.
func NewDefaultDatabaseMigratorMigrationError(err error) error {
	return fmt.Errorf("agate database migrator failed to apply migrations: %w", err)
}
