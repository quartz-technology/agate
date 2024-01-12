package indexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAgateBeaconAPIClient(t *testing.T) {
	client := NewAgateBeaconAPIClient()

	require.NotNil(t, client)
	require.Nil(t, client.clt)
}
