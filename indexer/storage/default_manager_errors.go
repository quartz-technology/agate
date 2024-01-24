package storage

import "fmt"

// NewDefaultManagerRelaysCreationError is raised if the DefaultManager encounters an
// error while creating a relay entity in the database.
func NewDefaultManagerRelaysCreationError(err error) error {
	return fmt.Errorf("default agate storage manager failed to create relays: %w", err)
}

// NewDefaultManagerTransactionedRelaysCreationError is raised if the DefaultManager
// encounters an error while executing the transaction responsible for storing multiple relays in
// the database.
func NewDefaultManagerTransactionedRelaysCreationError(err error) error {
	return fmt.Errorf(
		"default agate storage manager failed to execute transaction which stores relays: %w",
		err,
	)
}

// NewDefaultManagerBidsCreationError is raised if the DefaultManager encounters an
// error while creating a list of bid entity in the database.
func NewDefaultManagerBidsCreationError(err error) error {
	return fmt.Errorf("default agate storage manager failed to create bids: %w", err)
}

// NewDefaultManagerSubmissionsCreationError is raised if the DefaultManager encounters an
// error while creating a list of submission entity in the database.
func NewDefaultManagerSubmissionsCreationError(err error) error {
	return fmt.Errorf(
		"default agate storage manager failed to create bids submissions: %w",
		err,
	)
}

// NewDefaultManagerTransactionedBidsSubmissionsCreationError is raised if the
// DefaultManager encounters an error while executing the transaction responsible for storing
// multiple bids and their submissions in the database.
func NewDefaultManagerTransactionedBidsSubmissionsCreationError(err error) error {
	return fmt.Errorf(
		"default agate storage manager failed to execute transaction which stores bids and"+
			" their respective submissions: %w",
		err,
	)
}

func NewDefaultManagerRelaysListingError(err error) error {
	return fmt.Errorf(
		"default agate storage manager failed to list relays prior to storing aggregated relay data: %w",
		err,
	)
}
