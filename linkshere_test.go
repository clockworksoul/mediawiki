package mediawiki

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinksHereBasic(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	c.Debug = os.Stdout

	r, err := c.Linkshere().Titles("Link target").Do(context.Background())
	require.NoError(t, err)

	CompareJSON(t, r.RawJSON, r, true)
}

func TestLinksHereLimit(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	c.Debug = os.Stdout

	r, err := c.Linkshere().Limit(1).Titles("Link target").Do(context.Background())
	require.NoError(t, err)

	CompareJSON(t, r.RawJSON, r, true)
}
