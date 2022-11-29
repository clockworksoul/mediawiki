package mediawiki

import (
	"context"
	"fmt"
	"time"
)

// LoginOption
type LoginOption func(map[string]string)

// Format to use for returning messages.
// One of the following values: html, none, raw, wikitext
// Default: wikitext
func (w *Client) WithLoginDomain(s string) LoginOption {
	return func(m map[string]string) {
		m["lgdomain"] = s
	}
}

// Log in and get authentication cookies.
//
// This action should only be used in combination with Special:BotPasswords;
// use for main-account login is deprecated and may fail without warning.
// To safely log in to the main account, use action=clientlogin.
func (w *Client) Login(ctx context.Context, username, password string, options ...LoginOption) (Response, error) {
	token, err := w.GetToken(ctx, LoginToken)
	if err != nil {
		return Response{}, err
	}

	w.username = username
	w.password = password
	w.loginBot = true

	v := Values{
		"action":     "login",
		"lgname":     username,
		"lgpassword": password,
		"lgtoken":    token,
	}

	for _, o := range options {
		o(v)
	}

	r := Response{}
	j, err := w.PostInto(ctx, v, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("error parsing response: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.BotLogin == nil {
		return r, fmt.Errorf("unexpected error in login")
	} else if r.BotLogin.Result != Success {
		return r, fmt.Errorf("login %s: (%s) %s", r.ClientLogin.Status, r.ClientLogin.MessageCode, r.ClientLogin.Message)
	}

	w.lastLoginTime = time.Now()

	return r, nil
}
