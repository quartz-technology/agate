package client

import "fmt"

// NewDefaultRelayAPIServiceInitializationError is raised when the DefaultRelayAPI
// initialization has failed because the <redax> SDK could not be initialized.
func NewDefaultRelayAPIServiceInitializationError(err error) error {
	return fmt.Errorf("failed to initialize default relay API SDK: %w", err)
}

// NewDefaultRelayAPIBidsReceivedRetrievalError is raised if the request made to the relay to
// get the bids received for a specific slot has failed.
func NewDefaultRelayAPIBidsReceivedRetrievalError(err error) error {
	return fmt.Errorf("default relay API SDK failed to retrieve bids received by relay: %w", err)
}

// NewDefaultRelayAPIBidsDeliveredRetrievalError is raised if the request made to the relay to
// get the bids delivered for a specific slot has failed.
func NewDefaultRelayAPIBidsDeliveredRetrievalError(err error) error {
	return fmt.Errorf("default relay API SDK failed to retrieve bids delivered by relay: %w", err)
}
