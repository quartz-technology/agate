package client

import (
	"errors"
	"fmt"
)

var ErrDefaultBeaconAPIServiceTypeAssertion = errors.New("failed to type assert beacon API HTTP service")

// NewDefaultBeaconAPIServiceInitializationError is raised if the client used to connect to the beacon
// node fails to initialize.
func NewDefaultBeaconAPIServiceInitializationError(err error) error {
	return fmt.Errorf("failed to initialize default beacon API HTTP service: %w", err)
}

// NewDefaultBeaconAPIEventSubscriptionError is raised if the client connected to the node fails to create an event
// subscription.
func NewDefaultBeaconAPIEventSubscriptionError(err error) error {
	return fmt.Errorf("default beacon API HTTP service failed to create event subscription: %w", err)
}
