package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllusersStandard(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Allusers().Do(context.Background())
	require.NoError(t, err)
	require.Nil(t, r.Error)
	assert.NotEmpty(t, r.Query.Allusers)
	assert.Len(t, r.Query.Allusers, 2)
	assert.Equal(t, "Mtitmus", r.Query.Allusers[1].Name)
	assert.NotZero(t, r.Query.Allusers[1].UserId)

	CompareJSON(t, r.RawJSON, r, false)
}

func TestAllusersContinue(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Allusers().Limit(1).Do(context.Background())
	require.NoError(t, err)
	require.Nil(t, r.Error)
	assert.Len(t, r.Query.Allusers, 1)
	assert.NotEmpty(t, r.Continue.Continue)

	CompareJSON(t, r.RawJSON, r, false)
}
