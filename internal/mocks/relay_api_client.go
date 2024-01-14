package mocks

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/common"
	datav1 "github.com/quartz-technology/redax-go/sdk/data/v1"
)

type MockRelayAPIClient struct {
	getRelayDataForSlotImpl func(context.Context, phase0.Slot) (*common.RelayData, error)
	getRelayAPIURLImp       func() string
}

func NewMockRelayAPIClient(
	getRelayDataForSlotImpl func(context.Context, phase0.Slot) (*common.RelayData, error),
	getRelayAPIURLImp func() string,
) *MockRelayAPIClient {
	return &MockRelayAPIClient{
		getRelayDataForSlotImpl: getRelayDataForSlotImpl,
		getRelayAPIURLImp:       getRelayAPIURLImp,
	}
}

func (m *MockRelayAPIClient) WithDefaultInternalImplementations() *MockRelayAPIClient {
	m.getRelayDataForSlotImpl = func(ctx context.Context, slot phase0.Slot) (*common.RelayData, error) {
		return &common.RelayData{
			BidsReceived:  make([]*datav1.BidReceived, 0),
			BidsDelivered: make([]*datav1.BidDelivered, 0),
		}, nil
	}

	m.getRelayAPIURLImp = func() string {
		return "https://example.com"
	}

	return m
}

func (m *MockRelayAPIClient) GetRelayDataForSlot(
	ctx context.Context,
	slot phase0.Slot,
) (*common.RelayData, error) {
	return m.getRelayDataForSlotImpl(ctx, slot)
}

func (m *MockRelayAPIClient) GetRelayAPIURL() string {
	return m.getRelayAPIURLImp()
}
