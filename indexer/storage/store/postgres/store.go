package postgres

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quartz-technology/agate/indexer/storage/store"
	"github.com/quartz-technology/agate/indexer/storage/store/models"
)

// The DefaultStore is an implementation of the Store which connects to a Postgres database to
// store the relay data.
type DefaultStore struct {
	connectionPool *pgxpool.Pool
	tx             pgx.Tx
}

// NewDefaultStore creates an empty and non-initialized DefaultStore.
func NewDefaultStore() *DefaultStore {
	return &DefaultStore{
		connectionPool: nil,
		tx:             nil,
	}
}

// Init initializes a DefaultStore given the postgres database URL.
func (store *DefaultStore) Init(ctx context.Context, databaseURL string) error {
	connectionPool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return NewDefaultStoreDatabaseConnectionError(err)
	}

	store.connectionPool = connectionPool

	return nil
}

func (store *DefaultStore) CreateRelays(ctx context.Context, relays []*models.Relay) error {
	tx := store.tx

	if tx == nil {
		return NewNonInitializedTransactionerError("CreateRelays")
	}

	relayBulkInsertQuery := `
	INSERT INTO relays (url)
	VALUES (@url)
	ON CONFLICT DO NOTHING
`
	relayBatch := &pgx.Batch{}

	for _, relay := range relays {
		args := pgx.NamedArgs{
			"url": relay.URL,
		}

		relayBatch.Queue(relayBulkInsertQuery, args)
	}

	br := tx.SendBatch(ctx, relayBatch)

	if err := br.Close(); err != nil {
		return NewDefaultStoreBatchOperationError(err)
	}

	return nil
}

func (store *DefaultStore) ListRelays(ctx context.Context) ([]models.Relay, error) {
	rows, err := store.connectionPool.Query(ctx, "SELECT id, url FROM relays")
	if err != nil {
		return nil, NewDefaultStoreRelayListingError(err)
	}

	relays, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Relay])
	if err != nil {
		return nil, NewDefaultStoreRelayListDecodingError(err)
	}

	return relays, nil
}

func (store *DefaultStore) ExecInTx(ctx context.Context, fn func(store store.Store) error) error {
	//nolint:exhaustruct
	tx, err := store.connectionPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return NewDefaultStoreTransactionInitializationError(err)
	}

	txStore := &DefaultStore{
		connectionPool: store.connectionPool,
		tx:             tx,
	}

	if err := fn(txStore); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return NewDefaultStoreTransactionRollbackError(rbErr)
		}

		return NewDefaultStoreTransactionExecutionError(err)
	}

	if err := tx.Commit(ctx); err != nil {
		return NewDefaultStoreTransactionCommitError(err)
	}

	return nil
}

func (store *DefaultStore) CreateBids(ctx context.Context, bids []*models.Bid) error {
	tx := store.tx

	if tx == nil {
		return NewNonInitializedTransactionerError("CreateBids")
	}

	bidBulkInsertQuery := `
	INSERT INTO bids (slot, parent_hash, block_hash, fee_recipient, gas_limit, gas_used, value, num_tx, proposer, builder)
	VALUES (@slot, @parent_hash, @block_hash, @fee_recipient, @gas_limit, @gas_used, @value, @num_tx, @proposer, @builder)
	ON CONFLICT DO NOTHING
`
	bidBatch := &pgx.Batch{}

	for _, bid := range bids {
		args := pgx.NamedArgs{
			"slot":          uint64(bid.Slot),
			"parent_hash":   bid.ParentHash,
			"block_hash":    bid.BlockHash,
			"fee_recipient": bid.FeeRecipient,
			"gas_limit":     bid.GasLimit,
			"gas_used":      bid.GasUsed,
			"value":         bid.Value.Bytes(),
			"num_tx":        bid.NumTx,
			"proposer":      bid.Proposer,
			"builder":       bid.Builder,
		}

		bidBatch.Queue(bidBulkInsertQuery, args)
	}

	if err := tx.SendBatch(ctx, bidBatch).Close(); err != nil {
		return NewDefaultStoreBatchOperationError(err)
	}

	return nil
}

func (store *DefaultStore) CreateSubmissions(ctx context.Context, submissions []*models.Submission) error {
	tx := store.tx

	if tx == nil {
		return NewNonInitializedTransactionerError("CreateSubmissions")
	}

	submissionBulkInsertQuery := `
	INSERT INTO submissions (relay_id, bid_block_hash, is_delivered, is_optimistic, submitted_at)
	VALUES (@relay_id, @bid_block_hash, @is_delivered, @is_optimistic, @submitted_at)
`
	submissionBatch := &pgx.Batch{}

	for _, submission := range submissions {
		args := pgx.NamedArgs{
			"relay_id":       submission.RelayID,
			"bid_block_hash": submission.BidBlockHash,
			"is_delivered":   submission.IsDelivered,
			"is_optimistic":  submission.IsOptimistic,
			"submitted_at":   submission.SubmittedAt,
		}

		submissionBatch.Queue(submissionBulkInsertQuery, args)
	}

	if err := tx.SendBatch(ctx, submissionBatch).Close(); err != nil {
		return NewDefaultStoreBatchOperationError(err)
	}

	return nil
}

func (store *DefaultStore) Close() {
	store.connectionPool.Close()
}
