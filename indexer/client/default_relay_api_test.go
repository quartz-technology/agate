package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDefaultRelayAPI(t *testing.T) {
	t.Parallel()

	client := NewDefaultRelayAPI("https://example.com")

	require.NotNil(t, client)
	require.Nil(t, client.sdk)
}

func TestDefaultRelayAPI_Init(t *testing.T) {
	t.Parallel()

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
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			client := NewDefaultRelayAPI(tc.relayAPIURL)
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

func TestDefaultRelayAPI_GetRelayAPIURL(t *testing.T) {
	t.Parallel()

	relayAPIURL := "https://example.com"
	client := NewDefaultRelayAPI(relayAPIURL)

	require.Equal(t, relayAPIURL, client.GetRelayAPIURL())
}
