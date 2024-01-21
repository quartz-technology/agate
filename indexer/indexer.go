package indexer

import (
	"context"

	v1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/quartz-technology/agate/indexer/data_aggregator"
	"github.com/quartz-technology/agate/indexer/data_preprocessor"
	"github.com/quartz-technology/agate/indexer/head_listener"
	"github.com/quartz-technology/agate/indexer/storage_manager"
	"github.com/rs/zerolog/log"
)

type Indexer struct {
	listener     head_listener.HeadListener
	aggregator   data_aggregator.DataAggregator
	preprocessor *data_preprocessor.DataPreprocessor
	storage      storage_manager.StorageManager
}

func NewIndexer(listener head_listener.HeadListener,
	aggregator data_aggregator.DataAggregator,
	preprocessor *data_preprocessor.DataPreprocessor,
	storage storage_manager.StorageManager,
) *Indexer {
	return &Indexer{
		listener:     listener,
		aggregator:   aggregator,
		preprocessor: preprocessor,
		storage:      storage,
	}
}

func (indexer *Indexer) Start(ctx context.Context) error {
	headEvents := make(chan *v1.HeadEvent)
	defer close(headEvents)

	if err := indexer.listener.Listen(ctx, headEvents); err != nil {
		return NewIndexerListenerError(err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case headEvent := <-headEvents:
			log.Info().
				Uint64("slot", uint64(headEvent.Slot)).
				Msg("processing new head event")

			aggregatedRelayData, err := indexer.aggregator.AggregateDataForSlotFromRelays(ctx, headEvent.Slot)
			if err != nil {
				log.Err(err).Msg("error encountered while aggregating relay data")
			}

			preprocessedAggregatedRelayData := indexer.preprocessor.Preprocess(aggregatedRelayData)

			log.Info().Msg("aggregated relay data successfully preprocessed")

			if err := indexer.storage.StoreAggregatedRelayData(
				ctx,
				preprocessedAggregatedRelayData,
			); err != nil {
				log.Err(err).Msg("failed to store aggregated relay data")
			}

			log.Info().Msg("aggregated relay data successfully stored in database!")
		}
	}
}

func (indexer *Indexer) Stop() {
	indexer.storage.Shutdown()
}
