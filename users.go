package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Get information about a list of users.
//
// Flags:
// * This module requires read rights.

// Users
type UsersOption func(map[string]string)

type UsersResponse struct {
	QueryResponse
	Query *UsersQuery `json:"query,omitempty"`
}

type UsersQuery struct {
	Users []UsersResponseUser `json:"users"`
}

type UsersResponseUser struct {
	UserId  int    `json:"userid,omitempty"`
	Name    string `json:"name,omitempty"`
	Missing any    `json:"missing,omitempty"`
}

type UsersClient struct {
	o []QueryOption
	c *Client
}

func (c *Client) Users() *UsersClient {
	return &UsersClient{c: c}
}

// Usprop
func (w *UsersClient) Prop(s ...string) *UsersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["usprop"] = strings.Join(s, "|")
	})
	return w
}

// Usattachedwiki
// With usprop=centralids, indicate whether the user is attached with the wiki identified by this ID.
func (w *UsersClient) Attachedwiki(s string) *UsersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["usattachedwiki"] = s
	})
	return w
}

// Ususers
// A list of users to obtain information for.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *UsersClient) Users(s ...string) *UsersClient {
	w.o = append(w.o, func(m map[string]string) {
		m["ususers"] = strings.Join(s, "|")
	})
	return w
}

// Ususerids
// A list of user IDs to obtain information for.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *UsersClient) Userids(i ...int) *UsersClient {
	w.o = append(w.o, func(m map[string]string) {
		var s []string

		for _, n := range i {
			s = append(s, strconv.FormatInt(int64(n), 10))
		}

		m["ususerids"] = strings.Join(s, "|")
	})
	return w
}

func (w *UsersClient) Do(ctx context.Context) (UsersResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return UsersResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"list":   "users",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := UsersResponse{}
	j, err := w.c.PostInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	}

	return r, nil
}
