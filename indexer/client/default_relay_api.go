package client

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/redax-go/relay"
	"github.com/quartz-technology/redax-go/sdk"
	datav1 "github.com/quartz-technology/redax-go/sdk/data/v1"
)

// The DefaultRelayAPI is an implementation of the RelayAPI using the Quartz Technology
// <redax> SDK to interact with the relay.
type DefaultRelayAPI struct {
	sdk         *sdk.RelaySDK
	relayAPIURL string
}

// NewDefaultRelayAPI creates an empty and non-initialized DefaultRelayAPI.
// It stores the relay API URL used to create the <redax> SDK.
func NewDefaultRelayAPI(relayAPIURL string) *DefaultRelayAPI {
	return &DefaultRelayAPI{
		sdk:         nil,
		relayAPIURL: relayAPIURL,
	}
}

// Init initializes the DefaultRelayAPI using the previously stored relay API URL.
func (client *DefaultRelayAPI) Init() error {
	clt, err := relay.NewClient(relay.WithAPIURL(client.relayAPIURL))
	if err != nil {
		return NewDefaultRelayAPIServiceInitializationError(err)
	}

	client.sdk = sdk.NewRelaySDK(clt)

	return nil
}

// GetRelayDataForSlot implements RelayAPIClient.GetRelayDataForSlot.
func (client *DefaultRelayAPI) GetRelayDataForSlot(
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
		return nil, NewDefaultRelayAPIBidsReceivedRetrievalError(err)
	}

	res.BidsDelivered, err = client.sdk.Data().V1().GetBidsDelivered(
		ctx,
		datav1.NewGetBidsDeliveredRequest().WithSlot(slot),
	)
	if err != nil {
		return nil, NewDefaultRelayAPIBidsDeliveredRetrievalError(err)
	}

	return res, nil
}

func (client *DefaultRelayAPI) GetRelayAPIURL() string {
	return client.relayAPIURL
}
