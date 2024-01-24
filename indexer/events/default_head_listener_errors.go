package events

import "fmt"

// NewDefaultHeadListenerSubscriptionError is raised when the DefaultHeadListener fails to subscribe
// to new head events via its client.BeaconAPI.
func NewDefaultHeadListenerSubscriptionError(err error) error {
	return fmt.Errorf("default head listener failed to subscribe to head event: %w", err)
}
