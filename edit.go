package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type EditResponse struct {
	CoreResponse
	Edit *EditEditResponse `json:"edit,omitempty"`
}

type EditEditResponse struct {
	Result       string     `json:"result"`
	PageId       int        `json:"pageid"`
	Title        string     `json:"title"`
	ContentModel string     `json:"contentmodel"`
	OldRevId     int        `json:"oldrevid,omitempty"`
	NewRevId     int        `json:"newrevid,omitempty"`
	NewTimestamp *time.Time `json:"newtimestamp,omitempty"`
	Watched      any        `json:"watched,omitempty"`
	NoChange     any        `json:"nochange,omitempty"`
}

type EditOption func(map[string]string)

type EditClient struct {
	o []EditOption
	c *Client
}

func (c *Client) Edit() *EditClient {
	return &EditClient{c: c}
}

// Title
// Title of the page to edit. Cannot be used together with EditPageId.
func (w *EditClient) Title(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["title"] = s
	})
	return w
}

// PageId
// Page ID of the page to edit. Cannot be used together with title.
// Type: integer
func (w *EditClient) PageId(i int) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["pageid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Section
// Section identifier. 0 for the top section, new for a new section. Often a positive integer, but can also be non-numeric.
func (w *EditClient) Section(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["section"] = s
	})
	return w
}

// SectionTitle
// The title for a new section when using section=new.
func (w *EditClient) SectionTitle(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["sectiontitle"] = s
	})
	return w
}

// Text
// Page content.
func (w *EditClient) Text(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["text"] = s
	})
	return w
}

// Summary
// Edit summary.
// When this parameter is not provided or empty, an edit summary may be generated automatically.
// When using section=new and sectiontitle is not provided, the value of this parameter is used for the section title instead, and an edit summary is generated automatically.
func (w *EditClient) Summary(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["summary"] = s
	})
	return w
}

// Tags
// Change tags to apply to the revision.
// Values (separate with | or alternative): possible vandalism, repeating characters
func (w *EditClient) Tags(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["tags"] = s
	})
	return w
}

// Minor
// Mark this edit as a minor edit.
// Type: boolean (details)
func (w *EditClient) Minor(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["minor"] = strconv.FormatBool(b)
	})
	return w
}

// NotMinor
// Do not mark this edit as a minor edit even if the "Mark all edits minor by default" user preference is set.
// Type: boolean (details)
func (w *EditClient) NotMinor(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["motminor"] = strconv.FormatBool(b)
	})
	return w
}

// Bot
// Mark this edit as a bot edit.
// Type: boolean (details)
func (w *EditClient) Bot(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["bot"] = strconv.FormatBool(b)
	})
	return w
}

// BaseRevId
// ID of the base revision, used to detect edit conflicts. May be obtained through action=query&prop=revisions. Self-conflicts cause the edit to fail unless basetimestamp is set.
// Type: integer
func (w *EditClient) BaseRevId(i int) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["baserevid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// BaseTimestamp
// Timestamp of the base revision, used to detect edit conflicts. May be obtained through action=query&prop=revisions&rvprop=timestamp. Self-conflicts are ignored.
// Type: timestamp (allowed formats)
func (w *EditClient) BaseTimestamp(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["basetimestamp"] = s
	})
	return w
}

// StartTimestamp
// Timestamp when the editing process began, used to detect edit conflicts. An appropriate value may be obtained using curtimestamp when beginning the edit process (e.g. when loading the page content to edit).
// Type: timestamp (allowed formats)
func (w *EditClient) StartTimestamp(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["starttimestamp"] = s
	})
	return w
}

// Recreate
// Override any errors about the page having been deleted in the meantime.
// Type: boolean (details)
func (w *EditClient) Recreate(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["recreate"] = strconv.FormatBool(b)
	})
	return w
}

// CreateOnly
// Don't edit the page if it exists already.
// Type: boolean (details)
func (w *EditClient) CreateOnly(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["createonly"] = strconv.FormatBool(b)
	})
	return w
}

// NoCreate
// Throw an error if the page doesn't exist.
// Type: boolean (details)
func (w *EditClient) NoCreate(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["nocreate"] = strconv.FormatBool(b)
	})
	return w
}

// Watch
// Deprecated.
// Add the page to the current user's watchlist.
// Type: boolean (details)
func (w *EditClient) Watch(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watch"] = strconv.FormatBool(b)
	})
	return w
}

// Unwatch
// Deprecated.
// Remove the page from the current user's watchlist.
// Type: boolean (details)
func (w *EditClient) Unwatch(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["unwatch"] = strconv.FormatBool(b)
	})
	return w
}

// Watchlist
// Unconditionally add or remove the page from the current user's watchlist, use preferences (ignored for bot users) or do not change watch.
// One of the following values: nochange, preferences, unwatch, watch
// Default: preferences
func (w *EditClient) Watchlist(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlist"] = s
	})
	return w
}

// WatchListExpiry
// Watchlist expiry timestamp. Omit this parameter entirely to leave the current expiry unchanged.
// Type: expiry (details)
func (w *EditClient) WatchListExpiry(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlistexpiry"] = s
	})
	return w
}

// Md5
// The MD5 hash of the text parameter, or the prependtext and appendtext parameters concatenated. If set, the edit won't be done unless the hash is correct.
func (w *EditClient) Md5(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["md5"] = s
	})
	return w
}

// PrependText
// Add this text to the beginning of the page or section. Overrides text.
func (w *EditClient) PrependText(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["prependtext"] = s
	})
	return w
}

// AppendText
// Add this text to the end of the page or section. Overrides text.
// Use section=new to append a new section, rather than this parameter.
func (w *EditClient) AppendText(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["appendtext"] = s
	})
	return w
}

// Undo
// Undo this revision. Overrides text, prependtext and appendtext.
// Type: integer
// The value must be no less than 0.
func (w *EditClient) Undo(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["undo"] = s
	})
	return w
}

// UndoAfter
// Undo all revisions from undo to this one. If not set, just undo one revision.
// Type: integer
// The value must be no less than 0.
func (w *EditClient) UndoAfter(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["undoafter"] = s
	})
	return w
}

// Redirect
// Automatically resolve redirects.
// Type: boolean (details)
func (w *EditClient) Redirect(b bool) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["redirect"] = strconv.FormatBool(b)
	})
	return w
}

// ContentFormat
// Content serialization format used for the input text.
// One of the following values: application/json, application/octet-stream, application/unknown, application/x-binary, text/css, text/javascript, text/plain, text/unknown, text/x-wiki, unknown/unknown
func (w *EditClient) ContentFormat(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["contentformat"] = s
	})
	return w
}

// ContentModel
// Content model of the new content.
// One of the following values: GadgetDefinition, Json.JsonConfig, JsonSchema, Map.JsonConfig, MassMessageListContent, NewsletterContent, Scribunto, SecurePoll, Tabular.JsonConfig, css, flow-board, javascript, json, sanitized-css, text, translate-messagebundle, unknown, wikitext
func (w *EditClient) ContentModel(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["contentmodel"] = s
	})
	return w
}

// CaptchaWord
// Answer to the CAPTCHA
func (w *EditClient) CaptchaWord(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["captchaword"] = s
	})
	return w
}

// CaptchaId
// CAPTCHA ID from previous request
func (w *EditClient) CaptchaId(s string) *EditClient {
	w.o = append(w.o, func(m map[string]string) {
		m["captchaid"] = s
	})
	return w
}

func (w *EditClient) Do(ctx context.Context) (EditResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return EditResponse{}, err
	}

	token, err := w.c.GetToken(ctx, CSRFToken)
	if err != nil {
		return EditResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "edit",
		"token":  token,
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := EditResponse{}
	j, err := w.c.PostInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Edit == nil {
		return r, fmt.Errorf("unexpected error in edit")
	} else if r.Edit.Result != "Success" {
		return r, fmt.Errorf("write %s: (%s) %s", r.ClientLogin.Status, r.ClientLogin.MessageCode, r.ClientLogin.Message)
	}

	return r, nil
}
