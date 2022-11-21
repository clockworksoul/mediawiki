package mediawiki

import (
	"context"
	"fmt"
	"strconv"
)

// Delete a page.
//
// Flags:
// * This module requires read rights.
// * This module requires write rights.
// * This module only accepts POST requests.

type DeleteResponse struct {
	CoreResponse
	Delete *DeleteDeleteResponse `json:"delete,omitempty"`
}

type DeleteDeleteResponse struct {
	Title  string `json:"title"`
	Resaon string `json:"reason"`
	LogId  int    `json:"logid,omitempty"`
}

// Delete
type DeleteOption func(map[string]string)

type DeleteClient struct {
	o []DeleteOption
	c *Client
}

func (c *Client) Delete() *DeleteClient {
	return &DeleteClient{c: c}
}

// Title
// Title of the page to delete. Cannot be used together with pageid.
func (w *DeleteClient) Title(s string) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["title"] = s
	})
	return w
}

// Pageid
// Page ID of the page to delete. Cannot be used together with title.
func (w *DeleteClient) Pageid(i int) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["pageid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Reason
// Reason for the deletion. If not set, an automatically generated reason will be used.
func (w *DeleteClient) Reason(s string) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["reason"] = s
	})
	return w
}

// Tags
// Change tags to apply to the entry in the deletion log.
// Values (separate with | or alternative): possible vandalism, repeating characters
func (w *DeleteClient) Tags(s string) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["tags"] = s
	})
	return w
}

// Deletetalk
// Delete the talk page, if it exists.
func (w *DeleteClient) Deletetalk(b bool) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["deletetalk"] = strconv.FormatBool(b)
	})
	return w
}

// Watch
// Add the page to the current user's watchlist.
func (w *DeleteClient) Watch(b bool) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watch"] = strconv.FormatBool(b)
	})
	return w
}

// Watchlist
// Unconditionally add or remove the page from the current user's watchlist, use preferences (ignored for bot users) or do not change watch.
// One of the following values: nochange, preferences, unwatch, watch
// Default: preferences
func (w *DeleteClient) Watchlist(s string) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlist"] = s
	})
	return w
}

// Watchlistexpiry
// Watchlist expiry timestamp. Omit this parameter entirely to leave the current expiry unchanged.
func (w *DeleteClient) Watchlistexpiry(s string) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlistexpiry"] = s
	})
	return w
}

// Unwatch
// Remove the page from the current user's watchlist.
func (w *DeleteClient) Unwatch(b bool) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["unwatch"] = strconv.FormatBool(b)
	})
	return w
}

// Oldimage
// The name of the old image to delete as provided by action=query&prop=imageinfo&iiprop=archivename.
func (w *DeleteClient) Oldimage(s string) *DeleteClient {
	w.o = append(w.o, func(m map[string]string) {
		m["oldimage"] = s
	})
	return w
}

func (w *DeleteClient) Do(ctx context.Context) (DeleteResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return DeleteResponse{}, err
	}

	token, err := w.c.GetToken(ctx, CSRFToken)
	if err != nil {
		return DeleteResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "delete",
		"token":  token,
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := DeleteResponse{}
	j, err := w.c.PostInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Delete == nil {
		return r, fmt.Errorf("unexpected error in delete")
	}

	return r, nil
}
