package mediawiki

import (
	"context"
	"fmt"
	"strings"
)

type QueryRevisionsResponse struct {
	QueryResponse
}

type QueryRevisionsClient struct {
	o []QueryOption
	c *Client
}

// WithQueryProp
// Which properties to get for the queried pages.
func (c *Client) QueryRevisions() *QueryRevisionsClient {
	return &QueryRevisionsClient{c: c}
}

func (w *QueryRevisionsClient) Titles(s ...string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["titles"] = strings.Join(s, "|")
	})
	return w
}

func (w *QueryRevisionsClient) Do(ctx context.Context) (QueryRevisionsResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return QueryRevisionsResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action":        "query",
		"prop":          "revisions",
		"rvslots":       "*",
		"rvprop":        "content",
		"formatversion": "2",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := QueryRevisionsResponse{}
	j, err := w.c.GetInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to get: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	}

	return r, nil
}
