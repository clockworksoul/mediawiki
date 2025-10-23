package mediawiki

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImagesBasic(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	c.Debug = os.Stdout

	r, err := c.Images().Titles("Axon").Do(context.Background())
	require.NoError(t, err)

	CompareJSON(t, r.RawJSON, r, false)
}
