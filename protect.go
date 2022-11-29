package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Change the protection level of a page.
//
// Flags:
// * This module requires read rights.
// * This module requires write rights.
// * This module only accepts POST requests.

// Protect

type ProtectResponse struct {
	CoreResponse
	Protect *struct {
		Title       string `json:"title,omitempty"`
		Reason      string `json:"reason,omitempty"`
		Protections []struct {
			Edit   string `json:"edit,omitempty"`
			Expiry string `json:"expiry,omitempty"`
		} `json:"protections,omitempty"`
	} `json:"protect,omitempty"`
}

type ProtectOption func(map[string]string)

type ProtectClient struct {
	o []ProtectOption
	c *Client
}

func (c *Client) Protect() *ProtectClient {
	return &ProtectClient{c: c}
}

// Title
// Title of the page to (un)protect. Cannot be used together with pageid.
func (w *ProtectClient) Title(s string) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["title"] = s
	})
	return w
}

// Pageid
// ID of the page to (un)protect. Cannot be used together with title.
func (w *ProtectClient) Pageid(i int) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["pageid"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Protections
// List of protection levels, formatted action=level (e.g. edit=sysop). A level of all means everyone is allowed to take the action, i.e. no restriction.
// Note: Any actions not listed will have restrictions removed.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *ProtectClient) Protections(s string) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["protections"] = s
	})
	return w
}

// Expiry
// Expiry timestamps. If only one timestamp is set, it'll be used for all protections. Use infinite, indefinite, infinity, or never, for a never-expiring protection.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
// Default: infinite
func (w *ProtectClient) Expiry(s string) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["expiry"] = s
	})
	return w
}

// Reason
// Reason for (un)protecting.
// Default: (empty)
func (w *ProtectClient) Reason(s string) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["reason"] = s
	})
	return w
}

// Tags
// Change tags to apply to the entry in the protection log.
func (w *ProtectClient) Tags(s ...string) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["tags"] = strings.Join(s, "|")
	})
	return w
}

// Cascade
// Enable cascading protection (i.e. protect transcluded templates and images used in this page). Ignored if none of the given protection levels support cascading.
func (w *ProtectClient) Cascade(b bool) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["cascade"] = strconv.FormatBool(b)
	})
	return w
}

// Watch
// If set, add the page being (un)protected to the current user's watchlist.
func (w *ProtectClient) Watch(b bool) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watch"] = strconv.FormatBool(b)
	})
	return w
}

// Watchlist
// Unconditionally add or remove the page from the current user's watchlist, use preferences (ignored for bot users) or do not change watch.
// One of the following values: nochange, preferences, unwatch, watch
// Default: preferences
func (w *ProtectClient) Watchlist(s string) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlist"] = s
	})
	return w
}

// Watchlistexpiry
// Watchlist expiry timestamp. Omit this parameter entirely to leave the current expiry unchanged.
func (w *ProtectClient) Watchlistexpiry(s string) *ProtectClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlistexpiry"] = s
	})
	return w
}

func (w *ProtectClient) Do(ctx context.Context) (ProtectResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return ProtectResponse{}, err
	}

	token, err := w.c.GetToken(ctx, CSRFToken)
	if err != nil {
		return ProtectResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "protect",
		"token":  token,
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := ProtectResponse{}
	j, err := w.c.PostInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Protect == nil {
		return r, fmt.Errorf("unexpected error in protect")
	}

	return r, nil
}
