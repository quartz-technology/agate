package data

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/common"
)

// Aggregator is used to aggregate data from multiple relays.
type Aggregator interface {
	// AggregateDataForSlotFromRelays collects data from multiple relays for a specific slot.
	AggregateDataForSlotFromRelays(
		ctx context.Context,
		slot phase0.Slot,
	) (*common.AggregatedRelayData, error)
}
