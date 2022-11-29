package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProtectStandard(t *testing.T) {
	name := "Protection test standard"
	ctx := context.Background()
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	defer func() {
		c.Debug = nil
		_, err := c.Delete().Title(name).Do(context.Background())
		assert.NoError(t, err)
	}()

	r, err := c.Edit().Title(name).Text("This is a test.").Summary("Automated test.").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Edit)
	assert.Equal(t, Success, r.Edit.Result)

	rp, err := c.Protect().Title(name).Protections("edit=sysop").Reason("This is a test.").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, rp.Error)

	CompareJSON(t, rp.RawJSON, rp, false)
}

func TestProtectError(t *testing.T) {
	ctx := context.Background()
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	rp, err := c.Protect().Title("No such page").Protections("edit=sysop").Reason("This is a test.").Do(ctx)
	require.Error(t, err)
	assert.NotNil(t, rp.Error)
	assert.Nil(t, rp.Protect)

	CompareJSON(t, rp.RawJSON, rp, false)
}
