package indexer

import (
	"context"

	v1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/quartz-technology/agate/indexer/data"
	"github.com/quartz-technology/agate/indexer/events"
	"github.com/quartz-technology/agate/indexer/storage"
	"github.com/rs/zerolog/log"
)

type Indexer struct {
	listener       events.HeadListener
	aggregator     data.Aggregator
	preprocessor   *data.Preprocessor
	storageManager storage.Manager
}

func New(listener events.HeadListener,
	aggregator data.Aggregator,
	preprocessor *data.Preprocessor,
	storage storage.Manager,
) *Indexer {
	return &Indexer{
		listener:       listener,
		aggregator:     aggregator,
		preprocessor:   preprocessor,
		storageManager: storage,
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

			if err := indexer.storageManager.StoreAggregatedRelayData(
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
	indexer.storageManager.Shutdown()
}
