package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryRevisionsStandard(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err :=
		c.QueryRevisions().
			Titles("Main_Page", "Help:Introduction to Yextipedia").
			Do(context.Background())

	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Query)
	assert.Len(t, r.Query.Pages, 2)
	assert.NotEmpty(t, r.Query.Pages[1].Revisions[0].Slots["main"].Content)

	CompareJSON(t, r.RawJSON, r, false)
}

func TestQueryRevisionsWarning(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.QueryRevisions().Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)

	CompareJSON(t, r.RawJSON, r, false)
}
