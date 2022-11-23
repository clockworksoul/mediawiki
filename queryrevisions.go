package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Get revision information.
// May be used in several ways:
//
// Get data about a set of pages (last revision), by setting titles or pageids.
// Get revisions for one given page, by using titles or pageids with start, end, or limit.
// Get data about a set of revisions by setting their IDs with revids.
//
// Flags:
// * This module requires read rights.
// * This module can be used as a generator.

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

// Rvprop
func (w *QueryRevisionsClient) Prop(s ...string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvprop"] = strings.Join(s, "|")
	})
	return w
}

// Rvslots
// Which revision slots to return data for, when slot-related properties are included in rvprops. If omitted, data from the main slot will be returned in a backwards-compatible format.
// Values (separate with | or alternative): main
// To specify all values, use *.
func (w *QueryRevisionsClient) Slots(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvslots"] = s
	})
	return w
}

// Rvlimit
// Limit how many revisions will be returned.
// May only be used with a single page (mode #2).
// The value must be between 1 and 500.
func (w *QueryRevisionsClient) Limit(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvlimit"] = s
	})
	return w
}

// Rvexpandtemplates
// Use action=expandtemplates instead. Expand templates in revision content (requires rvprop=content).
func (w *QueryRevisionsClient) Expandtemplates(b bool) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvexpandtemplates"] = strconv.FormatBool(b)
	})
	return w
}

// Rvgeneratexml
// Use action=expandtemplates or action=parse instead. Generate XML parse tree for revision content (requires rvprop=content).
func (w *QueryRevisionsClient) Generatexml(b bool) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvgeneratexml"] = strconv.FormatBool(b)
	})
	return w
}

// Rvparse
// Use action=parse instead. Parse revision content (requires rvprop=content). For performance reasons, if this option is used, rvlimit is enforced to 1.
func (w *QueryRevisionsClient) Parse(b bool) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvparse"] = strconv.FormatBool(b)
	})
	return w
}

// Rvsection
// Only retrieve the content of the section with this identifier.
func (w *QueryRevisionsClient) Section(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvsection"] = s
	})
	return w
}

// Rvdiffto
// Use action=compare instead. Revision ID to diff each revision to. Use prev, next and cur for the previous, next and current revision respectively.
func (w *QueryRevisionsClient) Diffto(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvdiffto"] = s
	})
	return w
}

// Rvdifftotext
// Use action=compare instead. Text to diff each revision to. Only diffs a limited number of revisions. Overrides rvdiffto. If rvsection is set, only that section will be diffed against this text.
func (w *QueryRevisionsClient) Difftotext(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvdifftotext"] = s
	})
	return w
}

// Rvdifftotextpst
// Use action=compare instead. Perform a pre-save transform on the text before diffing it. Only valid when used with rvdifftotext.
func (w *QueryRevisionsClient) Difftotextpst(b bool) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvdifftotextpst"] = strconv.FormatBool(b)
	})
	return w
}

// Rvcontentformat
// Serialization format used for rvdifftotext and expected for output of content.
// One of the following values: application/json, application/octet-stream, application/unknown, application/x-binary, text/css, text/javascript, text/plain, text/unknown, text/x-wiki, unknown/unknown
func (w *QueryRevisionsClient) Contentformat(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvcontentformat"] = s
	})
	return w
}

// Rvstartid
// Start enumeration from this revision's timestamp. The revision must exist, but need not belong to this page.
// May only be used with a single page (mode #2).
func (w *QueryRevisionsClient) Startid(i int) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvstartid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Rvendid
// Stop enumeration at this revision's timestamp. The revision must exist, but need not belong to this page.
// May only be used with a single page (mode #2).
func (w *QueryRevisionsClient) Endid(i int) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvendid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Rvstart
// From which revision timestamp to start enumeration.
// May only be used with a single page (mode #2).
func (w *QueryRevisionsClient) Start(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvstart"] = s
	})
	return w
}

// Rvend
// Enumerate up to this timestamp.
// May only be used with a single page (mode #2).
func (w *QueryRevisionsClient) End(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvend"] = s
	})
	return w
}

// Rvdir
func (w *QueryRevisionsClient) Dir(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvdir"] = s
	})
	return w
}

// Rvuser
// Only include revisions made by user.
// May only be used with a single page (mode #2).
func (w *QueryRevisionsClient) User(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvuser"] = s
	})
	return w
}

// Rvexcludeuser
// Exclude revisions made by user.
// May only be used with a single page (mode #2).
func (w *QueryRevisionsClient) Excludeuser(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvexcludeuser"] = s
	})
	return w
}

// Rvtag
// Only list revisions tagged with this tag.
func (w *QueryRevisionsClient) Tag(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvtag"] = s
	})
	return w
}

// Rvcontinue
// When more results are available, use this to continue.
func (w *QueryRevisionsClient) Continue(s string) *QueryRevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvcontinue"] = s
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
