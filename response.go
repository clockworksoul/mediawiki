package mediawiki

import (
	"encoding/json"
	"io"
	"time"
)

type CoreResponse struct {
	RawJSON     string                   `json:"-"`
	ClientLogin *ResponseClientLogin     `json:"clientlogin,omitempty"`
	Error       *ResponseError           `json:"error,omitempty"`
	Warnings    map[string]ResponseError `json:"warnings,omitempty"`
}

type Response struct {
	RawJSON       string               `json:"-"`
	BatchComplete interface{}          `json:"batchcomplete"`
	BotLogin      *ResponseBotLogin    `json:"login"`
	ClientLogin   *ResponseClientLogin `json:"clientlogin"`
	Edit          *ResponseEdit        `json:"edit"`
	Error         *ResponseError       `json:"error"`
	Query         *ResponseQuery       `json:"query"`
	Upload        *ResponseUpload      `json:"upload"`
	Warnings      *ResponseWarnings    `json:"warnings"`
}

type ResponseError struct {
	Code   string `json:"code,omitempty"`
	Docref string `json:"docref,omitempty"`
	Info   string `json:"info,omitempty"`
	Star   string `json:"*,omitempty"`
}

type ResponseQuery struct {
	Pages  []ResponseQueryPage `json:"pages"`
	Tokens map[string]string   `json:"tokens"`
}

type ResponseQueryPage struct {
	PageId       int                         `json:"pageid"`
	Namespace    int                         `json:"ns"`
	Title        string                      `json:"title"`
	Revisions    []ResponseQueryPageRevision `json:"revisions,omitempty"`
	Missing      string                      `json:"missing,omitempty"`
	CategoryInfo map[string]int              `json:"categoryinfo,omitempty"`
}

type ResponseQueryPageRevision struct {
	Slots map[string]ResponseQueryPageRevisionSlot `json:"slots"`
}

type ResponseQueryPageRevisionSlot struct {
	Content       string `json:"content"`
	ContentModel  string `json:"contentmodel"`
	ContentFormat string `json:"contentformat"`
}

type ResponseWarnings struct {
	Tokens map[string]string `json:"tokens"`
}

type ResponseClientLogin struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	MessageCode string `json:"messagecode"`
}

type ResponseBotLogin struct {
	Result   string `json:"result"`
	UserId   int    `json:"lguserid"`
	UserName string `json:"lgusername"`
}

type ResponseEdit struct {
	Result       string    `json:"result"`
	PageId       int       `json:"pageid"`
	Title        string    `json:"title"`
	ContentModel string    `json:"contentmodel"`
	OldRevId     int       `json:"oldrevid,omitempty"`
	NewRevId     int       `json:"newrevid,omitempty"`
	NewTimestamp time.Time `json:"newtimestamp,omitempty"`
	Watched      string    `json:"watched"`
}

type ResponseUpload struct {
	Filename string            `json:"filename"`
	Result   string            `json:"result"`
	Warnings *ResponseWarnings `json:"warnings"`
}

func ParseResponseReader(in io.Reader, v any) error {
	b, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	return ParseResponse(b, v)
}

func ParseResponse(b []byte, v any) error {
	return json.Unmarshal(b, v)
}
