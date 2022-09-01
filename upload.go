package mediawiki

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
)

func (w *Client) Upload(ctx context.Context, name string, file io.Reader, filename string) (Response, error) {
	if err := w.checkKeepAlive(ctx); err != nil {
		return Response{}, err
	}

	token, err := w.GetToken(ctx, CSRFToken)
	if err != nil {
		return Response{}, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("action", "upload")
	writer.WriteField("filename", filename)
	writer.WriteField("ignorewarnings", "true")
	writer.WriteField("token", token)
	writer.WriteField("format", "json")

	part, _ := writer.CreateFormFile("file", filename)
	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", w.apiURL.String(), body)
	if err != nil {
		return Response{}, err
	}

	req.Header.Add("User-Agent", w.UserAgent)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	if w.Debug != nil {
		reqdump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Fprintf(w.Debug, "Err dumping request: %v\n", err)
		} else {
			w.Debug.Write(reqdump)
		}
	}

	resp, err := w.Client.Do(req)
	if err != nil {
		return Response{}, err
	}

	if w.Debug != nil {
		respdump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Fprintf(w.Debug, "Err dumping response: %v\n", err)
		} else {
			w.Debug.Write(respdump)
		}
	}

	if resp.StatusCode >= 400 {
		return Response{}, fmt.Errorf(resp.Status)
	}

	r, err := ParseResponseReader(resp.Body)
	if err != nil {
		return r, err
	}

	if e := r.Error; e != nil {
		if (e.Code == "backend-fail-alreadyexists" || e.Code == "fileexists-no-change") && !w.FailFileExists {
			return r, nil
		}

		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	} else if r.Upload == nil {
		return r, fmt.Errorf("unexpected error in upload")
	} else if r.Upload.Result != "Success" && r.Upload.Result != "Warning" {
		return r, fmt.Errorf("upload failure")
	}

	return r, nil
}
