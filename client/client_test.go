package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Call(t *testing.T) {
	_, err := Call("")
	require.Error(t, err, "Empty Call should return an error")
}
