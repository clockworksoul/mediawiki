package mediawiki

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUploadGood(t *testing.T) {
	name := "test-kitten.jpg"

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	// c.Debug = os.Stdout

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	defer func() {
		c.Debug = nil
		_, err := c.Delete().Title("File:" + name).Do(context.Background())
		assert.NoError(t, err)
	}()

	f, err := os.Open("test/kitten.jpg")
	require.NoError(t, err)

	r, err := c.Upload().Filename(name).File(f).Ignorewarnings(true).Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Upload)

	// Trick for forcing embedded HTML to use consistent encoding
	var i any
	json.Unmarshal([]byte(r.RawJSON), &i)
	b, _ := json.MarshalIndent(i, "", "  ")

	CompareJSON(t, string(b), r, false)
}

func TestUploadError(t *testing.T) {
	name := "test-kitten-2.jpg"

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	defer func() {
		c.Debug = nil
		_, err := c.Delete().Title("File:" + name).Do(context.Background())
		assert.NoError(t, err)
	}()

	// First upload should be as expected
	f, err := os.Open("test/kitten.jpg")
	require.NoError(t, err)
	r, err := c.Upload().Filename(name).File(f).Ignorewarnings(true).Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Upload)

	// Second upload should fail
	f, err = os.Open("test/kitten.jpg")
	require.NoError(t, err)
	r, err = c.Upload().Filename(name).File(f).Ignorewarnings(true).Do(context.Background())
	require.Error(t, err)
	assert.NotNil(t, r.Error)
	require.Nil(t, r.Upload)

	// Trick for forcing embedded HTML to use consistent encoding
	var i any
	json.Unmarshal([]byte(r.RawJSON), &i)
	b, _ := json.MarshalIndent(i, "", "  ")

	CompareJSON(t, string(b), r, false)
}

func TestUploadWarning(t *testing.T) {
	name1 := "test-kitten-warning-1.jpg"
	name2 := "test-kitten-warning-2.jpg"

	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	defer func() {
		c.Debug = nil
		_, err := c.Delete().Title("File:" + name1).Do(context.Background())
		assert.NoError(t, err)
	}()

	// First upload should be as expected
	f, err := os.Open("test/kitten.jpg")
	require.NoError(t, err)
	r, err := c.Upload().Filename(name1).File(f).Ignorewarnings(true).Do(context.Background())
	require.NoError(t, err)
	assert.Nil(t, r.Error)
	require.NotNil(t, r.Upload)

	f, err = os.Open("test/kitten.jpg")
	require.NoError(t, err)
	r, err = c.Upload().Filename(name2).File(f).Do(context.Background())
	require.NoError(t, err)
	require.NotNil(t, r.Upload)
	require.Equal(t, "Warning", r.Upload.Result)

	// Trick for forcing embedded HTML to use consistent encoding
	var i any
	json.Unmarshal([]byte(r.RawJSON), &i)
	b, _ := json.MarshalIndent(i, "", "  ")

	CompareJSON(t, string(b), r, false)
}
