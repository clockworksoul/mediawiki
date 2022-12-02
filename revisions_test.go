package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRevisionsStandard(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err :=
		c.Revisions().
			Titles("Main_Page", "Help:Introduction to Yextipedia").
			Do(context.Background())

	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Query)
	assert.Len(t, r.Query.Pages, 2)

	r.Query.Pages[1].Revisions[0].Tags = []string{}
	CompareJSON(t, r.RawJSON, r, false)
}

func TestRevisionsContent(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Revisions().Titles("Main_Page").Prop("content").Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Query)
	assert.NotEmpty(t, r.Query.Pages[0].Revisions[0].Slots["main"].Content)
}

func TestRevisionsOpts(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	// c.Debug = os.Stdout

	r, err :=
		c.Revisions().
			Titles("Main_Page", "Help:Introduction to Yextipedia").
			Slots("*").
			Prop("ids", "flags", "timestamp", "user", "userid", "size", "slotsize", "sha1", "slotsha1", "contentmodel", "comment", "content", "roles").
			Do(context.Background())

	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Query)
	assert.Len(t, r.Query.Pages, 2)

	r.Query.Pages[1].Revisions[0].Tags = []string{}
	CompareJSON(t, r.RawJSON, r, false)
}

func TestRevisionsWarning(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Revisions().Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)

	CompareJSON(t, r.RawJSON, r, false)
}
