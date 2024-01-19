package storage_manager

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
)

// The DatabaseMigrator is used to apply the database migrations, ideally at the indexer's startup.
type DatabaseMigrator interface {
	// Migrate is used to migrate the schemas in db/migrations.
	Migrate() error
}

// The DefaultDatabaseMigrator is an implementation of the DatabaseMigrator which applies the
// migrations to a Postgres database.
type DefaultDatabaseMigrator struct {
	clt *migrate.Migrate
}

// NewDefaultDatabaseMigrator creates an empty and non-initialized DefaultDatabaseMigrator.
func NewDefaultDatabaseMigrator() *DefaultDatabaseMigrator {
	return &DefaultDatabaseMigrator{
		clt: nil,
	}
}

// Init initializes a DefaultDatabaseMigrator given the migrations source and the database URL.
func (migrator *DefaultDatabaseMigrator) Init(sourceURL, databaseURL string) error {
	clt, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return NewDefaultDatabaseMigratorInitializationError(err)
	}

	migrator.clt = clt

	return nil
}

func (migrator *DefaultDatabaseMigrator) Migrate() error {
	if err := migrator.clt.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}

		return NewDefaultDatabaseMigratorMigrationError(err)
	}

	return nil
}
