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

// Transcludedin

type TranscludedinResponse struct {
	CoreResponse
	BatchComplete any                    `json:"batchcomplete,omitempty"`
	Continue      *TranscludedinContinue `json:"continue,omitempty"`
	Query         *TranscludedinQuery    `json:"query,omitempty"`
}

type TranscludedinContinue struct {
	Lhcontinue string `json:"lhcontinue"`
	Continue   string `json:"continue"`
}

type TranscludedinQuery struct {
	Pages map[string]TranscludedinFromPage `json:"pages"`
}

type TranscludedinFromPage struct {
	Pageid        int                 `json:"pageid"`
	Ns            Namespace           `json:"ns"`
	Title         string              `json:"title"`
	Missing       any                 `json:"missing,omitempty"`
	Transcludedin []TranscludedinPage `json:"transcludedin,omitempty"`
}

type TranscludedinPage struct {
	Pageid   int       `json:"pageid"`
	Ns       Namespace `json:"ns"`
	Title    string    `json:"title"`
	Redirect any       `json:"redirect"`
}

type TranscludedinOption func(map[string]string)

type TranscludedinClient struct {
	o []TranscludedinOption
	c *Client
}

func (c *Client) Transcludedin() *TranscludedinClient {
	return &TranscludedinClient{c: c}
}

// Prop. Which properties to get:
//
// * pageid - Page ID of each page.
// * title - Title of each page.
// * redirect - Flag if the page is a redirect.
//
// Values (separate with | or alternative): pageid, redirect, title
// Default: pageid|title|redirect
func (w *TranscludedinClient) Prop(s ...string) *TranscludedinClient {
	w.o = append(w.o, func(m map[string]string) {
		m["tiprop"] = strings.Join(s, "|")
	})
	return w
}

// Namespace. Only include pages in these namespaces.
//
// Values (separate with | or alternative): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, etc.
//
// To specify all values, use NamespaceAll.
func (w *TranscludedinClient) Namespace(ns ...Namespace) *TranscludedinClient {
	w.o = append(w.o, func(m map[string]string) {
		if len(ns) > 0 && ns[0] == NamespaceAll {
			m["tinamespace"] = "*"
			return
		}

		var s []string
		for _, n := range ns {
			s = append(s, strconv.FormatInt(int64(n), 10))
		}
		m["tinamespace"] = strings.Join(s, "|")
	})
	return w

}

// Show
// Show only items that meet these criteria:
// * redirect - Only show redirects.
// * !redirect - Only show non-redirects.
//
// Values (separate with | or alternative): !redirect, redirect
func (w *TranscludedinClient) Show(s ...string) *TranscludedinClient {
	w.o = append(w.o, func(m map[string]string) {
		m["tishow"] = strings.Join(s, "|")
	})
	return w
}

// Limit
// How many to return.
// Type: integer
// The value must be between 1 and 500. For "max", use 0.
// Default: 10
func (w *TranscludedinClient) Limit(i int) *TranscludedinClient {
	w.o = append(w.o, func(m map[string]string) {
		m["tilimit"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Continue
// When more results are available, use this to continue.
func (w *TranscludedinClient) Continue(s string) *TranscludedinClient {
	w.o = append(w.o, func(m map[string]string) {
		m["ticontinue"] = s
	})
	return w
}

// Titles. A list of titles to work on.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *TranscludedinClient) Titles(s ...string) *TranscludedinClient {
	w.o = append(w.o, func(m map[string]string) {
		m["titles"] = strings.Join(s, "|")
	})
	return w
}

func (w *TranscludedinClient) Do(ctx context.Context) (TranscludedinResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return TranscludedinResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"prop":   "transcludedin",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := TranscludedinResponse{}
	j, err := w.c.GetInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Query == nil {
		return r, fmt.Errorf("unexpected error in query")
	}

	return r, nil
}
