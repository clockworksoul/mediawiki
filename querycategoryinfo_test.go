package mediawiki

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediawikiQueryCategorylist(t *testing.T) {
	ctx := context.Background()

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	r, err := c.QueryCategoryInfo().Prop("categoryinfo").Titles("Category:Automatically converted pages").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)

	m, _ := json.Marshal(r)
	assert.JSONEq(t, r.RawJSON, string(m))
}
