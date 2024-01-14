package common

import datav1 "github.com/quartz-technology/redax-go/sdk/data/v1"

// RelayData is used by the data aggregator service to transmit collected data from the relays
// for a specific slot to the service responsible for storing it into the database.
type RelayData struct {
	BidsReceived  []*datav1.BidReceived
	BidsDelivered []*datav1.BidDelivered
}

// AggregatedRelayData is a 1:1 mapping between a relay API URL and the data collected by the
// data aggregator service from this relay for a specific slot.
// It is also used by the service responsible for storing this data.
type AggregatedRelayData = map[string]*RelayData
