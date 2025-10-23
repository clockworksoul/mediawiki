package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Fetch data from and about MediaWiki.
// All data modifications will first have to use query to acquire a token to prevent abuse from malicious sites.
//
// Flags:
// * This module requires read rights.

// Images

type ImagesResponse struct {
	CoreResponse
	BatchComplete any             `json:"batchcomplete,omitempty"`
	Continue      *ImagesContinue `json:"continue,omitempty"`
	Query         *ImagesQuery    `json:"query,omitempty"`
}

type ImagesContinue struct {
	Imcontinue string `json:"imcontinue"`
	Continue   string `json:"continue"`
}

type ImagesQuery struct {
	Pages map[string]ImagesPage `json:"pages"`
}

type ImagesPage struct {
	Pageid  int               `json:"pageid"`
	Ns      Namespace         `json:"ns"`
	Title   string            `json:"title"`
	Missing any               `json:"missing,omitempty"`
	Images  []ImagesPageImage `json:"images,omitempty"`
}

type ImagesPageImage struct {
	Ns    Namespace `json:"ns"`
	Title string    `json:"title"`
}

type ImagesOption func(map[string]string)

type ImagesClient struct {
	o []ImagesOption
	c *Client
}

func (c *Client) Images() *ImagesClient {
	return &ImagesClient{c: c}
}

// How many files to return.
// Type: integer or max
// The value must be between 1 and 500.
// Default: 10
func (w *ImagesClient) Limit(i int) *ImagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["imlimit"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Continue
// When more results are available, use this to continue.
func (w *ImagesClient) Continue(s string) *ImagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["imcontinue"] = s
	})
	return w
}

// Only list these files. Useful for checking whether a certain page has a certain file.
//
// Maximum number of values is 50 (500 for clients that are allowed higher limits).
func (w *ImagesClient) Images(s ...string) *ImagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["imimages"] = strings.Join(s, "|")
	})
	return w
}

// Dir
// The direction in which to list.

// One of the following values: ascending, descending
// Default: ascending
func (w *ImagesClient) Dir(s string) *ImagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["imdir"] = s
	})
	return w
}

// Titles. A list of titles to work on.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *ImagesClient) Titles(s ...string) *ImagesClient {
	w.o = append(w.o, func(m map[string]string) {
		m["titles"] = strings.Join(s, "|")
	})
	return w
}

func (w *ImagesClient) Do(ctx context.Context) (ImagesResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return ImagesResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"prop":   "images",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := ImagesResponse{}
	j, err := w.c.GetInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Query == nil {
		return r, fmt.Errorf("unexpected error in query")
	}

	return r, nil
}
