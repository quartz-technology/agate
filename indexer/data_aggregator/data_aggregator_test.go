package data_aggregator

import (
	"context"
	"errors"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/agate/internal/mocks"
	datav1 "github.com/quartz-technology/redax-go/sdk/data/v1"
	"github.com/stretchr/testify/require"
)

func TestNewAgateDataAggregator(t *testing.T) {
	aggregator := NewAgateDataAggregator()

	require.NotNil(t, aggregator)
	require.NotNil(t, aggregator.relayAPIClients)
	require.Len(t, aggregator.relayAPIClients, 0)
}

func TestAgateDataAggregator_Init(t *testing.T) {
	aggregator := NewAgateDataAggregator()
	aggregator.Init()

	require.Nil(t, aggregator.relayAPIClients)
	require.Len(t, aggregator.relayAPIClients, 0)
}

func TestAgateDataAggregator_AggregateDataForSlotFromRelays(t *testing.T) {
	testCases := map[string]struct {
		relayAPIClient           []RelayAPIClient
		slot                     phase0.Slot
		expectedAggregatedData   *common.AggregatedRelayData
		expectedAggregationError error
	}{
		"should aggregate data with errors": {
			relayAPIClient: []RelayAPIClient{
				mocks.NewMockRelayAPIClient(
					func(ctx context.Context, slot phase0.Slot) (*common.RelayData, error) {
						return &common.RelayData{
							BidsReceived:  make([]*datav1.BidReceived, 0),
							BidsDelivered: make([]*datav1.BidDelivered, 0),
						}, nil
					},
					func() string {
						return "https://mock.functional.relay.com"
					},
				),
				mocks.NewMockRelayAPIClient(
					func(ctx context.Context, slot phase0.Slot) (*common.RelayData, error) {
						return nil, errors.New("mocked error")
					},
					func() string {
						return "https://mock.failing.relay.com"
					},
				),
			},
			slot: 42,
			expectedAggregatedData: &common.AggregatedRelayData{
				"https://mock.functional.relay.com": &common.RelayData{
					BidsReceived:  make([]*datav1.BidReceived, 0),
					BidsDelivered: make([]*datav1.BidDelivered, 0),
				},
			},
			expectedAggregationError: &AgateDataAggregationError{
				Slot: 42,
				RelayErrors: map[string]error{
					"https://mock.failing.relay.com": errors.New("mocked error"),
				},
			},
		},
		"should aggregate data without errors": {
			relayAPIClient: []RelayAPIClient{
				mocks.NewMockRelayAPIClient(
					func(ctx context.Context, slot phase0.Slot) (*common.RelayData, error) {
						return &common.RelayData{
							BidsReceived:  make([]*datav1.BidReceived, 0),
							BidsDelivered: make([]*datav1.BidDelivered, 0),
						}, nil
					},
					func() string {
						return "https://mock.functional.relay1.com"
					},
				), mocks.NewMockRelayAPIClient(
					func(ctx context.Context, slot phase0.Slot) (*common.RelayData, error) {
						return &common.RelayData{
							BidsReceived:  make([]*datav1.BidReceived, 0),
							BidsDelivered: make([]*datav1.BidDelivered, 0),
						}, nil
					},
					func() string {
						return "https://mock.functional.relay2.com"
					},
				),
			},
			slot: 0,
			expectedAggregatedData: &common.AggregatedRelayData{
				"https://mock.functional.relay1.com": &common.RelayData{
					BidsReceived:  make([]*datav1.BidReceived, 0),
					BidsDelivered: make([]*datav1.BidDelivered, 0),
				},
				"https://mock.functional.relay2.com": &common.RelayData{
					BidsReceived:  make([]*datav1.BidReceived, 0),
					BidsDelivered: make([]*datav1.BidDelivered, 0),
				},
			},
			expectedAggregationError: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			aggregator := NewAgateDataAggregator()
			aggregator.Init(tc.relayAPIClient...)

			aggregatedRelayData, err := aggregator.AggregateDataForSlotFromRelays(context.Background(), tc.slot)

			require.Equal(t, tc.expectedAggregatedData, aggregatedRelayData)
			require.Equal(t, tc.expectedAggregationError, err)
		})
	}
}
