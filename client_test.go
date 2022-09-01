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

func TestMediawikiClientGetToken(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	token, err := c.GetToken(context.Background(), LoginToken)
	require.NoError(t, err)
	require.Greater(t, len(token), 0)
}

func TestMediawikiClientClientLogin(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	r, err := c.ClientLogin(context.Background(), username, password)
	require.NoError(t, err)
	require.NotNil(t, r.ClientLogin)
	assert.Equal(t, "PASS", r.ClientLogin.Status)
}

func TestMediawikiClientParseResponse(t *testing.T) {
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

	r, err := ParseResponse([]byte(mock))
	require.NoError(t, err)

	assert.Equal(t, "Foo!", r.Batchcomplete)

	assert.NotNil(t, r.Warnings)
	assert.Equal(t, "Warning!", r.Warnings.Tokens["*"])

	assert.NotNil(t, r.Query)
	assert.Equal(t, "!!TOKEN!!", r.Query.Tokens["logintoken"])

	assert.NotNil(t, r.Error)
	assert.Equal(t, "anerror", r.Error.Code)
	assert.Equal(t, "You got an error.", r.Error.Info)
}

func TestMediawikiClientProtect(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.ClientLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Write(context.Background(), "Protection test", "This is a test.", "Automated test")
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Edit)
	assert.Equal(t, "Success", r.Edit.Result)

	r, err = c.Protect(context.Background(), "Protection test", "This is a test")
	require.NoError(t, err)
	assert.Nil(t, r.Error)
}

func TestMediawikiClientUpload(t *testing.T) {
	// c, err := New(apiUrl, agent)
	// require.NoError(t, err)

	// _, err = c.ClientLogin(context.Background(), username, password)
	// require.NoError(t, err)

	// img := documents.File{Path: "/Users/mtitmus/workspace/orgmd/images/hippo.jpg"}

	// in, err := img.Reader()
	// require.NoError(t, err)

	// _, err = c.Upload(context.Background(), "/Users/mtitmus/workspace/orgmd/images/hippo.jpg", in, "hippo.jpg")
	// require.NoError(t, err)
}

func TestMediawikiClientWrite(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.ClientLogin(context.Background(), username, password)
	require.NoError(t, err)

	r, err := c.Write(context.Background(), "Test", "This is a test.", "Automated test")
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Edit)
	assert.Equal(t, "Success", r.Edit.Result)
}
