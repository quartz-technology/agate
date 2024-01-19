package data_preprocessor

import (
	"testing"

	"github.com/attestantio/go-builder-client/api/v1"
	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/agate/indexer/storage_manager/store/dto"
	datav1 "github.com/quartz-technology/redax-go/sdk/data/v1"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultDataPreprocessor(t *testing.T) {
	preprocessor := NewDefaultDataPreprocessor()

	require.NotNil(t, preprocessor)
}

func TestDefaultDataPreprocessor_Preprocess(t *testing.T) {
	testCases := map[string]struct {
		input  *common.AggregatedRelayData
		output *DataPreprocessorOutput[[]*DefaultPreprocessedRelayData]
	}{
		"should preprocess non-delivered bid": {
			input: &common.AggregatedRelayData{
				"https://example.com": &common.RelayData{
					BidsReceived: []*datav1.BidReceived{
						{
							BidTrace:        v1.BidTrace{},
							BidReceivedMeta: datav1.BidReceivedMeta{},
						},
					},
					BidsDelivered: []*datav1.BidDelivered{},
				},
			},
			output: &DataPreprocessorOutput[[]*DefaultPreprocessedRelayData]{
				Output: []*DefaultPreprocessedRelayData{
					{
						Bid: &dto.Bid{},
						Submissions: []*dto.Submission{
							{
								RelayURL: "https://example.com",
							},
						},
					},
				},
			},
		},
		"should preprocess delivered bid": {
			input: &common.AggregatedRelayData{
				"https://example.com": &common.RelayData{
					BidsReceived: []*datav1.BidReceived{
						{
							BidTrace:        v1.BidTrace{},
							BidReceivedMeta: datav1.BidReceivedMeta{},
						},
					},
					BidsDelivered: []*datav1.BidDelivered{
						{
							BidTrace:         v1.BidTrace{},
							BidDeliveredMeta: datav1.BidDeliveredMeta{},
						},
					},
				},
			},
			output: &DataPreprocessorOutput[[]*DefaultPreprocessedRelayData]{
				Output: []*DefaultPreprocessedRelayData{
					{
						Bid: &dto.Bid{},
						Submissions: []*dto.Submission{
							{
								RelayURL:    "https://example.com",
								IsDelivered: true,
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			preprocessor := NewDefaultDataPreprocessor()
			data := preprocessor.Preprocess(tc.input)

			require.Equal(t, tc.output, data)
		})
	}
}
