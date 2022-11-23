package mediawiki

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryAllpages(t *testing.T) {
	ctx := context.Background()

	c, err := New(apiUrl, agent)
	require.NoError(t, err)
	c.Debug = os.Stdout

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	r, err := c.QueryAllpages().Limit(1).From("T").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)

	m, _ := json.Marshal(r)
	assert.JSONEq(t, r.RawJSON, string(m))
}
