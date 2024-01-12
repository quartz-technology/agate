package mocks

import (
	"context"

	client "github.com/attestantio/go-eth2-client"
	v1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

type MockBeaconAPIClient struct {
	subscribeToHeadEventsImpl func(context.Context, client.EventHandlerFunc) error
}

func NewMockBeaconAPIClient(
	subscribeToHeadEventsImpl func(context.Context, client.EventHandlerFunc) error,
) *MockBeaconAPIClient {
	return &MockBeaconAPIClient{
		subscribeToHeadEventsImpl: subscribeToHeadEventsImpl,
	}
}

func (m *MockBeaconAPIClient) WithDefaultInternalImplementations() *MockBeaconAPIClient {
	m.subscribeToHeadEventsImpl = func(ctx context.Context, handlerFunc client.EventHandlerFunc) error {
		handlerFunc(&v1.Event{
			Topic: "",
			Data: &v1.HeadEvent{
				Slot:                      0,
				Block:                     phase0.Root{},
				State:                     phase0.Root{},
				EpochTransition:           false,
				CurrentDutyDependentRoot:  phase0.Root{},
				PreviousDutyDependentRoot: phase0.Root{},
			},
		})

		return nil
	}

	return m
}

func (m *MockBeaconAPIClient) SubscribeToHeadEvents(
	ctx context.Context,
	handler client.EventHandlerFunc,
) error {
	return m.subscribeToHeadEventsImpl(ctx, handler)
}
