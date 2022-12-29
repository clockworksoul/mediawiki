package mediawiki

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTranscludedinBasic(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	c.Debug = os.Stdout

	r, err := c.Transcludedin().Titles("Template:Test").Do(context.Background())
	require.NoError(t, err)

	CompareJSON(t, r.RawJSON, r, true)
}

func TestTranscludedinLimit(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	c.Debug = os.Stdout

	r, err := c.Transcludedin().Limit(1).Titles("Template:Test").Do(context.Background())
	require.NoError(t, err)

	CompareJSON(t, r.RawJSON, r, true)
}
