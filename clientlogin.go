package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ClientLoginOption
type ClientLoginOption func(map[string]string)

// Only use these authentication requests, by the id returned from action=query&meta=authmanagerinfo with amirequestsfor=login or from a previous response from this module.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *Client) WithClientLoginRequests(s ...string) ClientLoginOption {
	return func(m map[string]string) {
		m["loginrequests"] = strings.Join(s, "|")
	}
}

// Format to use for returning messages.
// One of the following values: html, none, raw, wikitext
// Default: wikitext
func (w *Client) WithClientLoginMessageFormat(s string) ClientLoginOption {
	return func(m map[string]string) {
		m["loginmessageformat"] = s
	}
}

// Merge field information for all authentication requests into one array.
// Type: boolean (details)
func (w *Client) WithClientLoginMergeRequestFields(b bool) ClientLoginOption {
	return func(m map[string]string) {
		m["loginmergerequestfields"] = strconv.FormatBool(b)
	}
}

// Preserve state from a previous failed login attempt, if possible.
// Type: boolean (details)
func (w *Client) WithClientLoginPreserveState(b bool) ClientLoginOption {
	return func(m map[string]string) {
		m["loginpreservestate"] = strconv.FormatBool(b)
	}
}

// Return URL for third-party authentication flows, must be absolute. Either this or logincontinue is required.
// Upon receiving a REDIRECT response, you will typically open a browser or web view to the specified redirecttarget URL for a third-party authentication flow. When that completes, the third party will send the browser or web view to this URL. You should extract any query or POST parameters from the URL and pass them as a logincontinue request to this API module.
func (w *Client) WithClientLoginReturnUrl(s string) ClientLoginOption {
	return func(m map[string]string) {
		m["loginreturnurl"] = s
	}
}

// This request is a continuation after an earlier UI or REDIRECT response. Either this or loginreturnurl is required.
// Type: boolean (details)
func (w *Client) WithClientLoginContinue(b bool) ClientLoginOption {
	return func(m map[string]string) {
		m["logincontinue"] = strconv.FormatBool(b)
	}
}

// This module accepts additional parameters depending on the available authentication requests.
// Use action=query&meta=authmanagerinfo with amirequestsfor=login (or a previous response
// from this module, if applicable) to determine the requests available and the fields that they use.
func (w *Client) WithClientLoginAdditionalParam(key, s string) ClientLoginOption {
	return func(m map[string]string) {
		m[key] = s
	}
}

// Log in to the wiki using the interactive flow.
func (w *Client) ClientLogin(ctx context.Context, username, password string, options ...ClientLoginOption) (Response, error) {
	token, err := w.GetToken(ctx, LoginToken)
	if err != nil {
		return Response{}, err
	}

	w.username = username
	w.password = password
	w.loginBot = false

	v := Values{
		"action":     "clientlogin",
		"username":   username,
		"password":   password,
		"logintoken": token,
	}

	for _, o := range options {
		o(v)
	}

	if v["loginreturnurl"] == "" && v["logincontinue"] == "" {
		v["loginreturnurl"] = fmt.Sprintf("%s://%s\n", w.apiURL.Scheme, w.apiURL.Host)
	}

	r := Response{}
	j, err := w.PostInto(ctx, v, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("error parsing response: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.ClientLogin == nil {
		return r, fmt.Errorf("unexpected error in login")
	} else if r.ClientLogin.Status != "PASS" {
		return r, fmt.Errorf("login %s: (%s) %s", r.ClientLogin.Status, r.ClientLogin.MessageCode, r.ClientLogin.Message)
	}

	w.lastLoginTime = time.Now()

	return r, nil
}
