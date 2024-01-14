package head_listener

import (
	"context"

	client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/rs/zerolog"
)

// BeaconAPIClient is used to interact with a beacon node.
type BeaconAPIClient interface {
	// SubscribeToHeadEvents is used to create a subscription to new head events and will perform
	// the handler's logic.
	SubscribeToHeadEvents(ctx context.Context, handler client.EventHandlerFunc) error
}

// AgateBeaconAPIClient is the implementation of the BeaconAPIClient for agate,
// using the go-eth2-client library to interact with the beacon node.
type AgateBeaconAPIClient struct {
	clt *http.Service
}

// NewAgateBeaconAPIClient creates a new and non-initialized AgateBeaconAPIClient.
func NewAgateBeaconAPIClient() *AgateBeaconAPIClient {
	return &AgateBeaconAPIClient{
		clt: nil,
	}
}

// Init initializes an AgateBeaconAPIClient using the beacon API URL to connect to the beacon node.
func (client *AgateBeaconAPIClient) Init(ctx context.Context, beaconAPIURL string) error {
	clt, err := http.New(
		ctx,
		http.WithAddress(beaconAPIURL),
		http.WithLogLevel(zerolog.Disabled),
	)
	if err != nil {
		return NewBeaconAPIClientInitializationError(err)
	}

	client.clt = clt.(*http.Service)

	return nil
}

func (client *AgateBeaconAPIClient) SubscribeToHeadEvents(
	ctx context.Context,
	handler client.EventHandlerFunc,
) error {
	if err := client.clt.Events(ctx, []string{"head"}, handler); err != nil {
		return NewEventSubscriptionError(err)
	}

	return nil
}
