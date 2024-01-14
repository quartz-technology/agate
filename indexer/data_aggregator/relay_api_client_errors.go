package data_aggregator

import "fmt"

// NewAgateRelayAPIClientSDKInitializationError is raised when the AgateRelayAPIClient
// initialization has failed because the <redax> SDK could not be initialized.
func NewAgateRelayAPIClientSDKInitializationError(err error) error {
	return fmt.Errorf("failed to initialize relay SDK: %w", err)
}

// NewAgateRelayAPIClientBidsReceivedRetrievalError is raised if the request made to the relay to
// get the bids received for a specific slot has failed.
func NewAgateRelayAPIClientBidsReceivedRetrievalError(err error) error {
	return fmt.Errorf("failed to retrieve bids received by relay: %w", err)
}

// NewAgateRelayAPIClientBidsDeliveredRetrievalError is raised if the request made to the relay to
// get the bids delivered for a specific slot has failed.
func NewAgateRelayAPIClientBidsDeliveredRetrievalError(err error) error {
	return fmt.Errorf("failed to retrieve bids delivered by relay: %w", err)
}
