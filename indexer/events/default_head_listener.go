package events

import (
	"context"

	v1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/quartz-technology/agate/indexer/client"
)

// DefaultHeadListener is the implementation of the HeadListener for agate,
// using the client.BeaconAPI as a way to listen for new heads.
type DefaultHeadListener struct {
	beaconAPIClient client.BeaconAPI
}

// NewDefaultHeadListener creates an empty and non-initialized DefaultHeadListener.
func NewDefaultHeadListener() *DefaultHeadListener {
	return &DefaultHeadListener{
		beaconAPIClient: nil,
	}
}

// Init initializes a new DefaultHeadListener using a client.BeaconAPI to listen to new head events.
func (listener *DefaultHeadListener) Init(beaconAPIClient client.BeaconAPI) {
	listener.beaconAPIClient = beaconAPIClient
}

func (listener *DefaultHeadListener) Listen(
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
		return NewDefaultHeadListenerSubscriptionError(err)
	}

	return nil
}
