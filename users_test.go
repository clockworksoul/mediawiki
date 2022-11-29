package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersStandard(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Users().Users("Mtitmus").Do(context.Background())

	require.NoError(t, err)
	require.Nil(t, r.Error)
	assert.NotEmpty(t, r.Query.Users)
	assert.Len(t, r.Query.Users, 1)
	assert.Equal(t, "Mtitmus", r.Query.Users[0].Name)
	assert.NotZero(t, r.Query.Users[0].UserId)
	assert.Nil(t, r.Query.Users[0].Missing)

	CompareJSON(t, r.RawJSON, r, false)
}

func TestUsersMissing(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Users().Users("nosuchuser").Do(context.Background())

	require.NoError(t, err)
	require.Nil(t, r.Error)
	assert.NotEmpty(t, r.Query.Users)
	assert.Len(t, r.Query.Users, 1)
	assert.Equal(t, "Nosuchuser", r.Query.Users[0].Name)
	assert.Zero(t, r.Query.Users[0].UserId)
	assert.NotNil(t, r.Query.Users[0].Missing)

	CompareJSON(t, r.RawJSON, r, false)
}
