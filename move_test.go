package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMoveStandard(t *testing.T) {
	name := "Move test"
	name2 := "Move test target"

	ctx := context.Background()
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	defer func() {
		c.Debug = nil
		c.Delete().Title(name).Do(context.Background())
		c.Delete().Title(name2).Do(context.Background())
		c.Delete().Title("Talk:" + name2).Do(context.Background())
	}()

	r, err := c.Edit().Title(name).Text("This is a test.").Summary("Automated test.").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Edit)
	assert.Equal(t, Success, r.Edit.Result)

	c.Edit().Title("Talk:" + name).Text("This is a test.").Summary("Automated test.").Do(ctx)

	rp, err := c.Move().From(name).To(name2).Movetalk(true).Noredirect(true).Reason("Because I want to.").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, rp.Error)

	CompareJSON(t, rp.RawJSON, rp, false)
}

func TestMoveDoesntExist(t *testing.T) {
	name := "This page doesn't exist"
	name2 := "This page still doesn't exist"

	ctx := context.Background()
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	rp, err := c.Move().From(name).To(name2).Reason("Because I want to.").Do(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "missingtitle: ")

	CompareJSON(t, rp.RawJSON, rp, false)
}
