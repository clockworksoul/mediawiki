package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryAllpages(t *testing.T) {
	ctx := context.Background()

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	r, err := c.QueryAllpages().Limit(1).From("T").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)

	CompareJSON(t, r.RawJSON, r, false)
}
