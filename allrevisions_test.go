package mediawiki

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllrevisionsStandard(t *testing.T) {
	ctx := context.Background()

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	r, err := c.Allrevisions().Limit(1).Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	assert.Nil(t, r.Warnings)
	assert.NotNil(t, r.Continue)
	assert.NotNil(t, r.Query)
	assert.Len(t, r.Query.Allrevisions, 1)

	CompareJSON(t, r.RawJSON, r, false)
}

func TestAllrevisionsFirst(t *testing.T) {
	ctx := context.Background()

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	r, err := c.Allrevisions().Namespace(NamespaceMain).Dir("newer").Limit(0).Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	assert.Nil(t, r.Warnings)
	assert.NotNil(t, r.Continue)
	assert.NotNil(t, r.Query)
	assert.Len(t, r.Query.Allrevisions, 1)

	var oldest AllrevisionsResponseQueryRevision
	var ot = time.Now()

	for _, arev := range r.Query.Allrevisions {
		for i, rev := range arev.Revisions {
			if rev.Timestamp.Before(ot) {
				oldest = arev
				ot = *rev.Timestamp
				fmt.Println(i)
			}
		}
	}

	fmt.Printf("OLDEST: %q %v\n", oldest.Title, ot)

	// CompareJSON(t, r.RawJSON, r, false)
}

func TestAllrevisionsError(t *testing.T) {
	ctx := context.Background()

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	r, err := c.Allrevisions().Diffto("-1").Limit(1).Do(ctx)
	require.Error(t, err)
	assert.NotNil(t, r.Error)

	CompareJSON(t, r.RawJSON, r, false)
}

func TestAllrevisionsWarn(t *testing.T) {
	ctx := context.Background()

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(ctx, username, password)
	require.NoError(t, err)

	r, err := c.Allrevisions().Limit(1).Prop("parsetree").Do(ctx)
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	assert.NotNil(t, r.Warnings)
	assert.NotNil(t, r.Query)

	CompareJSON(t, r.RawJSON, r, false)
}
