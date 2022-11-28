package mediawiki

import (
	"strings"
	"time"
)

type QueryResponse struct {
	CoreResponse
	BatchComplete any                 `json:"batchcomplete"`
	Query         *QueryResponseQuery `json:"query,omitempty"`
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

// Namespace and full page title (including all subpage levels).
func (p QueryResponseQueryPage) FullPageName() string {
	return p.Title
}

// Full page title (including all subpage levels) without the namespace.
func (p QueryResponseQueryPage) PageName() string {
	if p.Namespace == NamespaceMain {
		return p.Title
	}

	if i := strings.Index(p.Title, ":"); i >= 0 {
		return p.Title[i+1:]
	}

	return p.Title
}

// Page title of the page in the immediately superior subpage level without the namespace. Would return Title/Foo on page Help:Title/Foo/Bar.
func (p QueryResponseQueryPage) BasePageName() string {
	name := p.PageName()
	split := strings.Split(name, "/")

	if len(split) == 1 {
		return name
	}

	return strings.Join(split[:len(split)-1], "/")
}

// Name of the root of the current page. Would return Title on page Help:Title/Foo/Bar.
func (p QueryResponseQueryPage) RootPageName() string {
	name := p.PageName()
	split := strings.Split(name, "/")
	return split[0]
}

// The subpage title. Would return Bar on page Help:Title/Foo/Bar. If no subpage exists the value of PageName() is returned.
func (p QueryResponseQueryPage) SubPageName() string {
	name := p.PageName()
	split := strings.Split(name, "/")
	return split[len(split)-1]
}

// Full page name of the associated subject (e.g. article or file). Useful on talk pages.
func (p QueryResponseQueryPage) ArticlePageName() string {
	switch {
	case p.Namespace%2 == 0:
		return p.Title
	case p.Namespace == NamespaceTalk:
		return strings.Replace(p.Title, "Talk:", "", 1)
	default:
		return strings.Replace(p.Title, " talk:", ":", 1)
	}
}

// Full page name of the associated talk page.
func (p QueryResponseQueryPage) TalkPageName() string {
	switch {
	case p.Namespace%2 == 1:
		return p.Title
	case p.Namespace == 0:
		return "Talk:" + p.Title
	default:
		return strings.Replace(p.Title, ":", " talk:", 1)
	}
}

type QueryOption func(map[string]string)
