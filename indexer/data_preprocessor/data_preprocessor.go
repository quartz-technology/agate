package data_preprocessor

import (
	"github.com/quartz-technology/agate/indexer/common"
)

// The DataPreprocessor is used to transform the raw data acquired by the data aggregator
// service, which is then used by the storage manager to save it in a database.
type DataPreprocessor[T any] interface {
	// Preprocess transforms the data_aggregator.DataAggregator's aggregation output into a data
	// structure that the storage_manager.StorageManager service can store in the database.
	Preprocess(data *common.AggregatedRelayData) *DataPreprocessorOutput[T]
}
