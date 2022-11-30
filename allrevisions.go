package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// List all revisions.
//
// Flags:
// * This module requires read rights.
// * This module can be used as a generator.

// Allrevisions

type AllrevisionsResponse struct {
	CoreResponse
	Batchcomplete any `json:"batchcomplete,omitempty"`
	Continue      *struct {
		Arvcontinue string `json:"arvcontinue,omitempty"`
		Continue    string `json:"continue,omitempty"`
	} `json:"continue,omitempty"`
	Query *struct {
		Allrevisions []struct {
			PageId    int `json:"pageid,omitempty"`
			Revisions []struct {
				RevId     int        `json:"revid,omitempty"`
				ParentId  *int       `json:"parentid,omitempty"`
				Parsetree string     `json:"parsetree,omitempty"`
				User      string     `json:"user,omitempty"`
				Timestamp *time.Time `json:"timestamp,omitempty"`
				Comment   string     `json:"comment,omitempty"`
			} `json:"revisions,omitempty"`
			Namespace int    `json:"ns"`
			Title     string `json:"title,omitempty"`
		} `json:"allrevisions,omitempty"`
	} `json:"query,omitempty"`
}

type AllrevisionsOption func(map[string]string)

type AllrevisionsClient struct {
	o []AllrevisionsOption
	c *Client
}

func (c *Client) Allrevisions() *AllrevisionsClient {
	return &AllrevisionsClient{c: c}
}

// prop
func (w *AllrevisionsClient) Prop(s ...string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvprop"] = strings.Join(s, "|")
	})
	return w
}

// slots
// Which revision slots to return data for, when slot-related properties are included in arvprops. If omitted, data from the main slot will be returned in a backwards-compatible format.
// To specify all values, use *.
func (w *AllrevisionsClient) Slots(s ...string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvslots"] = strings.Join(s, "|")
	})
	return w
}

// limit
// Limit how many revisions will be returned.
// The value must be between 1 and 500. A value of <= 0 indicates "max"
func (w *AllrevisionsClient) Limit(i int) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		s := "max"
		if i > 0 {
			s = strconv.FormatInt(int64(i), 10)
		}

		m["arvlimit"] = s
	})
	return w
}

// expandtemplates
// Use action=expandtemplates instead. Expand templates in revision content (requires arvprop=content).
func (w *AllrevisionsClient) Expandtemplates(b bool) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvexpandtemplates"] = strconv.FormatBool(b)
	})
	return w
}

// generatexml
// Use action=expandtemplates or action=parse instead. Generate XML parse tree for revision content (requires arvprop=content).
func (w *AllrevisionsClient) Generatexml(b bool) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvgeneratexml"] = strconv.FormatBool(b)
	})
	return w
}

// parse
// Use action=parse instead. Parse revision content (requires arvprop=content). For performance reasons, if this option is used, arvlimit is enforced to 1.
func (w *AllrevisionsClient) Parse(b bool) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvparse"] = strconv.FormatBool(b)
	})
	return w
}

// section
// Only retrieve the content of the section with this identifier.
func (w *AllrevisionsClient) Section(s string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvsection"] = s
	})
	return w
}

// diffto
// Use action=compare instead. Revision ID to diff each revision to. Use prev, next and cur for the previous, next and current revision respectively.
func (w *AllrevisionsClient) Diffto(s string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvdiffto"] = s
	})
	return w
}

// difftotext
// Use action=compare instead. Text to diff each revision to. Only diffs a limited number of revisions. Overrides arvdiffto. If arvsection is set, only that section will be diffed against this text.
func (w *AllrevisionsClient) Difftotext(s string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvdifftotext"] = s
	})
	return w
}

// difftotextpst
// Use action=compare instead. Perform a pre-save transform on the text before diffing it. Only valid when used with arvdifftotext.
func (w *AllrevisionsClient) Difftotextpst(b bool) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvdifftotextpst"] = strconv.FormatBool(b)
	})
	return w
}

// contentformat
// Serialization format used for arvdifftotext and expected for output of content.
// One of the following values: application/json, application/octet-stream, application/unknown, application/x-binary, text/css, text/javascript, text/plain, text/unknown, text/x-wiki, unknown/unknown
func (w *AllrevisionsClient) Contentformat(s string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvcontentformat"] = s
	})
	return w
}

// user
// Only list revisions by this user.
func (w *AllrevisionsClient) User(s string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvuser"] = s
	})
	return w
}

// namespace
// Only list pages in this namespace.
// Note: Due to miser mode, using this may result in fewer than limit results returned before continuing; in extreme cases, zero results may be returned.
// To specify all values, use a value of less than 0.
func (w *AllrevisionsClient) Namespace(i ...int) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		var s []string

		for _, n := range i {
			if n < 0 {
				s = append(s, "*")
			} else {
				s = append(s, strconv.FormatInt(int64(n), 10))
			}
		}

		m["ususerids"] = strings.Join(s, "|")
	})
	return w
}

// start
// The timestamp to start enumerating from.
func (w *AllrevisionsClient) Start(t time.Time) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvstart"] = t.Format("2006-01-02T15:04:05Z")
	})
	return w
}

// end
// The timestamp to stop enumerating at.
func (w *AllrevisionsClient) End(t time.Time) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvend"] = t.Format("2006-01-02T15:04:05Z")
	})
	return w
}

// dir
func (w *AllrevisionsClient) Dir(s string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvdir"] = s
	})
	return w
}

// excludeuser
// Don't list revisions by this user.
func (w *AllrevisionsClient) Excludeuser(s string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvexcludeuser"] = s
	})
	return w
}

// continue
// When more results are available, use this to continue.
func (w *AllrevisionsClient) Continue(s string) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvcontinue"] = s
	})
	return w
}

// generatetitles
// When being used as a generator, generate titles rather than revision IDs.
func (w *AllrevisionsClient) Generatetitles(b bool) *AllrevisionsClient {
	w.o = append(w.o, func(m map[string]string) {
		m["arvgeneratetitles"] = strconv.FormatBool(b)
	})
	return w
}

func (w *AllrevisionsClient) Do(ctx context.Context) (AllrevisionsResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return AllrevisionsResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"list":   "allrevisions",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := AllrevisionsResponse{}
	j, err := w.c.PostInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Query == nil {
		return r, fmt.Errorf("unexpected error in allrevisions")
	}

	return r, nil
}
