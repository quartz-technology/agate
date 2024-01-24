package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDefaultBeaconAPI(t *testing.T) {
	t.Parallel()

	client := NewDefaultBeaconAPI()

	require.NotNil(t, client)
	require.Nil(t, client.clt)
}
