package storage

import (
	"context"

	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/agate/indexer/storage/store/dto"
)

// The Manager is used to store preprocessed aggregated relay data into a database.
type Manager interface {
	// StoreRelays is used to store the relay entities at Agate's startup.
	StoreRelays(ctx context.Context, relays []*dto.Relay) error
	// StoreAggregatedRelayData is used to store the preprocessed aggregated relay data once
	// received by the data aggregator.
	StoreAggregatedRelayData(ctx context.Context, data *common.DataPreprocessorOutput) error
	// Shutdown stops the manager sub-services.
	Shutdown()
}
