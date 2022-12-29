package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Fetch data from and about MediaWiki.
// All data modifications will first have to use query to acquire a token to prevent abuse from malicious sites.
//
// Flags:
// * This module requires read rights.

// Linkshere

type LinkshereResponse struct {
	CoreResponse
	BatchComplete any                `json:"batchcomplete,omitempty"`
	Continue      *LinkshereContinue `json:"continue,omitempty"`
	Query         *LinkshereQuery    `json:"query,omitempty"`
}

type LinkshereContinue struct {
	Lhcontinue string `json:"lhcontinue"`
	Continue   string `json:"continue"`
}

type LinkshereQuery struct {
	Pages map[string]LinkshereFromPage `json:"pages"`
}

type LinkshereFromPage struct {
	Pageid    int             `json:"pageid"`
	Ns        Namespace       `json:"ns"`
	Title     string          `json:"title"`
	Linkshere []LinksherePage `json:"linkshere"`
}

type LinksherePage struct {
	Pageid int       `json:"pageid"`
	Ns     Namespace `json:"ns"`
	Title  string    `json:"title"`
}

type LinkshereOption func(map[string]string)

type LinkshereClient struct {
	o []LinkshereOption
	c *Client
}

func (c *Client) Linkshere() *LinkshereClient {
	return &LinkshereClient{c: c}
}

// Prop. Which properties to get:
//
// * pageid - Page ID of each page.
// * title - Title of each page.
// * redirect - Flag if the page is a redirect.
//
// Values (separate with | or alternative): pageid, redirect, title
// Default: pageid|title|redirect
func (w *LinkshereClient) Prop(s ...string) *LinkshereClient {
	w.o = append(w.o, func(m map[string]string) {
		m["lhprop"] = strings.Join(s, "|")
	})
	return w
}

// Namespace. Only include pages in these namespaces.
//
// Values (separate with | or alternative): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, etc.
//
// To specify all values, use NamespaceAll.
func (w *LinkshereClient) Namespace(ns ...Namespace) *LinkshereClient {
	w.o = append(w.o, func(m map[string]string) {
		if len(ns) > 0 && ns[0] == NamespaceAll {
			m["lhnamespace"] = "*"
			return
		}

		var s []string
		for _, n := range ns {
			s = append(s, strconv.FormatInt(int64(n), 10))
		}
		m["lhnamespace"] = strings.Join(s, "|")
	})
	return w

}

// Show
// Show only items that meet these criteria:
// * redirect - Only show redirects.
// * !redirect - Only show non-redirects.
//
// Values (separate with | or alternative): !redirect, redirect
func (w *LinkshereClient) Show(s ...string) *LinkshereClient {
	w.o = append(w.o, func(m map[string]string) {
		m["lhshow"] = strings.Join(s, "|")
	})
	return w
}

// Limit
// How many to return.
// Type: integer
// The value must be between 1 and 500. For "max", use 0.
// Default: 10
func (w *LinkshereClient) Limit(i int) *LinkshereClient {
	w.o = append(w.o, func(m map[string]string) {
		m["lhlimit"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Continue
// When more results are available, use this to continue.
func (w *LinkshereClient) Continue(s string) *LinkshereClient {
	w.o = append(w.o, func(m map[string]string) {
		m["lhcontinue"] = s
	})
	return w
}

// Titles. A list of titles to work on.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *LinkshereClient) Titles(s ...string) *LinkshereClient {
	w.o = append(w.o, func(m map[string]string) {
		m["titles"] = strings.Join(s, "|")
	})
	return w
}

func (w *LinkshereClient) Do(ctx context.Context) (LinkshereResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return LinkshereResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"prop":   "linkshere",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := LinkshereResponse{}
	j, err := w.c.PostInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
		// } else if r.Linkshere == nil {
		// 	return r, fmt.Errorf("unexpected error in query")
	}

	return r, nil
}
