package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Enumerate all pages sequentially in a given namespace.
//
// Flags:
// * This module requires read rights.
// * This module can be used as a generator.

type AllpagesResponse struct {
	QueryResponse
	BatchComplete string                    `json:"batchcomplete"`
	Continue      *AllpagesResponseContinue `json:"continue,omitempty"`
	Query         *AllpagesResponseQuery    `json:"query,omitempty"`
}

type AllpagesResponseContinue struct {
	GapContinue string `json:"gapcontinue"`
	Continue    string `json:"continue"`
}

type AllpagesResponseQuery struct {
	Pages map[string]QueryResponseQueryPage `json:"pages"`
}

// Allpages
type AllpagesOption func(map[string]string)

type AllpagesClient struct {
	o []QueryOption
	c *Client
}

func (c *Client) Allpages() *AllpagesClient {
	return &AllpagesClient{c: c}
}

// from
// The page title to start enumerating from.
func (w *AllpagesClient) From(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapfrom"] = s
	})
	return w
}

// continue
// When more results are available, use this to continue.
func (w *AllpagesClient) Continue(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapcontinue"] = s
	})
	return w
}

// to
// The page title to stop enumerating at.
func (w *AllpagesClient) To(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapto"] = s
	})
	return w
}

// prefix
// Search for all page titles that begin with this value.
func (w *AllpagesClient) Prefix(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapprefix"] = s
	})
	return w
}

// namespace
// The namespace to enumerate.
// One of the following values: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 90, 91, 92, 93, 100, 101, 102, 103, 104, 105, 106, 107, 486, 487, 710, 711, 828, 829, 1198, 1199, 2300, 2301, 2302, 2303, 2600, 5500, 5501
// Default: 0
func (w *AllpagesClient) Namespace(i int) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapnamespace"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// filterredir
// Which pages to list.
// Note: Due to miser mode, using this may result in fewer than aplimit results returned before continuing; in extreme cases, zero results may be returned.
// One of the following values: all, nonredirects, redirects
// Default: all
func (w *AllpagesClient) Filterredir(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapfilterredir"] = s
	})
	return w
}

// minsize
// Limit to pages with at least this many bytes.
func (w *AllpagesClient) Minsize(i int) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapminsize"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// maxsize
// Limit to pages with at most this many bytes.
func (w *AllpagesClient) Maxsize(i int) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapmaxsize"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// prtype
// Limit to protected pages only.
// Values (separate with | or alternative): edit, move, upload
func (w *AllpagesClient) Prtype(s ...string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapprtype"] = strings.Join(s, "|")
	})
	return w
}

// prlevel
// Filter protections based on protection level (must be used with apprtype= parameter).
// Values (separate with | or alternative): Can be empty, or autoconfirmed, sysop
func (w *AllpagesClient) Prlevel(s ...string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapprlevel"] = strings.Join(s, "|")
	})
	return w
}

// prfiltercascade
// Filter protections based on cascadingness (ignored when apprtype isn't set).
// One of the following values: all, cascading, noncascading
// Default: all
func (w *AllpagesClient) Prfiltercascade(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapprfiltercascade"] = s
	})
	return w
}

// limit
// How many total pages to return.
// The value must be between 1 and 500. A value <= 0 indicates a value of "max"
// Default: 10
func (w *AllpagesClient) Limit(i int) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		if i <= 0 {
			m["gaplimit"] = "max"
		} else {
			m["gaplimit"] = strconv.FormatInt(int64(i), 10)
		}
	})
	return w
}

// dir
// The direction in which to list.
// One of the following values: ascending, descending
// Default: ascending
func (w *AllpagesClient) Dir(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapdir"] = s
	})
	return w
}

// filterlanglinks
// Filter based on whether a page has langlinks. Note that this may not consider langlinks added by extensions.
// One of the following values: all, withlanglinks, withoutlanglinks
// Default: all
func (w *AllpagesClient) Filterlanglinks(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapfilterlanglinks"] = s
	})
	return w
}

// prexpiry
// Which protection expiry to filter the page on:
// indefinite - Get only pages with indefinite protection expiry.
// definite -  Get only pages with a definite (specific) protection expiry.
// all - Get pages with any protections expiry.
// One of the following values: all, definite, indefinite
// Default: all
func (w *AllpagesClient) Prexpiry(s string) *AllpagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["gapprexpiry"] = s
	})
	return w
}

func (w *AllpagesClient) Do(ctx context.Context) (AllpagesResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return AllpagesResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action":    "query",
		"generator": "allpages",
		"prop":      "info",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := AllpagesResponse{}
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
