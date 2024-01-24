package client

import (
	"context"

	client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/rs/zerolog"
)

// DefaultBeaconAPI is the default implementation of the BeaconAPI for agate,
// using the go-eth2-client library to send queries to the beacon node.
type DefaultBeaconAPI struct {
	clt *http.Service
}

// NewDefaultBeaconAPI creates a new and non-initialized DefaultBeaconAPI.
func NewDefaultBeaconAPI() *DefaultBeaconAPI {
	return &DefaultBeaconAPI{
		clt: nil,
	}
}

// Init initializes an DefaultBeaconAPI using the beacon API URL to connect to the beacon node.
func (client *DefaultBeaconAPI) Init(ctx context.Context, beaconAPIURL string) error {
	var ok bool

	clt, err := http.New(
		ctx,
		http.WithAddress(beaconAPIURL),
		http.WithLogLevel(zerolog.Disabled),
	)
	if err != nil {
		return NewDefaultBeaconAPIServiceInitializationError(err)
	}

	client.clt, ok = clt.(*http.Service)
	if !ok {
		return ErrDefaultBeaconAPIServiceTypeAssertion
	}

	return nil
}

func (client *DefaultBeaconAPI) SubscribeToHeadEvents(
	ctx context.Context,
	handler client.EventHandlerFunc,
) error {
	if err := client.clt.Events(ctx, []string{"head"}, handler); err != nil {
		return NewDefaultBeaconAPIEventSubscriptionError(err)
	}

	return nil
}
