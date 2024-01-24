package data_aggregator

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/redax-go/relay"
	"github.com/quartz-technology/redax-go/sdk"
	datav1 "github.com/quartz-technology/redax-go/sdk/data/v1"
)

// RelayAPIClient is used to collect data from a relay.
type RelayAPIClient interface {
	// GetRelayDataForSlot retrieves the bids received and delivered to a relay for a specific slot.
	GetRelayDataForSlot(ctx context.Context, slot phase0.Slot) (*common.RelayData, error)
	// GetRelayAPIURL returns the API URL of the relay this client makes request to.
	GetRelayAPIURL() string
}

// The AgateRelayAPIClient is an implementation of the RelayAPIClient using the Quartz Technology
// <redax> SDK to interact with the relay.
type AgateRelayAPIClient struct {
	sdk         *sdk.RelaySDK
	relayAPIURL string
}

// NewAgateRelayAPIClient creates an empty and non-initialized AgateRelayAPIClient.
// It stores the relay API URL used to create the <redax> SDK.
func NewAgateRelayAPIClient(relayAPIURL string) *AgateRelayAPIClient {
	return &AgateRelayAPIClient{
		sdk:         nil,
		relayAPIURL: relayAPIURL,
	}
}

// Init initializes the AgateRelayAPIClient using the previously stored relay API URL.
func (client *AgateRelayAPIClient) Init() error {
	clt, err := relay.NewClient(relay.WithAPIURL(client.relayAPIURL))
	if err != nil {
		return NewAgateRelayAPIClientSDKInitializationError(err)
	}

	client.sdk = sdk.NewRelaySDK(clt)

	return nil
}

// GetRelayDataForSlot implements RelayAPIClient.GetRelayDataForSlot.
func (client *AgateRelayAPIClient) GetRelayDataForSlot(
	ctx context.Context,
	slot phase0.Slot,
) (*common.RelayData, error) {
	var (
		err error
		res = new(common.RelayData)
	)

	res.BidsReceived, err = client.sdk.Data().V1().GetBidsReceived(
		ctx,
		datav1.NewGetBidsReceivedRequest().WithSlot(slot),
	)
	if err != nil {
		return nil, NewAgateRelayAPIClientBidsReceivedRetrievalError(err)
	}

	res.BidsDelivered, err = client.sdk.Data().V1().GetBidsDelivered(
		ctx,
		datav1.NewGetBidsDeliveredRequest().WithSlot(slot),
	)
	if err != nil {
		return nil, NewAgateRelayAPIClientBidsDeliveredRetrievalError(err)
	}

	return res, nil
}

func (client *AgateRelayAPIClient) GetRelayAPIURL() string {
	return client.relayAPIURL
}
