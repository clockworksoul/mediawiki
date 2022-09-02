package mediawiki

import (
	"context"
	"fmt"
	"strings"
)

func (w *Client) QueryRevisions(ctx context.Context, titles ...string) (Response, error) {
	if err := w.checkKeepAlive(ctx); err != nil {
		return Response{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action":        "query",
		"prop":          "revisions",
		"titles":        strings.Join(titles, "|"),
		"rvslots":       "*",
		"rvprop":        "content",
		"formatversion": "2",
	}

	// Make the request.
	r, err := w.Get(ctx, parameters)
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	}

	return r, nil
}
