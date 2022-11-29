package mediawiki

import (
	"context"
	"fmt"
	"strconv"
)

type CategoryMembers struct {
	QueryResponse
	BatchComplete string                   `json:"batchcomplete"`
	Continue      *CategoryMembersContinue `json:"continue,omitempty"`
	Query         *CategoryMembersQuery    `json:"query"`
}

type CategoryMembersQuery struct {
	CategoryMembers []QueryResponseQueryPage `json:"categorymembers"`
}

type CategoryMembersContinue struct {
	CmContinue string `json:"cmcontinue"`
	Continue   string `json:"continue"`
}

type CategoryMembersClient struct {
	o []QueryOption
	c *Client
}

func (c *Client) CategoryMembers() *CategoryMembersClient {
	return &CategoryMembersClient{c: c}
}

func (w *CategoryMembersClient) Title(s string) *CategoryMembersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cmtitle"] = s
	})
	return w
}

func (w *CategoryMembersClient) PageId(i int) *CategoryMembersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cmpageid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

func (w *CategoryMembersClient) Continue(s string) *CategoryMembersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cmcontinue"] = s
	})
	return w
}

func (w *CategoryMembersClient) Limit(i int) *CategoryMembersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cmlimit"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

func (w *CategoryMembersClient) Do(ctx context.Context) (CategoryMembers, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return CategoryMembers{}, err
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
	r := CategoryMembers{}
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
