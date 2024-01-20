package storage_manager

import (
	"context"

	"github.com/quartz-technology/agate/indexer/data_preprocessor"
	"github.com/quartz-technology/agate/indexer/storage_manager/store/dto"
)

// The StorageManager is used to store preprocessed aggregated relay data into a database.
type StorageManager[T any] interface {
	// StoreRelays is used to store the relay entities at Agate's startup.
	StoreRelays(ctx context.Context, relays []*dto.Relay) error
	// StoreAggregatedRelayData is used to store the preprocessed aggregated relay data once
	// received by the data aggregator.
	StoreAggregatedRelayData(
		ctx context.Context,
		data *data_preprocessor.DataPreprocessorOutput[T],
	) error
}
