package indexer

import "fmt"

// NewHeadEventSubscriptionError is raised when the HeadListener fails to subscribe to new head
// events via its BeaconAPIClient.
func NewHeadEventSubscriptionError(err error) error {
	return fmt.Errorf("failed to subscribe to head event: %w", err)
}
