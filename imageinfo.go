package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Fetch data from and about MediaWiki.
// All data modifications will first have to use query to acquire a token to prevent abuse from malicious sites.
//
// Flags:
// * This module requires read rights.

// Imageinfo

type ImageinfoResponse struct {
	CoreResponse
	BatchComplete any             `json:"batchcomplete,omitempty"`
	Query         *ImageinfoQuery `json:"query,omitempty"`
}

type ImageinfoQuery struct {
	Pages map[string]ImageinfoPage `json:"pages"`
}

type ImageinfoPage struct {
	Pageid          int                      `json:"pageid"`
	Ns              Namespace                `json:"ns"`
	Title           string                   `json:"title"`
	ImageRepository string                   `json:"imagerepository"`
	Missing         any                      `json:"missing,omitempty"`
	Imageinfo       []ImageinfoPageImageinfo `json:"imageinfo,omitempty"`
}

type ImageinfoPageImageinfo struct {
	Timestamp           *time.Time       `json:"timestamp,omitempty"`
	User                string           `json:"user"`
	UserId              int              `json:"userid,omitempty"`
	Size                int              `json:"size,omitempty"`
	Width               int              `json:"width,omitempty"`
	Height              int              `json:"height,omitempty"`
	ParsedComment       string           `json:"parsedcomment"`
	Comment             string           `json:"comment"`
	CanonicalTitle      string           `json:"canonicaltitle,omitempty"`
	Url                 string           `json:"url,omitempty"`
	DescriptionUrl      string           `json:"descriptionurl,omitempty"`
	DescriptionShortUrl string           `json:"descriptionshorturl,omitempty"`
	Sha1                string           `json:"sha1,omitempty"`
	Metadata            []map[string]any `json:"metadata,omitempty"`
	CommonMetadata      []string         `json:"commonmetadata"`
	Extmetadata         map[string]any   `json:"extmetadata,omitempty"`
	Mime                string           `json:"mime,omitempty"`
	Mediatype           string           `json:"mediatype,omitempty"`
}

type ImageinfoOption func(map[string]string)

type ImageinfoClient struct {
	o []ImageinfoOption
	c *Client
}

func (c *Client) Imageinfo() *ImageinfoClient {
	return &ImageinfoClient{c: c}
}

// iiprop
func (w *ImageinfoClient) Prop(s ...string) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iiprop"] = strings.Join(s, "|")
	})
	return w
}

// How many files to return.
// Type: integer or max
// The value must be between 1 and 500.
// Default: 10
func (w *ImageinfoClient) Limit(i int) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iilimit"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// start
// From which revision timestamp to start enumeration.
// May only be used with a single page (mode #2).
func (w *ImageinfoClient) Start(t time.Time) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iistart"] = t.Format("2006-01-02T15:04:05Z")
	})
	return w
}

// end
// Enumerate up to this timestamp.
// May only be used with a single page (mode #2).
func (w *ImageinfoClient) End(t time.Time) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iiend"] = t.Format("2006-01-02T15:04:05Z")
	})
	return w
}

// If iiprop=url is set, a URL to an image scaled to this width will be returned.
//
// For performance reasons if this option is used, no more than 50 scaled images will be returned.
//
// Type: integer
// Default: -1
func (w *ImageinfoClient) URLWidth(i int) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iiurlwidth"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Similar to iiurlwidth.
//
// Type: integer
// Default: -1
func (w *ImageinfoClient) URLHeight(i int) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iiurlheight"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Version of metadata to use. If latest is specified, use latest version.
// Defaults to 1 for backwards compatibility.
//
// Default: 1
func (w *ImageinfoClient) MetadataVersion(s string) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iimetadataversion"] = s
	})
	return w
}

// Continue
// When more results are available, use this to continue.
func (w *ImageinfoClient) Continue(s string) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["iicontinue"] = s
	})
	return w
}

// Titles. A list of titles to work on.
// Separate values with | or alternative.
// Maximum number of values is 50 (500 for clients allowed higher limits).
func (w *ImageinfoClient) Titles(s ...string) *ImageinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["titles"] = strings.Join(s, "|")
	})
	return w
}

func (w *ImageinfoClient) Do(ctx context.Context) (ImageinfoResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return ImageinfoResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"prop":   "imageinfo",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := ImageinfoResponse{}
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
