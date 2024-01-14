package data_aggregator

import (
	"context"
	"sync"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/common"
)

// DataAggregator is used to aggregate data from multiple relays.
type DataAggregator interface {
	// AggregateDataForSlotFromRelays collects data from multiple relays for a specific slot.
	AggregateDataForSlotFromRelays(
		ctx context.Context,
		slot phase0.Slot,
	) (*common.AggregatedRelayData, error)
}

// relayResponse is used by the AgateDataAggregator during the map/reduce phase of the relay data
// aggregation process.
//
// As the AgateDataAggregator uses one goroutine per relay to query, it uses a channel of
// relayResponse to get back the responses from all the goroutines.
type relayResponse struct {
	// The API URL of the queried relay.
	relayAPIURL string
	// The content retrieved from the relay.
	data *common.RelayData
	// An optional error if the request made to the relay has encountered an error.
	err error
}

// The AgateDataAggregator is an implementation of the DataAggregator which uses multiple
// RelayAPIClient (one per relay) to aggregate relay data.
type AgateDataAggregator struct {
	relayAPIClients []RelayAPIClient
}

// NewAgateDataAggregator creates an empty and non-initialized AgateDataAggregator.
func NewAgateDataAggregator() *AgateDataAggregator {
	return &AgateDataAggregator{
		relayAPIClients: make([]RelayAPIClient, 0),
	}
}

// Init initializes an AgateDataAggregator service given multiple clients able to collect data
// from relays.
func (aggregator *AgateDataAggregator) Init(relayAPIClients ...RelayAPIClient) {
	aggregator.relayAPIClients = relayAPIClients
}

// AggregateDataForSlotFromRelays implements DataAggregator.AggregateDataForSlotFromRelays.
//
// It uses a map/reduce procedure to distribute each request made to the relays to distinct
// workers (the map phase).
// Then, it waits for the results to come back and aggregates those in a single placeholder
// (the reduce phase).
//
// In the scenario where part of the data collection has failed, the aggregator still returns the
// successfully collected data.
func (aggregator *AgateDataAggregator) AggregateDataForSlotFromRelays(
	ctx context.Context,
	slot phase0.Slot,
) (*common.AggregatedRelayData, error) {
	wg := sync.WaitGroup{}
	data := make(common.AggregatedRelayData)

	numRelays := len(aggregator.relayAPIClients)
	relayResponses := make(chan *relayResponse, numRelays)

	for _, relayAPIClient := range aggregator.relayAPIClients {
		wg.Add(1)

		go func(relayAPIClient RelayAPIClient, slot phase0.Slot) {
			var (
				relayDataResult       = new(common.RelayData)
				relayDataError  error = nil
			)

			defer func() {
				relayResponses <- &relayResponse{
					relayAPIURL: relayAPIClient.GetRelayAPIURL(),
					data:        relayDataResult,
					err:         relayDataError,
				}
				wg.Done()
			}()

			relayDataResult, relayDataError = relayAPIClient.GetRelayDataForSlot(ctx, slot)
		}(relayAPIClient, slot)
	}

	wg.Wait()

	close(relayResponses)

	err := NewAgateDataAggregationError()

	for relayRes := range relayResponses {
		if relayRes.err != nil {
			err.RecordFailure(relayRes.relayAPIURL, relayRes.err)
		} else {
			data[relayRes.relayAPIURL] = relayRes.data
		}
	}

	if len(err.RelayErrors) > 0 {
		err.Slot = slot

		// We still want to return the data that has been successfully retrieved from the relays.
		return &data, err
	}

	return &data, nil
}
