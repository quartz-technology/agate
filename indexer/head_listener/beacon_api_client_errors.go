package head_listener

import (
	"errors"
	"fmt"
)

var ErrBeaconAPIClientTypeAssertionError = errors.New("failed to type assert beacon API HTTP client")

// NewBeaconAPIClientInitializationError is raised if the client used to connect to the beacon
// node fails to initialize.
func NewBeaconAPIClientInitializationError(err error) error {
	return fmt.Errorf("failed to initialize beacon API client service: %w", err)
}

// NewEventSubscriptionError is raised if the client connected to the node fails to create an event
// subscription.
func NewEventSubscriptionError(err error) error {
	return fmt.Errorf("failed to create event subscription: %w", err)
}
