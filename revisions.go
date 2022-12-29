package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
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

type RevisionsResponse struct {
	QueryResponse
	Query *RevisionsResponseQuery `json:"query,omitempty"`
}

type RevisionsResponseQuery struct {
	Normalized []RevisionsResponseNormalized `json:"normalized,omitempty"`
	Pages      []RevisionsResponsePage       `json:"pages,omitempty"`
}

type RevisionsResponseNormalized struct {
	Fromencoded bool   `json:"fromencoded"`
	From        string `json:"from"`
	To          string `json:"to"`
}

type RevisionsResponsePage struct {
	Namespace Namespace                   `json:"ns"`
	Title     string                      `json:"title,omitempty"`
	Missing   any                         `json:"missing,omitempty"`
	Pageid    int                         `json:"pageid,omitempty"`
	Revisions []RevisionsResponseRevision `json:"revisions,omitempty"`
}

type RevisionsResponseRevision struct {
	Revid         int                              `json:"revid,omitempty"`
	Parentid      int                              `json:"parentid"`
	Minor         bool                             `json:"minor"`
	User          string                           `json:"user,omitempty"`
	Userid        int                              `json:"userid,omitempty"`
	Timestamp     *time.Time                       `json:"timestamp,omitempty"`
	Size          int                              `json:"size,omitempty"`
	Sha1          string                           `json:"sha1,omitempty"`
	Roles         []string                         `json:"roles,omitempty"`
	Slots         map[string]RevisionsResponseSlot `json:"slots,omitempty"`
	Comment       string                           `json:"comment"`
	Parsedcomment string                           `json:"parsedcomment,omitempty"`
	Tags          []string                         `json:"tags,omitempty"`
}

type RevisionsResponseSlot struct {
	Size          int    `json:"size,omitempty"`
	Sha1          string `json:"sha1,omitempty"`
	Contentmodel  string `json:"contentmodel,omitempty"`
	Contentformat string `json:"contentformat,omitempty"`
	Content       string `json:"content,omitempty"`
}

type RevisionsClient struct {
	o []QueryOption
	c *Client
}

// WithQueryProp
// Which properties to get for the queried pages.
func (c *Client) Revisions() *RevisionsClient {
	return &RevisionsClient{c: c}
}

func (w *RevisionsClient) Titles(s ...string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["titles"] = strings.Join(s, "|")
	})
	return w
}

// prop
func (w *RevisionsClient) Prop(s ...string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvprop"] = strings.Join(s, "|")
	})
	return w
}

// slots
// Which revision slots to return data for, when slot-related properties are included in rvprops. If omitted, data from the main slot will be returned in a backwards-compatible format.
// Values (separate with | or alternative): main
// To specify all values, use *.
func (w *RevisionsClient) Slots(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvslots"] = s
	})
	return w
}

// limit
// Limit how many revisions will be returned.
// May only be used with a single page (mode #2).
// The value must be between 1 and 500.
func (w *RevisionsClient) Limit(i int) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvlimit"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// expandtemplates
// Use action=expandtemplates instead. Expand templates in revision content (requires rvprop=content).
func (w *RevisionsClient) Expandtemplates(b bool) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvexpandtemplates"] = strconv.FormatBool(b)
	})
	return w
}

// generatexml
// Use action=expandtemplates or action=parse instead. Generate XML parse tree for revision content (requires rvprop=content).
func (w *RevisionsClient) Generatexml(b bool) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvgeneratexml"] = strconv.FormatBool(b)
	})
	return w
}

// parse
// Use action=parse instead. Parse revision content (requires rvprop=content). For performance reasons, if this option is used, rvlimit is enforced to 1.
func (w *RevisionsClient) Parse(b bool) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvparse"] = strconv.FormatBool(b)
	})
	return w
}

// section
// Only retrieve the content of the section with this identifier.
func (w *RevisionsClient) Section(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvsection"] = s
	})
	return w
}

// diffto
// Use action=compare instead. Revision ID to diff each revision to. Use prev, next and cur for the previous, next and current revision respectively.
func (w *RevisionsClient) Diffto(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvdiffto"] = s
	})
	return w
}

// difftotext
// Use action=compare instead. Text to diff each revision to. Only diffs a limited number of revisions. Overrides rvdiffto. If rvsection is set, only that section will be diffed against this text.
func (w *RevisionsClient) Difftotext(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvdifftotext"] = s
	})
	return w
}

// difftotextpst
// Use action=compare instead. Perform a pre-save transform on the text before diffing it. Only valid when used with rvdifftotext.
func (w *RevisionsClient) Difftotextpst(b bool) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvdifftotextpst"] = strconv.FormatBool(b)
	})
	return w
}

// contentformat
// Serialization format used for rvdifftotext and expected for output of content.
// One of the following values: application/json, application/octet-stream, application/unknown, application/x-binary, text/css, text/javascript, text/plain, text/unknown, text/x-wiki, unknown/unknown
func (w *RevisionsClient) Contentformat(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvcontentformat"] = s
	})
	return w
}

// startid
// Start enumeration from this revision's timestamp. The revision must exist, but need not belong to this page.
// May only be used with a single page (mode #2).
func (w *RevisionsClient) Startid(i int) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvstartid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// endid
// Stop enumeration at this revision's timestamp. The revision must exist, but need not belong to this page.
// May only be used with a single page (mode #2).
func (w *RevisionsClient) Endid(i int) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvendid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// start
// From which revision timestamp to start enumeration.
// May only be used with a single page (mode #2).
func (w *RevisionsClient) Start(t time.Time) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvstart"] = t.Format("2006-01-02T15:04:05Z")
	})
	return w
}

// end
// Enumerate up to this timestamp.
// May only be used with a single page (mode #2).
func (w *RevisionsClient) End(t time.Time) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvend"] = t.Format("2006-01-02T15:04:05Z")
	})
	return w
}

// dir
func (w *RevisionsClient) Dir(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvdir"] = s
	})
	return w
}

// user
// Only include revisions made by user.
// May only be used with a single page (mode #2).
func (w *RevisionsClient) User(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvuser"] = s
	})
	return w
}

// excludeuser
// Exclude revisions made by user.
// May only be used with a single page (mode #2).
func (w *RevisionsClient) Excludeuser(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvexcludeuser"] = s
	})
	return w
}

// tag
// Only list revisions tagged with this tag.
func (w *RevisionsClient) Tag(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvtag"] = s
	})
	return w
}

// continue
// When more results are available, use this to continue.
func (w *RevisionsClient) Continue(s string) *RevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["rvcontinue"] = s
	})
	return w
}

func (w *RevisionsClient) Do(ctx context.Context) (RevisionsResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return RevisionsResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action":        "query",
		"prop":          "revisions",
		"formatversion": "2",
		"rvslots":       "main",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := RevisionsResponse{}
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
