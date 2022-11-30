package mediawiki

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEditGood(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	text := "This is a test. " + time.Now().String()
	r, err := c.Edit().Title("TestMediawikiEditGood").Text("This is a test. " + text).Summary("Automated test.").Do(context.Background())

	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Edit)
	assert.Equal(t, Success, r.Edit.Result)

	CompareJSON(t, r.RawJSON, r, false)
}

func TestEditError(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Edit().Title("TestMediawikiEditError").Do(context.Background())

	require.Error(t, err)
	assert.NotNil(t, r.Error)
	require.Nil(t, r.Edit)

	CompareJSON(t, r.RawJSON, r, false)
}

func TestEditRepeated(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	text := "This is a test. " + time.Now().String()
	r, err := c.Edit().Title("TestMediawikiEditRepeated").Text(text).Summary("Automated test.").Watchlist("unwatch").Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Edit)
	assert.Nil(t, r.Edit.NoChange)
	assert.Equal(t, Success, r.Edit.Result)

	r, err = c.Edit().Title("TestMediawikiEditRepeated").Text(text).Summary("Automated test.").Watchlist("unwatch").Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Edit)
	assert.NotNil(t, r.Edit.NoChange)
	assert.Equal(t, Success, r.Edit.Result)

	CompareJSON(t, r.RawJSON, r, false)
}

func CompareJSON(t *testing.T, a, b any, verbose bool) bool {
	t.Helper()

	var ja, jb string
	var ok bool

	if ja, ok = a.(string); !ok {
		m, _ := json.Marshal(a)
		ja = string(m)
	}
	if jb, ok = b.(string); !ok {
		m, _ := json.Marshal(b)
		jb = string(m)
	}

	require.NotEmpty(t, ja)
	require.NotEmpty(t, jb)

	if verbose || !assert.JSONEq(t, ja, jb) {
		buf := &bytes.Buffer{}
		json.Indent(buf, []byte(ja), "", "  ")
		t.Log(buf.String())

		buf = &bytes.Buffer{}
		json.Indent(buf, []byte(jb), "", "  ")
		t.Log(buf.String())
		return false
	}

	return true
}
