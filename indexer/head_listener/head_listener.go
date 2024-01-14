package head_listener

import (
	"context"

	v1 "github.com/attestantio/go-eth2-client/api/v1"
)

// HeadListener is used to listen to head events on the Ethereum beacon chain.
type HeadListener interface {
	// Listen starts the listening process in the background and populates the headEvents channel
	// with new head events as they arrive.
	Listen(ctx context.Context, headEvents chan<- *v1.HeadEvent) error
}

// AgateHeadListener is the implementation of the HeadListener for agate,
// using the BeaconAPIClient as a way to listen for new heads.
type AgateHeadListener struct {
	beaconAPIClient BeaconAPIClient
}

// NewAgateHeadListener creates an empty and non-initialized AgateHeadListener.
func NewAgateHeadListener() *AgateHeadListener {
	return &AgateHeadListener{
		beaconAPIClient: nil,
	}
}

// Init initializes a new AgateHeadListener using a BeaconAPIClient to listen to new head events.
func (listener *AgateHeadListener) Init(beaconAPIClient BeaconAPIClient) {
	listener.beaconAPIClient = beaconAPIClient
}

func (listener *AgateHeadListener) Listen(
	ctx context.Context,
	headEvents chan<- *v1.HeadEvent,
) error {
	if err := listener.beaconAPIClient.SubscribeToHeadEvents(ctx, func(event *v1.Event) {
		headEvent, ok := event.Data.(*v1.HeadEvent)
		if !ok {
			// TODO: Log error.
			return
		}

		headEvents <- headEvent
	}); err != nil {
		return NewHeadEventSubscriptionError(err)
	}

	return nil
}
