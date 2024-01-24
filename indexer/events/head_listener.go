package events

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
