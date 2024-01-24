package client

import (
	"context"

	client "github.com/attestantio/go-eth2-client"
)

// BeaconAPI is used to interact with a beacon node.
type BeaconAPI interface {
	// SubscribeToHeadEvents is used to create a subscription to new head events and will perform
	// the handler's logic.
	SubscribeToHeadEvents(ctx context.Context, handler client.EventHandlerFunc) error
}
