package data_preprocessor

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/agate/indexer/storage_manager/store/dto"
)

// The DataPreprocessor is used to transform the raw data acquired by the data aggregator
// service, which is then used by the storage manager to save it in a database.
type DataPreprocessor struct{}

// NewDataPreprocessor creates an empty and non-initialized DataPreprocessor.
func NewDataPreprocessor() *DataPreprocessor {
	return &DataPreprocessor{}
}

// Preprocess transforms the data_aggregator.DataAggregator's aggregation output into a data
// structure that the storage_manager.StorageManager service can store in the database.
//
//nolint:funlen
func (preprocessor *DataPreprocessor) Preprocess(
	data *common.AggregatedRelayData,
) *DataPreprocessorOutput {
	// A mapping between each relay (represented by their URL) and the bids each one of them has
	// received.
	bidsPerRelay := make(map[string][]*dto.Bid)
	// A mapping between each bid and their submissions made to relay(s).
	submissionsPerBid := make(map[*dto.Bid][]*dto.Submission)

	for relayAPIURL, relayData := range *data {
		for _, bidReceived := range relayData.BidsReceived {
			bid := &dto.Bid{
				Slot:         phase0.Slot(bidReceived.Slot),
				ParentHash:   bidReceived.ParentHash,
				BlockHash:    bidReceived.BlockHash,
				FeeRecipient: bidReceived.ProposerFeeRecipient,
				GasLimit:     bidReceived.GasLimit,
				GasUsed:      bidReceived.GasUsed,
				Value:        bidReceived.Value,
				NumTx:        bidReceived.NumTx,
				Proposer:     bidReceived.ProposerPubkey,
				Builder:      bidReceived.BuilderPubkey,
			}

			bidsPerRelay[relayAPIURL] = append(bidsPerRelay[relayAPIURL], bid)

			submission := &dto.Submission{
				RelayURL:     relayAPIURL,
				IsDelivered:  false,
				IsOptimistic: bidReceived.OptimisticSubmission,
				SubmittedAt:  bidReceived.TimestampMs,
			}

			// Looks into all the bids delivered to this relay.
			// If one of them is the bid currently in process, mark it as delivered.
			// Usually, there is only one bid delivered by each - but this might require more
			// investigation.
			for _, bidDelivered := range relayData.BidsDelivered {
				if bidReceived.BlockHash.String() == bidDelivered.BlockHash.String() {
					submission.IsDelivered = true
				}
			}

			submissionsPerBid[bid] = append(submissionsPerBid[bid], submission)
		}
	}

	res := make([]*PreprocessedRelayData, 0)

	// Uses the two mappings created previously to reduce the input mapping into an array.
	for _, bidDTOs := range bidsPerRelay {
		for _, bidDTO := range bidDTOs {
			item := new(PreprocessedRelayData)

			item.Bid = bidDTO
			item.Submissions = submissionsPerBid[bidDTO]

			res = append(res, item)
		}
	}

	return &DataPreprocessorOutput{
		Output: res,
	}
}
