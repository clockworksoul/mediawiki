package mediawiki

import (
	"context"
	"fmt"
	"strconv"
)

type EditOption func(map[string]string)

// WithEditTitle
// Title of the page to edit. Cannot be used together with EditPageId.
func (w *Client) WithEditTitle(s string) EditOption {
	return func(m map[string]string) {
		m["title"] = s
	}
}

// WithEditPageId
// Page ID of the page to edit. Cannot be used together with title.
// Type: integer
func (w *Client) WithEditPageId(i int) EditOption {
	return func(m map[string]string) {
		m["pageid"] = strconv.FormatInt(int64(i), 10)
	}
}

// WithEditSection
// Section identifier. 0 for the top section, new for a new section. Often a positive integer, but can also be non-numeric.
func (w *Client) WithEditSection(s string) EditOption {
	return func(m map[string]string) {
		m["section"] = s
	}
}

// WithEditSectionTitle
// The title for a new section when using section=new.
func (w *Client) WithEditSectionTitle(s string) EditOption {
	return func(m map[string]string) {
		m["sectiontitle"] = s
	}
}

// WithEditText
// Page content.
func (w *Client) WithEditText(s string) EditOption {
	return func(m map[string]string) {
		m["text"] = s
	}
}

// WithEditSummary
// Edit summary.
// When this parameter is not provided or empty, an edit summary may be generated automatically.
// When using section=new and sectiontitle is not provided, the value of this parameter is used for the section title instead, and an edit summary is generated automatically.
func (w *Client) WithEditSummary(s string) EditOption {
	return func(m map[string]string) {
		m["summary"] = s
	}
}

// WithEditTags
// Change tags to apply to the revision.
// Values (separate with | or alternative): possible vandalism, repeating characters
func (w *Client) WithEditTags(s string) EditOption {
	return func(m map[string]string) {
		m["tags"] = s
	}
}

// WithEditMinor
// Mark this edit as a minor edit.
// Type: boolean (details)
func (w *Client) WithEditMinor(b bool) EditOption {
	return func(m map[string]string) {
		m["minor"] = strconv.FormatBool(b)
	}
}

// WithEditNotMinor
// Do not mark this edit as a minor edit even if the "Mark all edits minor by default" user preference is set.
// Type: boolean (details)
func (w *Client) WithEditNotMinor(b bool) EditOption {
	return func(m map[string]string) {
		m["motminor"] = strconv.FormatBool(b)
	}
}

// WithEditBot
// Mark this edit as a bot edit.
// Type: boolean (details)
func (w *Client) WithEditBot(b bool) EditOption {
	return func(m map[string]string) {
		m["bot"] = strconv.FormatBool(b)
	}
}

// WithEditBaseRevId
// ID of the base revision, used to detect edit conflicts. May be obtained through action=query&prop=revisions. Self-conflicts cause the edit to fail unless basetimestamp is set.
// Type: integer
func (w *Client) WithEditBaseRevId(i int) EditOption {
	return func(m map[string]string) {
		m["baserevid"] = strconv.FormatInt(int64(i), 10)
	}
}

// WithEditBaseTimestamp
// Timestamp of the base revision, used to detect edit conflicts. May be obtained through action=query&prop=revisions&rvprop=timestamp. Self-conflicts are ignored.
// Type: timestamp (allowed formats)
func (w *Client) WithEditBaseTimestamp(s string) EditOption {
	return func(m map[string]string) {
		m["basetimestamp"] = s
	}
}

// WithEditStartTimestamp
// Timestamp when the editing process began, used to detect edit conflicts. An appropriate value may be obtained using curtimestamp when beginning the edit process (e.g. when loading the page content to edit).
// Type: timestamp (allowed formats)
func (w *Client) WithEditStartTimestamp(s string) EditOption {
	return func(m map[string]string) {
		m["starttimestamp"] = s
	}
}

// WithEditRecreate
// Override any errors about the page having been deleted in the meantime.
// Type: boolean (details)
func (w *Client) WithEditRecreate(b bool) EditOption {
	return func(m map[string]string) {
		m["recreate"] = strconv.FormatBool(b)
	}
}

// WithEditCreateOnly
// Don't edit the page if it exists already.
// Type: boolean (details)
func (w *Client) WithEditCreateOnly(b bool) EditOption {
	return func(m map[string]string) {
		m["createonly"] = strconv.FormatBool(b)
	}
}

// WithEditNoCreate
// Throw an error if the page doesn't exist.
// Type: boolean (details)
func (w *Client) WithEditNoCreate(b bool) EditOption {
	return func(m map[string]string) {
		m["nocreate"] = strconv.FormatBool(b)
	}
}

// WithEditWatch
// Deprecated.
// Add the page to the current user's watchlist.
// Type: boolean (details)
func (w *Client) WithEditWatch(b bool) EditOption {
	return func(m map[string]string) {
		m["watch"] = strconv.FormatBool(b)
	}
}

// WithEditUnwatch
// Deprecated.
// Remove the page from the current user's watchlist.
// Type: boolean (details)
func (w *Client) WithEditUnwatch(b bool) EditOption {
	return func(m map[string]string) {
		m["unwatch"] = strconv.FormatBool(b)
	}
}

// WithEditWatchlist
// Unconditionally add or remove the page from the current user's watchlist, use preferences (ignored for bot users) or do not change watch.
// One of the following values: nochange, preferences, unwatch, watch
// Default: preferences
func (w *Client) WithEditWatchlist(s string) EditOption {
	return func(m map[string]string) {
		m["watchlist"] = s
	}
}

// WithEditWatchListExpiry
// Watchlist expiry timestamp. Omit this parameter entirely to leave the current expiry unchanged.
// Type: expiry (details)
func (w *Client) WithEditWatchListExpiry(s string) EditOption {
	return func(m map[string]string) {
		m["watchlistexpiry"] = s
	}
}

// WithEditMd5
// The MD5 hash of the text parameter, or the prependtext and appendtext parameters concatenated. If set, the edit won't be done unless the hash is correct.
func (w *Client) WithEditMd5(s string) EditOption {
	return func(m map[string]string) {
		m["md5"] = s
	}
}

// WithEditPrependText
// Add this text to the beginning of the page or section. Overrides text.
func (w *Client) WithEditPrependText(s string) EditOption {
	return func(m map[string]string) {
		m["prependtext"] = s
	}
}

// WithEditAppendText
// Add this text to the end of the page or section. Overrides text.
// Use section=new to append a new section, rather than this parameter.
func (w *Client) WithEditAppendText(s string) EditOption {
	return func(m map[string]string) {
		m["appendtext"] = s
	}
}

// WithEditUndo
// Undo this revision. Overrides text, prependtext and appendtext.
// Type: integer
// The value must be no less than 0.
func (w *Client) WithEditUndo(s string) EditOption {
	return func(m map[string]string) {
		m["undo"] = s
	}
}

// WithEditUndoAfter
// Undo all revisions from undo to this one. If not set, just undo one revision.
// Type: integer
// The value must be no less than 0.
func (w *Client) WithEditUndoAfter(s string) EditOption {
	return func(m map[string]string) {
		m["undoafter"] = s
	}
}

// WithEditRedirect
// Automatically resolve redirects.
// Type: boolean (details)
func (w *Client) WithEditRedirect(b bool) EditOption {
	return func(m map[string]string) {
		m["redirect"] = strconv.FormatBool(b)
	}
}

// WithEditContentFormat
// Content serialization format used for the input text.
// One of the following values: application/json, application/octet-stream, application/unknown, application/x-binary, text/css, text/javascript, text/plain, text/unknown, text/x-wiki, unknown/unknown
func (w *Client) WithEditContentFormat(s string) EditOption {
	return func(m map[string]string) {
		m["contentformat"] = s
	}
}

// WithEditContentModel
// Content model of the new content.
// One of the following values: GadgetDefinition, Json.JsonConfig, JsonSchema, Map.JsonConfig, MassMessageListContent, NewsletterContent, Scribunto, SecurePoll, Tabular.JsonConfig, css, flow-board, javascript, json, sanitized-css, text, translate-messagebundle, unknown, wikitext
func (w *Client) WithEditContentModel(s string) EditOption {
	return func(m map[string]string) {
		m["contentmodel"] = s
	}
}

// WithEditCaptchaWord
// Answer to the CAPTCHA
func (w *Client) WithEditCaptchaWord(s string) EditOption {
	return func(m map[string]string) {
		m["captchaword"] = s
	}
}

// WithEditCaptchaId
// CAPTCHA ID from previous request
func (w *Client) WithEditCaptchaId(s string) EditOption {
	return func(m map[string]string) {
		m["captchaid"] = s
	}
}

func (w *Client) Edit(ctx context.Context, options ...EditOption) (Response, error) {
	if err := w.checkKeepAlive(ctx); err != nil {
		return Response{}, err
	}

	token, err := w.GetToken(ctx, CSRFToken)
	if err != nil {
		return Response{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "edit",
		"token":  token,
	}

	for _, o := range options {
		o(parameters)
	}

	// Make the request.
	r, err := w.Post(ctx, parameters)
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Edit == nil {
		return r, fmt.Errorf("unexpected error in edit")
	} else if r.Edit.Result != "Success" {
		return r, fmt.Errorf("write %s: (%s) %s", r.ClientLogin.Status, r.ClientLogin.Messagecode, r.ClientLogin.Message)
	}

	return r, nil
}
