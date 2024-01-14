package data_aggregator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAgateRelayAPIClient(t *testing.T) {
	client := NewAgateRelayAPIClient("https://example.com")

	require.NotNil(t, client)
	require.Nil(t, client.sdk)
}

func TestAgateRelayAPIClient_Init(t *testing.T) {
	testCases := map[string]struct {
		relayAPIURL string
		success     bool
	}{
		"should create relay API client": {
			relayAPIURL: "https://example.com",
			success:     true,
		},
		"should fail to create relay API client": {
			relayAPIURL: "https://example.com/",
			success:     false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			client := NewAgateRelayAPIClient(tc.relayAPIURL)
			err := client.Init()

			if tc.success {
				require.NoError(t, err)
				require.NotNil(t, client.sdk)
			} else {
				require.Error(t, err)
				require.Nil(t, client.sdk)
			}
		})
	}
}

func TestAgateRelayAPIClient_GetRelayAPIURL(t *testing.T) {
	relayAPIURL := "https://example.com"
	client := NewAgateRelayAPIClient(relayAPIURL)

	require.Equal(t, relayAPIURL, client.GetRelayAPIURL())
}
