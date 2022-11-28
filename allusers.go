package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// https://www.mediawiki.org/wiki/API:Allusers
// Get information about a list of users.
//
// Flags:
// * This module requires read rights.

type AllUsersDir string

const (
	AllusersAscending  AllUsersDir = "ascending"
	AllusersDescending AllUsersDir = "descending"
)

// Allusers
type AllusersOption func(map[string]string)

type AllusersResponse struct {
	QueryResponse
	Query    *AllusersQuery            `json:"query,omitempty"`
	Continue *AllusersResponseContinue `json:"continue,omitempty"`
}

type AllusersResponseContinue struct {
	From     string `json:"aufrom"`
	Continue string `json:"continue"`
}

type AllusersQuery struct {
	Allusers []AllusersResponseAllusers `json:"allusers"`
}

type AllusersResponseAllusers struct {
	UserId int    `json:"userid,omitempty"`
	Name   string `json:"name,omitempty"`
}

type AllusersClient struct {
	o []QueryOption
	c *Client
}

func (c *Client) Allusers() *AllusersClient {
	return &AllusersClient{c: c}
}

// From
// The username to start enumerating from.
func (w *AllusersClient) From(s string) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["aufrom"] = s
	})
	return w
}

// To
// The username to stop enumerating at.
func (w *AllusersClient) To(s string) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["auto"] = s
	})
	return w
}

// Prefix
// Search for all users that begin with this value.
func (w *AllusersClient) Prefix(s string) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["auprefix"] = s
	})
	return w
}

// Dir
// Direction to sort in.
// One of the following values: ascending, descending
// Default: ascending
func (w *AllusersClient) Dir(d AllUsersDir) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["audir"] = string(d)
	})
	return w
}

// Group
// Only include users in the given groups.
// Values: accountcreator, autopatrolled, bot, bureaucrat, checkuser, confirmed,
// flow-bot, import, interface-admin, ipblock-exempt, no-ipinfo,
// push-subscription-manager, steward, suppress, sysop, translationadmin,
// transwiki, uploader
func (w *AllusersClient) Group(s ...string) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["augroup"] = strings.Join(s, "|")
	})
	return w
}

// ExcludeGroup
// Exclude users in the given groups.
// Values: accountcreator, autopatrolled, bot, bureaucrat, checkuser, confirmed,
// flow-bot, import, interface-admin, ipblock-exempt, no-ipinfo,
// push-subscription-manager, steward, suppress, sysop, translationadmin,
// transwiki, uploader
func (w *AllusersClient) ExcludeGroup(s ...string) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["auexcludegroup"] = strings.Join(s, "|")
	})
	return w
}

// Rights
// Only include users with the given rights. Does not include rights granted by implicit or
// auto-promoted groups like *, user, or autoconfirmed.
func (w *AllusersClient) Rights(s ...string) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["auexcludegroup"] = strings.Join(s, "|")
	})
	return w
}

// Prop
// Which pieces of information to include:
// Values (separate with | or alternative): blockinfo, centralids, editcount, groups, implicitgroups, registration, rights
func (w *AllusersClient) Prop(s ...string) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["auprop"] = strings.Join(s, "|")
	})
	return w
}

// Limit
// A list of users to obtain information for.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *AllusersClient) Limit(i int) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["aulimit"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Witheditsonly
// Only list users who have made edits.
// Delete the talk page, if it exists.
func (w *AllusersClient) Witheditsonly(b bool) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["auwitheditsonly"] = strconv.FormatBool(b)
	})
	return w
}

// Activeusers
// Only list users active in the last 30 days.
func (w *AllusersClient) Activeusers(b bool) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["auactiveusers"] = strconv.FormatBool(b)
	})
	return w
}

// Attachedwiki
// With usprop=centralids, indicate whether the user is attached with the wiki identified by this ID.
func (w *AllusersClient) Attachedwiki(s string) *AllusersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["auattachedwiki"] = s
	})
	return w
}

func (w *AllusersClient) Do(ctx context.Context) (AllusersResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return AllusersResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"list":   "allusers",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := AllusersResponse{}
	j, err := w.c.PostInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
		// } else if r.Allusers == nil {
		// 	return r, fmt.Errorf("unexpected error in queryusers")
	}

	return r, nil
}
