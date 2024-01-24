//nolint:stylecheck
package storage

// The DatabaseMigrator is used to apply the database migrations, ideally at the indexer's startup.
type DatabaseMigrator interface {
	// Migrate is used to migrate the schemas in db/migrations.
	Migrate() error
}
