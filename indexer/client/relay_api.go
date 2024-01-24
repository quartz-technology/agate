package client

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/common"
)

// RelayAPI is used to collect data from a relay.
type RelayAPI interface {
	// GetRelayDataForSlot retrieves the bids received and delivered to a relay for a specific slot.
	GetRelayDataForSlot(ctx context.Context, slot phase0.Slot) (*common.RelayData, error)
	// GetRelayAPIURL returns the API URL of the relay this client makes request to.
	GetRelayAPIURL() string
}
