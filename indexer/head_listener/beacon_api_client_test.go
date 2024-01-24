package head_listener

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAgateBeaconAPIClient(t *testing.T) {
	t.Parallel()

	client := NewAgateBeaconAPIClient()

	require.NotNil(t, client)
	require.Nil(t, client.clt)
}
