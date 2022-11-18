package mediawiki

import (
	"context"
	"fmt"
	"strconv"
)

type QueryCategoryMembersClient struct {
	o []QueryOption
	c *Client
}

// WithQueryProp
// Which properties to get for the queried pages.
func (c *Client) QueryCategoryMembers() *QueryCategoryMembersClient {
	return &QueryCategoryMembersClient{c: c}
}

func (w *QueryCategoryMembersClient) Title(s string) *QueryCategoryMembersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cmtitle"] = s
	})
	return w
}

func (w *QueryCategoryMembersClient) PageId(i int) *QueryCategoryMembersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cmpageid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

func (w *QueryCategoryMembersClient) Continue(s string) *QueryCategoryMembersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cmcontinue"] = s
	})
	return w
}

func (w *QueryCategoryMembersClient) Limit(i int) *QueryCategoryMembersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cmlimit"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

func (w *QueryCategoryMembersClient) Do(ctx context.Context) (QueryCategoryMembers, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return QueryCategoryMembers{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"list":   "categorymembers",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := QueryCategoryMembers{}
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

type QueryCategoryMembers struct {
	QueryResponse
	BatchComplete string                        `json:"batchcomplete"`
	Continue      *QueryCategoryMembersContinue `json:"continue"`
	Query         *QueryCategoryMembersQuery    `json:"query"`
}

type QueryCategoryMembersQuery struct {
	CategoryMembers []ResponseQueryPage `json:"categorymembers"`
}

type QueryCategoryMembersContinue struct {
	CmContinue string `json:"cmcontinue"`
	Continue   string `json:"continue"`
}
