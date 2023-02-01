package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Move a page.
//
// Flags:
// * This module requires read rights.
// * This module requires write rights.
// * This module only accepts POST requests.

// Move

type MoveResponse struct {
	CoreResponse
	Move *MoveResponseMove `json:"move,omitempty"`
}

type MoveResponseMove struct {
	From            string `json:"from"`
	To              string `json:"to"`
	Reason          string `json:"reason"`
	RedirectCreated any    `json:"redirectcreated,omitempty"`
	Talkfrom        string `json:"talkfrom,omitempty"`
	Talkto          string `json:"talkto,omitempty"`
}

type MoveOption func(map[string]string)

type MoveClient struct {
	o []MoveOption
	c *Client
}

func (c *Client) Move() *MoveClient {
	return &MoveClient{c: c}
}

// From
// Title of the page to rename. Cannot be used together with fromid.
func (w *MoveClient) From(s string) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["from"] = s
	})
	return w
}

// Fromid
// Page ID of the page to rename. Cannot be used together with from.
func (w *MoveClient) Fromid(i int) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["fromid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// To
// Title to rename the page to.
func (w *MoveClient) To(s string) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["to"] = s
	})
	return w
}

// Reason
// Reason for the rename.
// Default: (empty)
func (w *MoveClient) Reason(s string) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["reason"] = s
	})
	return w
}

// Movetalk
// Rename the talk page, if it exists.
// Default: false
func (w *MoveClient) Movetalk(b bool) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["movetalk"] = strconv.FormatBool(b)
	})
	return w
}

// Movesubpages
// Rename subpages, if applicable.
// Default: false
func (w *MoveClient) Movesubpages(b bool) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["movesubpages"] = strconv.FormatBool(b)
	})
	return w
}

// Noredirect
// Don't create a redirect.
// Default: false
func (w *MoveClient) Noredirect(b bool) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["noredirect"] = strconv.FormatBool(b)
	})
	return w
}

// Watchlist
// Unconditionally add or remove the page from the current user's watchlist, use preferences (ignored for bot users) or do not change watch.
// One of the following values: nochange, preferences, unwatch, watch
// Default: preferences
func (w *MoveClient) Watchlist(s string) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlist"] = s
	})
	return w
}

// Watchlistexpiry
// Watchlist expiry timestamp. Omit this parameter entirely to leave the current expiry unchanged.
func (w *MoveClient) Watchlistexpiry(s string) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlistexpiry"] = s
	})
	return w
}

// Ignorewarnings
// Ignore any warnings.
func (w *MoveClient) Ignorewarnings(b bool) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["ignorewarnings"] = strconv.FormatBool(b)
	})
	return w
}

// Tags
// Change tags to apply to the entry in the move log and to the null revision on the destination page.
func (w *MoveClient) Tags(s ...string) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["tags"] = strings.Join(s, "|")
	})
	return w
}

// Token
// A "csrf" token retrieved from action=query&meta=tokens
func (w *MoveClient) Token(s string) *MoveClient {
	w.o = append(w.o, func(m map[string]string) {
		m["token"] = s
	})
	return w
}

func (w *MoveClient) Do(ctx context.Context) (MoveResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return MoveResponse{}, err
	}

	token, err := w.c.GetToken(ctx, CSRFToken)
	if err != nil {
		return MoveResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "move",
		"token":  token,
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := MoveResponse{}
	j, err := w.c.PostInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
		// } else if r.Move == nil {
		// 	return r, fmt.Errorf("unexpected error in move")
	}

	return r, nil
}
