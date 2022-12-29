package mediawiki

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
)

// If you modify this package, please change DefaultUserAgent.

// DefaultUserAgent is the HTTP User-Agent used by default.
const DefaultUserAgent = "go-mediawiki (https://github.com/clockworksoul/mediawiki)"

type Client struct {
	Client *http.Client

	apiURL *url.URL

	// HTTP user agent
	UserAgent string

	// API token cache.
	// Maps from name of token (e.g., "csrf") to token value.
	// Use GetToken to obtain tokens.
	Tokens *Tokens

	Debug io.Writer

	// Used for keep-alive
	lastLoginTime      time.Time
	username, password string
	loginBot           bool
	keepAliveMutex     sync.Mutex
}

type Token string

type Tokens struct {
	sync.RWMutex
	m map[Token]string
}

// These consts represents MW API token names.
// They are meant to be used with the GetToken method like so:
//
//	ClientInstance.GetToken(mwclient.CSRFToken)
const (
	CSRFToken                   Token = "csrf"
	DeleteGlobalAccountToken    Token = "deleteglobalaccount"
	PatrolToken                 Token = "patrol"
	RollbackToken               Token = "rollback"
	SetGlobalAccountStatusToken Token = "setglobalaccountstatus"
	UserRightsToken             Token = "userrights"
	WatchToken                  Token = "watch"
	LoginToken                  Token = "login"
)

// New returns a pointer to an initialized Client object. If the provided API URL
// is invalid (as defined by the net/url package), then it will return nil and
// the error from url.Parse().
//
// The userAgent parameter will be joined with the DefaultUserAgent const and
// used as HTTP User-Agent. If userAgent is an empty string, DefaultUserAgent
// will be used by itself as User-Agent. The User-Agent set by New can be
// overriden by setting the UserAgent field on the returned *Client.
//
// New disables maxlag by default. To enable it, simply set
// Client.Maxlag.On to true. The default timeout is 5 seconds and the default
// amount of retries is 3.
func New(inURL, userAgent string) (*Client, error) {
	apiurl, err := url.Parse(inURL)
	if err != nil {
		return nil, err
	}

	var ua string
	if userAgent != "" {
		ua = userAgent + " " + DefaultUserAgent
	} else {
		ua = DefaultUserAgent
	}

	client := &Client{}
	client.init(apiurl, ua)

	return client, nil
}

// Login attempts to login using the provided username and password.
// Do not use Login with OAuth.
func (w *Client) GetToken(ctx context.Context, token Token) (string, error) {
	if w.apiURL.String() == "" {
		return "", fmt.Errorf("api URL is undefined")
	}

	w.Tokens.Lock()
	defer w.Tokens.Unlock()

	if t, exists := w.Tokens.m[token]; exists {
		return t, nil
	}

	url := fmt.Sprintf("%s?format=json&action=query&meta=tokens&type=%s", w.apiURL.String(), token)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error requesting token: %w", err)
	}

	req.Header.Set("User-Agent", w.UserAgent)

	// if w.Debug != nil {
	// 	reqdump, err := httputil.DumpRequestOut(req, true)
	// 	if err != nil {
	// 		fmt.Fprintf(w.Debug, "Err dumping request: %v\n", err)
	// 	} else {
	// 		w.Debug.Write(reqdump)
	// 	}
	// }

	resp, err := w.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error receiving token: %w", err)
	}

	// if w.Debug != nil {
	// 	respdump, err := httputil.DumpResponse(resp, true)
	// 	if err != nil {
	// 		fmt.Fprintf(w.Debug, "Err dumping response: %v\n", err)
	// 	} else {
	// 		w.Debug.Write(respdump)
	// 	}
	// }

	r := Response{}
	err = ParseResponseReader(resp.Body, &r)
	if err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	if e := r.Error; e != nil {
		return "", fmt.Errorf("%s: %s", e.Code, e.Info)
	}

	ts := r.Query.Tokens[string(token)+"token"]
	if ts != "" {
		w.Tokens.m[token] = ts
	}

	return ts, nil
}

func (w *Client) BotLogin(ctx context.Context, username, password string) (Response, error) {
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
		return r, fmt.Errorf("login failure: (%s) %s", r.BotLogin.Result, r.BotLogin.Reason)
	}

	w.lastLoginTime = time.Now()

	return r, nil
}

func (w *Client) GetInto(ctx context.Context, v Values, a any) (string, error) {
	v["format"] = "json"

	query := w.apiURL.String() + "?" + v.Encode()

	if w.Debug != nil {
		fmt.Fprintln(w.Debug, query)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", query, nil)
	if err != nil {
		return "", fmt.Errorf("error constructing GET: %w", err)
	}

	req.Header.Set("User-Agent", w.UserAgent)

	resp, err := w.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error executing Get: %w", err)
	}

	b, _ := io.ReadAll(resp.Body)
	buf := &bytes.Buffer{}
	json.Indent(buf, b, "", "  ")
	j := buf.String()

	if w.Debug != nil {
		fmt.Fprintln(w.Debug, j)
	}

	err = ParseResponseReader(buf, a)
	if err != nil {
		return j, fmt.Errorf("error parsing response: %w", err)
	}

	if w.Debug != nil {
		fmt.Fprintln(w.Debug, "-----")
		b, _ = json.MarshalIndent(a, "", "  ")
		fmt.Fprintln(w.Debug, string(b))
	}

	return j, nil
}

func (w *Client) PostInto(ctx context.Context, v Values, a any) (string, error) {
	v["format"] = "json"

	req, err := http.NewRequestWithContext(ctx, "POST", w.apiURL.String(), strings.NewReader(v.Encode()))
	if err != nil {
		return "", fmt.Errorf("error constructing POST: %w", err)
	}

	req.Header.Set("User-Agent", w.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := w.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error executing POST: %w", err)
	}

	b, _ := io.ReadAll(resp.Body)
	buf := &bytes.Buffer{}
	json.Indent(buf, b, "", "  ")
	j := buf.String()

	if w.Debug != nil {
		fmt.Fprintln(w.Debug, j)
	}

	err = ParseResponseReader(buf, a)
	if err != nil {
		return j, fmt.Errorf("error parsing response: %w", err)
	}

	if w.Debug != nil {
		fmt.Fprintln(w.Debug, "-----")
		b, _ = json.MarshalIndent(a, "", "  ")
		fmt.Fprintln(w.Debug, string(b))
	}

	return j, nil
}

// checkKeepAlive checks for the presence of an active session cookie,
// and attempts to re-initialize the connection if one isn't found.
func (w *Client) checkKeepAlive(ctx context.Context) error {
	if w.username == "" && w.password == "" {
		return nil
	}

	w.keepAliveMutex.Lock()
	defer w.keepAliveMutex.Unlock()

	if w.hasCookie("^.*_session$") {
		return nil
	}

	if err := w.init(w.apiURL, w.UserAgent); err != nil {
		return fmt.Errorf("keep-alive re-init failure: %w", err)
	}

	if w.loginBot {
		if _, err := w.BotLogin(ctx, w.username, w.password); err != nil {
			return fmt.Errorf("keep-alive login failure: %w", err)
		}
	} else {
		if _, err := w.ClientLogin(ctx, w.username, w.password); err != nil {
			return fmt.Errorf("keep-alive login failure: %w", err)
		}
	}

	return nil
}

func (w *Client) hasCookie(regex string) bool {
	p := regexp.MustCompile(regex)

	for _, c := range w.Client.Jar.Cookies(w.apiURL) {
		if p.Match([]byte(c.Name)) {
			return true
		}
	}

	return false
}

// init is used by both New and checkKeepAlive to (re-)initialize
// the client.
func (w *Client) init(apiurl *url.URL, ua string) error {
	cookies, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	w.Client = &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           cookies,
		Timeout:       30 * time.Second,
	}
	w.apiURL = apiurl
	w.UserAgent = ua
	w.Tokens = &Tokens{m: map[Token]string{}}

	return nil
}
