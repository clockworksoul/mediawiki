package mediawiki

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Upload a file, or get the status of pending uploads.
// Several methods are available:
//
// Upload file contents directly, using the file parameter.
// Upload the file in pieces, using the filesize, chunk, and offset parameters.
// Have the MediaWiki server fetch a file from a URL, using the url parameter.
// Complete an earlier upload that failed due to warnings, using the filekey parameter.Note that
// the HTTP POST must be done as a file upload (i.e. using multipart/form-data) when sending
// the file.
//
// Flags:
// * This module requires read rights.
// * This module requires write rights.
// * This module only accepts POST requests.

type UploadResponse struct {
	CoreResponse
	Upload *UploadUploadResponse `json:"upload,omitempty"`
}

type UploadUploadResponse struct {
	Filename   string                 `json:"filename,omitempty"`
	Imageinfo  *UploadUploadImageinfo `json:"imageinfo,omitempty"`
	Result     Result                 `json:"result,omitempty"`
	Warnings   *UploadUploadWarnings  `json:"warnings,omitempty"`
	FileKey    string                 `json:"filekey,omitempty"`
	SessionKey string                 `json:"sessionkey,omitempty"`
}

type UploadUploadWarnings struct {
	Exists           string                        `json:"exists,omitempty"`
	NoChange         *UploadUploadWarningsNoChange `json:"nochange,omitempty"`
	DuplicateVersion string                        `json:"duplicate-version,omitempty"`
	WasDeleted       string                        `json:"was-deleted,omitempty"`
	Duplicate        []string                      `json:"duplicate,omitempty"`
	DuplicateArchive string                        `json:"duplicate-archive,omitempty"`
	BadFileName      string                        `json:"badfilename,omitempty"`
}

type UploadUploadWarningsNoChange struct {
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type UploadUploadImageinfo struct {
	Bitdepth       int    `json:"bitdepth,omitempty"`
	Canonicaltitle string `json:"canonicaltitle,omitempty"`
	Comment        string `json:"comment"`
	Commonmetadata []struct {
		Name  string `json:"name,omitempty"`
		Value any    `json:"value,omitempty"`
	} `json:"commonmetadata,omitempty"`
	Descriptionurl string `json:"descriptionurl,omitempty"`
	Extmetadata    struct {
		DateTime *struct {
			Value  *time.Time `json:"value,omitempty"`
			Source string     `json:"source,omitempty"`
			Hidden any        `json:"hidden,omitempty"`
		} `json:"DateTime,omitempty"`
		ObjectName *struct {
			Value  string `json:"value,omitempty"`
			Source string `json:"source,omitempty"`
			Hidden any    `json:"hidden,omitempty"`
		} `json:"ObjectName,omitempty"`
	} `json:"extmetadata,omitempty"`
	Height    int    `json:"height,omitempty"`
	HTML      string `json:"html,omitempty"`
	Mediatype string `json:"mediatype,omitempty"`
	Metadata  []struct {
		Name  string `json:"name,omitempty"`
		Value any    `json:"value,omitempty"`
	} `json:"metadata,omitempty"`
	Mime          string     `json:"mime,omitempty"`
	Parsedcomment string     `json:"parsedcomment"`
	Sha1          string     `json:"sha1,omitempty"`
	Size          int        `json:"size,omitempty"`
	Timestamp     *time.Time `json:"timestamp,omitempty"`
	URL           string     `json:"url,omitempty"`
	User          string     `json:"user,omitempty"`
	Userid        int        `json:"userid,omitempty"`
	Width         int        `json:"width,omitempty"`
}

type UploadOption func(map[string]string)

type UploadClient struct {
	o []UploadOption
	c *Client
	f io.Reader
}

func (c *Client) Upload() *UploadClient {
	return &UploadClient{c: c}
}

// Filename
// Target filename.
func (w *UploadClient) Filename(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["filename"] = s
	})
	return w
}

// Comment
// Upload comment. Also used as the initial page text for new files if text is not specified.
// Default: (empty)
func (w *UploadClient) Comment(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["comment"] = s
	})
	return w
}

// Tags
// Change tags to apply to the upload log entry and file page revision.
// Values (separate with | or alternative): possible vandalism, repeating characters
func (w *UploadClient) Tags(s ...string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["tags"] = strings.Join(s, "|")
	})
	return w
}

// Text
// Initial page text for new files.
func (w *UploadClient) Text(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["text"] = s
	})
	return w
}

// Watch
// Watch the page.
func (w *UploadClient) Watch(b bool) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watch"] = strconv.FormatBool(b)
	})
	return w
}

// Watchlist
// Unconditionally add or remove the page from the current user's watchlist, use preferences (ignored for bot users) or do not change watch.
// One of the following values: nochange, preferences, watch
// Default: preferences
func (w *UploadClient) Watchlist(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlist"] = s
	})
	return w
}

// Watchlistexpiry
// Watchlist expiry timestamp. Omit this parameter entirely to leave the current expiry unchanged.
func (w *UploadClient) Watchlistexpiry(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["watchlistexpiry"] = s
	})
	return w
}

// Ignorewarnings
// Ignore any warnings.
func (w *UploadClient) Ignorewarnings(b bool) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["ignorewarnings"] = strconv.FormatBool(b)
	})
	return w
}

// File
// File contents.
// Must be posted as a file upload using multipart/form-data.
func (w *UploadClient) File(f io.Reader) *UploadClient {
	w.f = f
	return w
}

// Url
// URL to fetch the file from.
func (w *UploadClient) Url(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["url"] = s
	})
	return w
}

// Filekey
// Key that identifies a previous upload that was stashed temporarily.
func (w *UploadClient) Filekey(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["filekey"] = s
	})
	return w
}

// Sessionkey
// Same as filekey, maintained for backward compatibility.
func (w *UploadClient) Sessionkey(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["sessionkey"] = s
	})
	return w
}

// Stash
// If set, the server will stash the file temporarily instead of adding it to the repository.
func (w *UploadClient) Stash(b bool) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["stash"] = strconv.FormatBool(b)
	})
	return w
}

// Filesize
// Filesize of entire upload.
// The value must be between 0 and 4,294,967,296.
func (w *UploadClient) Filesize(i int) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["filesize"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Offset
// Offset of chunk in bytes.
// The value must be no less than 0.
func (w *UploadClient) Offset(i int) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["offset"] = strconv.FormatInt(int64(i), 10)
	})
	return w
}

// Chunk
// Chunk contents.
// Must be posted as a file upload using multipart/form-data.
func (w *UploadClient) Chunk(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["chunk"] = s
	})
	return w
}

// Async
// Make potentially large file operations asynchronous when possible.
func (w *UploadClient) Async(b bool) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["async"] = strconv.FormatBool(b)
	})
	return w
}

// Checkstatus
// Only fetch the upload status for the given file key.
func (w *UploadClient) Checkstatus(b bool) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["checkstatus"] = strconv.FormatBool(b)
	})
	return w
}

// Token
// A "csrf" token retrieved from action=query&meta=tokens
func (w *UploadClient) Token(s string) *UploadClient {
	w.o = append(w.o, func(m map[string]string) {
		m["token"] = s
	})
	return w
}

func (w *UploadClient) Do(ctx context.Context) (UploadResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return UploadResponse{}, err
	}

	token, err := w.c.GetToken(ctx, CSRFToken)
	if err != nil {
		return UploadResponse{}, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Specify parameters to send.
	parameters := Values{
		"action": "upload",
		"token":  token,
		"format": "json",
	}

	for _, o := range w.o {
		o(parameters)
	}

	for k, v := range parameters {
		writer.WriteField(k, v)
	}

	if w.f != nil {
		part, _ := writer.CreateFormFile("file", parameters["filename"])
		io.Copy(part, w.f)
	}

	writer.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", w.c.apiURL.String(), body)
	if err != nil {
		return UploadResponse{}, err
	}

	req.Header.Add("User-Agent", w.c.UserAgent)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := w.c.Client.Do(req)
	if err != nil {
		return UploadResponse{}, err
	}

	if resp.StatusCode >= 400 {
		return UploadResponse{}, fmt.Errorf(resp.Status)
	}

	b, _ := io.ReadAll(resp.Body)
	buf := &bytes.Buffer{}
	json.Indent(buf, b, "", "  ")
	j := buf.String()

	if w.c.Debug != nil {
		fmt.Fprintln(w.c.Debug, j)
	}

	r := UploadResponse{}
	err = ParseResponseReader(buf, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("error parsing response: %w", err)
	}

	if w.c.Debug != nil {
		fmt.Fprintln(w.c.Debug, "-----")
		b, _ = json.MarshalIndent(r, "", "  ")
		fmt.Fprintln(w.c.Debug, string(b))
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Upload == nil {
		return r, fmt.Errorf("unexpected error in upload")
	} else if r.Upload.Result != Success && r.Upload.Result != Warning {
		return r, fmt.Errorf("upload failure")
	}

	return r, nil
}
