package store

import (
	"context"

	"github.com/quartz-technology/agate/indexer/storage_manager/store/models"
)

// Store is used to perform storage operations with a database.
// It should be used by the storage manager as the intermediate when interacting with the database.
type Store interface {
	// ExecInTx is used to perform a set of operations wrapped in a function in a single database
	// transaction. If one of them fails, the preceding operations all revert.
	//
	// Note that the database operations performed in the function callback should be made using
	// the store parameter of the said callback.
	ExecInTx(ctx context.Context, fn func(store Store) error) error

	// CreateRelays is used to create multiple relay entities in the database.
	CreateRelays(ctx context.Context, relays []*models.Relay) error
	// ListRelays is used to list all the relay entities stored in the database.
	ListRelays(ctx context.Context) ([]models.Relay, error)

	// CreateBids is used to create multiple bid entities in the database.
	CreateBids(ctx context.Context, bids []*models.Bid) error

	// CreateSubmissions is used to create multiple submission entities in the database.
	CreateSubmissions(ctx context.Context, submissions []*models.Submission) error

	// Close is used to shut down the connection from the Store to the database.
	Close()
}
