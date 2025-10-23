package mediawiki

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImageinfoBasic(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	c.Debug = os.Stdout

	r, err := c.Imageinfo().
		Prop("timestamp", "user", "userid", "comment", "parsedcomment", "canonicaltitle", "url", "size", "dimensions", "sha1", "mime", "thumbmime", "mediatype", "metadata", "commonmetadata", "extmetadata").
		Titles("File:Axon-black.png").Do(context.Background())
	require.NoError(t, err)

	CompareJSON(t, r.RawJSON, r, false)
}
