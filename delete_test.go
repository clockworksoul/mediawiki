package mediawiki

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediawikiDeleteGood(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	// Create the page to delete

	er, err := c.Edit().Title("TestMediawikiDeleteGood").Text("This is a test. " + time.Now().String()).Summary("Automated test.").Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, er.Error)

	// Delete the page
	dr, err := c.Delete().Title("TestMediawikiDeleteGood").Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, dr.Error)
	require.NotNil(t, dr.Delete)

	CompareJSON(t, dr.RawJSON, dr, false)
}

func TestMediawikiDeleteError(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	// Delete the page
	dr, err := c.Delete().Title("TestMediawikiDeleteError").Do(context.Background())
	require.Error(t, err)
	assert.NotNil(t, dr.Error)
	require.Nil(t, dr.Delete)

	CompareJSON(t, dr.RawJSON, dr, true)
}
