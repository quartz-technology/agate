package postgres

import "fmt"

// NewDefaultStoreDatabaseConnectionError is raised if the default agate store cannot connect to
// the postgres database.
func NewDefaultStoreDatabaseConnectionError(err error) error {
	return fmt.Errorf("default agate store failed to connect to postgres database: %w", err)
}

// NonInitializedTransactionerError is the error raised by the store when it attempts to perform
// a write operation on the database which should have been wrapped in a transaction - which is
// not initialized, therefore causing the error.
type NonInitializedTransactionerError struct {
	Method string
}

// NewNonInitializedTransactionerError creates an initialized NonInitializedTransactionerError
// given the name of the method attempting to perform the "unprotected" write operation in the
// database.
func NewNonInitializedTransactionerError(method string) error {
	return &NonInitializedTransactionerError{
		Method: method,
	}
}

func (e *NonInitializedTransactionerError) Error() string {
	return fmt.Sprintf("database transactioner is not initialized in Method %s", e.Method)
}

// NewDefaultStoreTransactionInitializationError is raised when the transaction object used to
// perform multiple operations in the database fails to be initialized.
func NewDefaultStoreTransactionInitializationError(err error) error {
	return fmt.Errorf("default agate store failed to initialize transactioner: %w", err)
}

// NewDefaultStoreBatchOperationError is raised when a batch operation fails to be performed.
func NewDefaultStoreBatchOperationError(err error) error {
	return fmt.Errorf(
		"default agate store failed to perform the batch (bulk) operation in the database: %w",
		err,
	)
}

// NewDefaultStoreRelayListingError is raised when the relays stored in the database can not be
// listed.
func NewDefaultStoreRelayListingError(err error) error {
	return fmt.Errorf("default agate store failed to list relays: %w", err)
}

// NewDefaultStoreRelayListDecodingError is raised when the received relay list fails to be
// decoded in a go object.
func NewDefaultStoreRelayListDecodingError(err error) error {
	return fmt.Errorf("default agate store failed to decode relay list: %w", err)
}

// NewDefaultStoreTransactionRollbackError is raised when a rollback triggered by an error which occurred
// during a transaction execution fails to complete.
func NewDefaultStoreTransactionRollbackError(err error) error {
	return fmt.Errorf("default agate store failed to rollback transaction: %w", err)
}

// NewDefaultStoreTransactionExecutionError is raised when one of the sub-operation of a
// transaction has failed.
func NewDefaultStoreTransactionExecutionError(err error) error {
	return fmt.Errorf(
		"default agate store failed to execute an operation of a transaction: %w",
		err,
	)
}

// NewDefaultStoreTransactionCommitError is raised when a transaction failed to be committed to
// the database.
func NewDefaultStoreTransactionCommitError(err error) error {
	return fmt.Errorf("default agate store failed to commit transaction: %w", err)
}

func NewAgateStoreCloseError(err error) error {
	return fmt.Errorf("failed to close postgres database connection pool: %w", err)
}
