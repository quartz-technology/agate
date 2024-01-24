package storage

import (
	"context"

	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/agate/indexer/storage/store"
	"github.com/quartz-technology/agate/indexer/storage/store/dto"
	"github.com/quartz-technology/agate/indexer/storage/store/models"
)

// The DefaultManager is an implementation of the Manager using the default store.
type DefaultManager struct {
	store store.Store
}

// NewDefaultManager creates an empty and non-initialized DefaultManager.
func NewDefaultManager() *DefaultManager {
	return &DefaultManager{
		store: nil,
	}
}

// Init initializes a DefaultManager given the store used to perform database operations.
func (manager *DefaultManager) Init(store store.Store) {
	manager.store = store
}

func (manager *DefaultManager) StoreRelays(
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
			return NewDefaultManagerRelaysCreationError(err)
		}

		return nil
	})
	if err != nil {
		return NewDefaultManagerTransactionedRelaysCreationError(err)
	}

	return nil
}

func (manager *DefaultManager) StoreAggregatedRelayData(
	ctx context.Context,
	data *common.DataPreprocessorOutput,
) error {
	relays, err := manager.store.ListRelays(ctx)
	if err != nil {
		return NewDefaultManagerRelaysListingError(err)
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
			return NewDefaultManagerBidsCreationError(err)
		}

		if err := store.CreateSubmissions(ctx, submissions); err != nil {
			return NewDefaultManagerSubmissionsCreationError(err)
		}

		return nil
	})
	if err != nil {
		return NewDefaultManagerTransactionedBidsSubmissionsCreationError(err)
	}

	return nil
}

func (manager *DefaultManager) Shutdown() {
	manager.store.Close()
}
