//nolint:exhaustruct
package data

import (
	"testing"

	"github.com/attestantio/go-builder-client/api/v1"
	"github.com/quartz-technology/agate/indexer/common"
	"github.com/quartz-technology/agate/indexer/storage/store/dto"
	datav1 "github.com/quartz-technology/redax-go/sdk/data/v1"
	"github.com/stretchr/testify/require"
)

func TestNewPreprocessor(t *testing.T) {
	t.Parallel()

	preprocessor := NewPreprocessor()

	require.NotNil(t, preprocessor)
}

func TestPreprocessor_Preprocess(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input  *common.AggregatedRelayData
		output *common.DataPreprocessorOutput
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
			output: &common.DataPreprocessorOutput{
				Output: []*common.PreprocessedRelayData{
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
			output: &common.DataPreprocessorOutput{
				Output: []*common.PreprocessedRelayData{
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
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			preprocessor := NewPreprocessor()
			data := preprocessor.Preprocess(tc.input)

			require.Equal(t, tc.output, data)
		})
	}
}
