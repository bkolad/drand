package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientPrivate(t *testing.T) {
	drands, dir := BatchNewDrand(5)
	defer CloseAllDrands(drands)
	defer os.RemoveAll(dir)

	client := NewClient()
	buff, err := client.Private(drands[0].priv.Public)
	require.Nil(t, err)
	require.NotNil(t, buff)
	require.Len(t, buff, 32)
}
