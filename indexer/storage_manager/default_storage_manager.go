package storage_manager

import (
	"context"

	"github.com/quartz-technology/agate/indexer/data_preprocessor"
	"github.com/quartz-technology/agate/indexer/storage_manager/store"
	"github.com/quartz-technology/agate/indexer/storage_manager/store/dto"
	"github.com/quartz-technology/agate/indexer/storage_manager/store/models"
)

// The DefaultStorageManager is an implementation of the StoreManager using the default
// preprocessed aggregated relay data format.
type DefaultStorageManager struct {
	store store.Store
}

// NewDefaultStorageManager creates an empty and non-initialized DefaultStorageManager.
func NewDefaultStorageManager() *DefaultStorageManager {
	return &DefaultStorageManager{
		store: nil,
	}
}

// Init initializes a DefaultStorageManager given the store used to perform database operations.
func (manager *DefaultStorageManager) Init(store store.Store) {
	manager.store = store
}

func (manager *DefaultStorageManager) StoreRelays(
	ctx context.Context,
	data []*dto.Relay,
) error {
	relays := make([]*models.Relay, 0)

	for _, relayDTO := range data {
		relays = append(
			relays,
			//nolint:exhaustruct
			&models.Relay{
				URL: relayDTO.URL,
			},
		)
	}

	err := manager.store.ExecInTx(ctx, func(store store.Store) error {
		if err := store.CreateRelays(ctx, relays); err != nil {
			return NewDefaultStorageManagerRelaysCreationError(err)
		}

		return nil
	})
	if err != nil {
		return NewDefaultStorageManagerTransactionedRelaysCreationError(err)
	}

	return nil
}

func (manager *DefaultStorageManager) StoreAggregatedRelayData(
	ctx context.Context,
	data *data_preprocessor.DataPreprocessorOutput,
) error {
	relays, err := manager.store.ListRelays(ctx)
	if err != nil {
		return NewDefaultStorageManagerRelaysListingError(err)
	}

	relayURLsToIDs := make(map[string]uint64)

	for _, relay := range relays {
		relayURLsToIDs[relay.URL] = relay.ID
	}

	bids := make([]*models.Bid, 0)
	submissions := make([]*models.Submission, 0)

	for _, relayData := range data.Output {
		bidModel := new(models.Bid).FromBidDTO(relayData.Bid)

		bids = append(bids, bidModel)

		for _, submission := range relayData.Submissions {
			//nolint:exhaustruct
			submissionModel := &models.Submission{
				RelayID:      relayURLsToIDs[submission.RelayURL],
				BidBlockHash: bidModel.BlockHash,
				IsDelivered:  submission.IsDelivered,
				IsOptimistic: submission.IsOptimistic,
				SubmittedAt:  submission.SubmittedAt,
			}

			submissions = append(submissions, submissionModel)
		}
	}

	err = manager.store.ExecInTx(ctx, func(store store.Store) error {
		if err := store.CreateBids(ctx, bids); err != nil {
			return NewDefaultStorageManagerBidsCreationError(err)
		}

		if err := store.CreateSubmissions(ctx, submissions); err != nil {
			return NewDefaultStorageManagerSubmissionsCreationError(err)
		}

		return nil
	})
	if err != nil {
		return NewDefaultStorageManagerTransactionedBidsSubmissionsCreationError(err)
	}

	return nil
}

func (manager *DefaultStorageManager) Shutdown() {
	manager.store.Close()
}
