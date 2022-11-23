package mediawiki

import "time"

type QueryResponse struct {
	RawJSON       string                       `json:"-"`
	BatchComplete bool                         `json:"batchcomplete"`
	Error         *QueryResponseError          `json:"error,omitempty"`
	Query         *QueryResponseQuery          `json:"query,omitempty"`
	Warnings      map[string]map[string]string `json:"warnings,omitempty"`
}

type QueryResponseError struct {
	Code   string `json:"code"`
	Info   string `json:"info"`
	Docref string `json:"docref"`
}

type QueryResponseNormalized struct {
	Fromencoded bool   `json:"fromencoded"`
	From        string `json:"from"`
	To          string `json:"to"`
}

type QueryResponseQuery struct {
	Normalized []QueryResponseNormalized `json:"normalized,omitempty"`
	Pages      []QueryResponseQueryPage  `json:"pages"`
	Tokens     map[string]string         `json:"tokens,omitempty"`
}

type QueryResponseQueryPage struct {
	PageId               int                              `json:"pageid,omitempty"`
	Namespace            int                              `json:"ns"`
	Title                string                           `json:"title"`
	Revisions            []QueryResponseQueryPageRevision `json:"revisions,omitempty"`
	Missing              bool                             `json:"missing,omitempty"`
	CategoryInfo         map[string]int                   `json:"categoryinfo,omitempty"`
	Contentmodel         string                           `json:"contentmodel,omitempty"`
	Pagelanguage         string                           `json:"pagelanguage,omitempty"`
	Pagelanguagehtmlcode string                           `json:"pagelanguagehtmlcode,omitempty"`
	Pagelanguagedir      string                           `json:"pagelanguagedir,omitempty"`
	Touched              *time.Time                       `json:"touched,omitempty"`
	Lastrevid            int                              `json:"lastrevid,omitempty"`
	Length               int                              `json:"length,omitempty"`
}

type QueryResponseQueryPageRevision struct {
	Slots map[string]QueryResponseQueryPageRevisionSlot `json:"slots"`
}

type QueryResponseQueryPageRevisionSlot struct {
	Content       string `json:"content"`
	ContentModel  string `json:"contentmodel"`
	ContentFormat string `json:"contentformat"`
}

type QueryOption func(map[string]string)
