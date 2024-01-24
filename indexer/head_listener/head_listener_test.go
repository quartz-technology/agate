//nolint:exhaustruct
package head_listener

import (
	"context"
	"testing"

	client "github.com/attestantio/go-eth2-client"
	v1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/quartz-technology/agate/internal/mocks"
	"github.com/stretchr/testify/require"
)

func TestNewAgateHeadListener(t *testing.T) {
	t.Parallel()

	listener := NewAgateHeadListener()

	require.NotNil(t, listener)
	require.Nil(t, listener.beaconAPIClient)
}

func TestAgateHeadListener_Init(t *testing.T) {
	t.Parallel()

	listener := NewAgateHeadListener()
	mockBeaconAPIClient := mocks.NewMockBeaconAPIClient(nil).WithDefaultInternalImplementations()

	listener.Init(mockBeaconAPIClient)

	require.NotNil(t, listener.beaconAPIClient)
}

func TestAgateHeadListener_Listen(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		mockEvents      []*v1.Event
		headEventsCount int
	}{
		"should capture unique head event": {
			mockEvents: []*v1.Event{
				{
					Data: &v1.HeadEvent{
						Slot: 0,
					},
				},
			},
			headEventsCount: 1,
		},
		"should capture multiple head events": {
			mockEvents: []*v1.Event{
				{
					Data: &v1.HeadEvent{
						Slot: 0,
					},
				},
				{
					Data: &v1.HeadEvent{
						Slot: 1,
					},
				},
			},
			headEventsCount: 2,
		},
		"should capture multiple head events only": {
			mockEvents: []*v1.Event{
				{
					Data: &v1.HeadEvent{
						Slot: 0,
					},
				},
				{
					Data: &v1.ChainReorgEvent{
						Slot: 0,
					},
				},
				{
					Data: &v1.HeadEvent{
						Slot: 1,
					},
				},
			},
			headEventsCount: 2,
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			listener := NewAgateHeadListener()
			mockBeaconAPIClient := mocks.NewMockBeaconAPIClient(
				func(ctx context.Context, handlerFunc client.EventHandlerFunc) error {
					for _, mockEvent := range tc.mockEvents {
						handlerFunc(mockEvent)
					}

					return nil
				},
			)

			headEvents := make(chan *v1.HeadEvent, tc.headEventsCount)

			listener.Init(mockBeaconAPIClient)

			err := listener.Listen(ctx, headEvents)
			require.NoError(t, err)

			acc := make([]*v1.HeadEvent, 0)

			for i := 0; i < tc.headEventsCount; i++ {
				headEvent := <-headEvents
				acc = append(acc, headEvent)
			}

			require.Len(t, acc, tc.headEventsCount)

			close(headEvents)
		})
	}
}
