package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type QueryCategoryInfoResponse struct {
	QueryResponse
	Query *QueryCategoryInfoResponseQuery `json:"query,omitempty"`
}

type QueryCategoryInfoResponseQuery struct {
	QueryResponseQuery
	Pages map[string]QueryResponseQueryPage `json:"pages"`
}

type QueryCategoryInfoResponsePage struct {
	QueryResponseQueryPage
	CategoryInfo map[string]QueryCategoryInfoResponsePagesCategoryInfo `json:"categoryinfo,omitempty"`
}

type QueryCategoryInfoResponsePagesCategoryInfo struct {
	Files   int `json:"files"`
	Pages   int `json:"pages"`
	Size    int `json:"size"`
	Subcats int `json:"subcats"`
}

type QueryCategoryinfoClient struct {
	o []QueryOption
	c *Client
}

func (c *Client) QueryCategoryInfo() *QueryCategoryinfoClient {
	return &QueryCategoryinfoClient{c: c}
}

// WithQueryProp
// Which properties to get for the queried pages.
func (w *QueryCategoryinfoClient) Prop(s ...string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["prop"] = strings.Join(s, "|")
	})
	return w
}

// WithQueryList
// Which lists to get.
func (w *QueryCategoryinfoClient) List(s ...string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["prop"] = strings.Join(s, "|")
	})
	return w
}

// WithQueryMeta
// Which metadata to get.
func (w *QueryCategoryinfoClient) Meta(s ...string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["meta"] = strings.Join(s, "|")
	})
	return w
}

// WithQueryIndexpageids
// Include an additional pageids section listing all returned page IDs.
func (w *QueryCategoryinfoClient) Indexpageids(b bool) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["indexpageids"] = strconv.FormatBool(b)
	})
	return w
}

// WithQueryExport
// Export the current revisions of all given or generated pages.
func (w *QueryCategoryinfoClient) Export(b bool) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["export"] = strconv.FormatBool(b)
	})
	return w
}

// WithQueryExportnowrap
// Return the export XML without wrapping it in an XML result (same format as Special:Export). Can only be used with query+export.
func (w *QueryCategoryinfoClient) Exportnowrap(b bool) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["exportnowrap"] = strconv.FormatBool(b)
	})
	return w
}

// WithQueryExportschema
// Target the given version of the XML dump format when exporting. Can only be used with query+export.
// One of the following values: 0.10, 0.11
// Default: 0.10
func (w *QueryCategoryinfoClient) Exportschema(s string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["exportschema"] = s
	})
	return w
}

// WithQueryIwurl
// Whether to get the full URL if the title is an interwiki link.
func (w *QueryCategoryinfoClient) Iwurl(b bool) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iwurl"] = strconv.FormatBool(b)
	})
	return w
}

// WithQueryContinue
// When more results are available, use this to continue.
func (w *QueryCategoryinfoClient) Continue(s string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["continue"] = s
	})
	return w
}

// WithQueryRawcontinue
// Return raw query-continue data for continuation.
func (w *QueryCategoryinfoClient) Rawcontinue(b bool) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rawcontinue"] = strconv.FormatBool(b)
	})
	return w
}

// WithQueryTitles
// A list of titles to work on.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *QueryCategoryinfoClient) Titles(s string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["titles"] = s
	})
	return w
}

// WithQueryPageids
// A list of page IDs to work on.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *QueryCategoryinfoClient) Pageids(s string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["pageids"] = s
	})
	return w
}

// WithQueryRevids
// A list of revision IDs to work on.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *QueryCategoryinfoClient) Revids(s string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["revids"] = s
	})
	return w
}

// WithQueryGenerator
func (w *QueryCategoryinfoClient) Generator(s string) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["generator"] = s
	})
	return w
}

// WithQueryRedirects
// Automatically resolve redirects in query+titles, query+pageids, and query+revids, and in pages returned by query+generator.
func (w *QueryCategoryinfoClient) Redirects(b bool) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["redirects"] = strconv.FormatBool(b)
	})
	return w
}

// WithQueryConverttitles
// Convert titles to other variants if necessary. Only works if the wiki's content language supports variant conversion. Languages that support variant conversion include ban, en, crh, gan, iu, kk, ku, shi, sr, tg, uz and zh.
func (w *QueryCategoryinfoClient) Converttitles(b bool) *QueryCategoryinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["converttitles"] = strconv.FormatBool(b)
	})
	return w
}

func (w *QueryCategoryinfoClient) Do(ctx context.Context) (QueryCategoryInfoResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return QueryCategoryInfoResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := QueryCategoryInfoResponse{}
	j, err := w.c.GetInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to get: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Query == nil {
		return r, fmt.Errorf("unexpected error in Do")
	}

	return r, nil
}
