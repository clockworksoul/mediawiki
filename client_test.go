package mediawiki

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	agent = "test-bot"
)

var (
	apiUrl   = os.Getenv("MEDIAWIKI_URL")
	password = os.Getenv("MEDIAWIKI_PASSWORD")
	username = os.Getenv("MEDIAWIKI_USERNAME")
)

func TestClientGetToken(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	token, err := c.GetToken(context.Background(), LoginToken)
	require.NoError(t, err)
	require.Greater(t, len(token), 0)
}

func TestClientBotLogin(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	r, err := c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)
	require.NotNil(t, r.BotLogin)
	assert.Equal(t, Success, r.BotLogin.Result)
}

func TestClientParseResponse(t *testing.T) {
	mock := `{
		"batchcomplete": "Foo!",
		"warnings": {
			"tokens": {
				"*": "Warning!"
			}
		},
		"query": {
			"tokens": {
				"logintoken": "!!TOKEN!!"
			}
		},
		"error": {
			"code": "anerror",
			"info": "You got an error."
		}
	}`

	r := Response{}
	err := ParseResponse([]byte(mock), &r)
	require.NoError(t, err)

	assert.Equal(t, "Foo!", r.BatchComplete)

	assert.NotNil(t, r.Warnings)
	assert.Equal(t, "Warning!", r.Warnings.Tokens["*"])

	assert.NotNil(t, r.Query)
	assert.Equal(t, "!!TOKEN!!", r.Query.Tokens["logintoken"])

	assert.NotNil(t, r.Error)
	assert.Equal(t, "anerror", r.Error.Code)
	assert.Equal(t, "You got an error.", r.Error.Info)
}

func TestClientProtect(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Edit().Title("Protection test").Text("This is a test.").Summary("Automated test.").Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Edit)
	assert.Equal(t, Success, r.Edit.Result)

	r2, err := c.Protect(context.Background(), "Protection test", "This is a test")
	require.NoError(t, err)
	assert.Nil(t, r2.Error)
}
