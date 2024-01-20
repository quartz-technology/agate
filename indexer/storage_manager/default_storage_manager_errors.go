package storage_manager

import "fmt"

// NewDefaultStorageManagerRelaysCreationError is raised if the DefaultStorageManager encounters an
// error while creating a relay entity in the database.
func NewDefaultStorageManagerRelaysCreationError(err error) error {
	return fmt.Errorf("default agate storage manager failed to create relays: %w", err)
}

// NewDefaultStorageManagerTransactionedRelaysCreationError is raised if the DefaultStorageManager
// encounters an error while executing the transaction responsible for storing multiple relays in
// the database.
func NewDefaultStorageManagerTransactionedRelaysCreationError(err error) error {
	return fmt.Errorf(
		"default agate storage manager failed to execute transaction which stores relays: %w",
		err,
	)
}

// NewDefaultStorageManagerBidsCreationError is raised if the DefaultStorageManager encounters an
// error while creating a list of bid entity in the database.
func NewDefaultStorageManagerBidsCreationError(err error) error {
	return fmt.Errorf("default agate storage manager failed to create bids: %w", err)
}

// NewDefaultStorageManagerSubmissionsCreationError is raised if the DefaultStorageManager encounters an
// error while creating a list of submission entity in the database.
func NewDefaultStorageManagerSubmissionsCreationError(err error) error {
	return fmt.Errorf(
		"default agate storage manager failed to create bids submissions: %w",
		err,
	)
}

// NewDefaultStorageManagerTransactionedBidsSubmissionsCreationError is raised if the
// DefaultStorageManager encounters an error while executing the transaction responsible for storing
// multiple bids and their submissions in the database.
func NewDefaultStorageManagerTransactionedBidsSubmissionsCreationError(err error) error {
	return fmt.Errorf(
		"default agate storage manager failed to execute transaction which stores bids and"+
			" their respective submissions: %w",
		err,
	)
}

func NewDefaultStorageManagerRelaysListingError(err error) error {
	return fmt.Errorf(
		"default agate store manager failed to list relays prior to storing aggregated relay data: %w",
		err,
	)
}
