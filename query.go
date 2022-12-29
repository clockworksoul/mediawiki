package mediawiki

import (
	"strings"
	"time"
)

// This contains structs that are used as response building blocks for multiple API commands.

type QueryResponse struct {
	CoreResponse
	BatchComplete any `json:"batchcomplete"`
}

type QueryResponseNormalized struct {
	Fromencoded bool   `json:"fromencoded,omitempty"`
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
	Namespace            Namespace                        `json:"ns"`
	Title                string                           `json:"title"`
	Revisions            []QueryResponseQueryPageRevision `json:"revisions,omitempty"`
	Missing              any                              `json:"missing,omitempty"`
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
	RevId     int                                           `json:"revid,omitempty"`
	ParentId  int                                           `json:"parentid"`
	User      string                                        `json:"user,omitempty"`
	Timestamp *time.Time                                    `json:"timestamp,omitempty"`
	Comment   string                                        `json:"comment"`
	Slots     map[string]QueryResponseQueryPageRevisionSlot `json:"slots,omitempty"`
}

type QueryResponseQueryPageRevisionSlot struct {
	Content       string `json:"content"`
	ContentModel  string `json:"contentmodel"`
	ContentFormat string `json:"contentformat"`
}

// Namespace and full page title (including all subpage levels).
func (p QueryResponseQueryPage) FullPageName() string {
	return strings.ReplaceAll(p.Title, "_", " ")
}

// Full page title (including all subpage levels) without the namespace.
func (p QueryResponseQueryPage) PageName() string {
	title := strings.ReplaceAll(p.Title, "_", " ")

	if p.Namespace == NamespaceMain {
		return title
	}

	if i := strings.Index(title, ":"); i >= 0 {
		return title[i+1:]
	}

	return title
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
	title := strings.ReplaceAll(p.Title, "_", " ")

	switch {
	case p.Namespace%2 == 0:
		return title
	case p.Namespace == NamespaceTalk:
		return strings.Replace(title, "Talk:", "", 1)
	default:
		return strings.Replace(title, " talk:", ":", 1)
	}
}

// Full page name of the associated talk page.
func (p QueryResponseQueryPage) TalkPageName() string {
	title := strings.ReplaceAll(p.Title, "_", " ")

	switch {
	case p.Namespace%2 == 1:
		return title
	case p.Namespace == 0:
		return "Talk:" + title
	default:
		return strings.Replace(title, ":", " talk:", 1)
	}
}

type QueryOption func(map[string]string)
