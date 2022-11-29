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
	CoreResponse
	RawJSON       string               `json:"-"`
	BatchComplete any                  `json:"batchcomplete,omitempty"`
	BotLogin      *ResponseBotLogin    `json:"login,omitempty"`
	ClientLogin   *ResponseClientLogin `json:"clientlogin,omitempty"`
	Edit          *ResponseEdit        `json:"edit,omitempty"`
	Query         *ResponseQuery       `json:"query,omitempty"`
	Warnings      *ResponseWarnings    `json:"warnings,omitempty"`
}

type ResponseError struct {
	Code        string `json:"code,omitempty"`
	Docref      string `json:"docref,omitempty"`
	Info        string `json:"info,omitempty"`
	Stasherrors []struct {
		Message string   `json:"message,omitempty"`
		Params  []string `json:"params,omitempty"`
		Code    string   `json:"code,omitempty"`
		Type    string   `json:"type,omitempty"`
	} `json:"stasherrors,omitempty"`
	Star string `json:"*,omitempty"`
}

type ResponseQuery struct {
	Pages  []QueryResponseQueryPage `json:"pages"`
	Tokens map[string]string        `json:"tokens"`
}

type ResponseWarnings struct {
	Tokens map[string]string `json:"tokens"`
}

type ResponseClientLogin struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	MessageCode string `json:"messagecode"`
}

type Result string

const (
	Error   Result = "Error"
	Success Result = "Success"
	Warning Result = "Warning"
)

type ResponseBotLogin struct {
	Result   Result `json:"result"`
	UserId   int    `json:"lguserid"`
	UserName string `json:"lgusername"`
}

type ResponseEdit struct {
	Result       Result    `json:"result"`
	PageId       int       `json:"pageid"`
	Title        string    `json:"title"`
	ContentModel string    `json:"contentmodel"`
	OldRevId     int       `json:"oldrevid,omitempty"`
	NewRevId     int       `json:"newrevid,omitempty"`
	NewTimestamp time.Time `json:"newtimestamp,omitempty"`
	Watched      string    `json:"watched"`
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
