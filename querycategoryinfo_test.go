package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryCategorylist(t *testing.T) {
	ctx := context.Background()

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	r, err := c.QueryCategoryInfo().Prop("categoryinfo").Titles("Category:Automatically converted pages").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)

	CompareJSON(t, r.RawJSON, r, false)
}
