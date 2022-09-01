package mediawiki

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

type Response struct {
	RawJSON       string               `json:"-"`
	Batchcomplete string               `json:"batchcomplete"`
	BotLogin      *ResponseBotLogin    `json:"login"`
	ClientLogin   *ResponseClientLogin `json:"clientlogin"`
	Edit          *ResponseEdit        `json:"edit"`
	Error         *ResponseError       `json:"error"`
	Query         *ResponseQuery       `json:"query"`
	Upload        *ResponseUpload      `json:"upload"`
	Warnings      *ResponseWarnings    `json:"warnings"`
}

type ResponseError struct {
	Code   string `json:"code"`
	Info   string `json:"info"`
	Docref string `json:"docref"`
}

type ResponseQuery struct {
	Tokens map[string]string `json:"tokens"`
}

type ResponseWarnings struct {
	Tokens map[string]string `json:"tokens"`
}

type ResponseClientLogin struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	Messagecode string `json:"messagecode"`
}

type ResponseBotLogin struct {
	Result     string `json:"result"`
	Lguserid   int    `json:"lguserid"`
	Lgusername string `json:"lgusername"`
}

type ResponseEdit struct {
	Result       string    `json:"result"`
	Pageid       int       `json:"pageid"`
	Title        string    `json:"title"`
	Contentmodel string    `json:"contentmodel"`
	Oldrevid     int       `json:"oldrevid"`
	Newrevid     int       `json:"newrevid"`
	Newtimestamp time.Time `json:"newtimestamp"`
	Watched      string    `json:"watched"`
}

type ResponseUpload struct {
	Filename string            `json:"filename"`
	Result   string            `json:"result"`
	Warnings *ResponseWarnings `json:"warnings"`
}

func ParseResponseReader(in io.Reader) (Response, error) {
	b, err := io.ReadAll(in)
	if err != nil {
		return Response{}, err
	}

	return ParseResponse(b)
}

func ParseResponse(b []byte) (Response, error) {
	var raw string

	// Gets the raw JSON for debugging purposes
	bb := &bytes.Buffer{}
	if err := json.Indent(bb, b, "", "  "); err != nil {
		raw = string(b)
	} else {
		raw = bb.String()
	}

	var r = Response{RawJSON: raw}

	if err := json.Unmarshal(b, &r); err != nil {
		return r, err
	}

	return r, nil
}
