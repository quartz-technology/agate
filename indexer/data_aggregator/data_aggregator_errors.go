package data_aggregator

import (
	"fmt"

	"github.com/attestantio/go-eth2-client/spec/phase0"
)

// AgateDataAggregationError is raised if at least one of the request made to collect data from a
// relay has encountered an error.
type AgateDataAggregationError struct {
	// The slot used as a parameter in the request made to the relay.
	Slot phase0.Slot
	// A 1:1 mapping between a relay API URL and the error that has been encountered when making
	// the request to get the data from the relay.
	RelayErrors map[string]error
}

// NewAgateDataAggregationError creates an empty and non-initialized AgateDataAggregationError.
func NewAgateDataAggregationError() *AgateDataAggregationError {
	return &AgateDataAggregationError{
		Slot:        0,
		RelayErrors: make(map[string]error),
	}
}

// RecordFailure is used during the reduce phase of the aggregation process of the
// AgateDataAggregator service.
// It saves a record of an error that has been encountered when collecting data from a relay.
func (err *AgateDataAggregationError) RecordFailure(relayAPIURL string, relayError error) {
	err.RelayErrors[relayAPIURL] = relayError
}

func (err *AgateDataAggregationError) Error() string {
	return fmt.Sprintf(
		"data aggregator encountered the following error(s) for slot %d: %v",
		err.Slot,
		err.RelayErrors,
	)
}
