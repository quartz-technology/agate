package data

import (
	"context"
	"errors"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/client"
	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/agate/internal/mocks"
	datav1 "github.com/quartz-technology/redax-go/sdk/data/v1"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultAggregator(t *testing.T) {
	t.Parallel()

	aggregator := NewDefaultAggregator()

	require.NotNil(t, aggregator)
	require.NotNil(t, aggregator.relayAPIClients)
	require.Len(t, aggregator.relayAPIClients, 0)
}

func TestDefaultAggregator_Init(t *testing.T) {
	t.Parallel()

	aggregator := NewDefaultAggregator()
	aggregator.Init()

	require.Nil(t, aggregator.relayAPIClients)
	require.Len(t, aggregator.relayAPIClients, 0)
}

func TestDefaultAggregator_AggregateDataForSlotFromRelays(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		relayAPIClient           []client.RelayAPI
		slot                     phase0.Slot
		expectedAggregatedData   *common.AggregatedRelayData
		expectedAggregationError error
	}{
		"should aggregate data with errors": {
			relayAPIClient: []client.RelayAPI{
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
			expectedAggregationError: &DefaultDataAggregationError{
				Slot: 42,
				RelayErrors: map[string]error{
					"https://mock.failing.relay.com": errors.New("mocked error"),
				},
			},
		},
		"should aggregate data without errors": {
			relayAPIClient: []client.RelayAPI{
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
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			aggregator := NewDefaultAggregator()
			aggregator.Init(tc.relayAPIClient...)

			aggregatedRelayData, err := aggregator.AggregateDataForSlotFromRelays(context.Background(), tc.slot)

			require.Equal(t, tc.expectedAggregatedData, aggregatedRelayData)
			require.Equal(t, tc.expectedAggregationError, err)
		})
	}
}
