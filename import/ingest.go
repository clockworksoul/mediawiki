package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type TextMode int

const (
	None TextMode = iota
	ModuleName
	ModuleFlag
	ModuleFlags
	ModuleDescription
	ModuleHelpUrl
	Parameters
	Var
	Description
	Info
)

func attrMap(attr []html.Attribute) map[string]string {
	m := map[string]string{}

	for _, a := range attr {
		m[a.Key] = a.Val
	}

	return m
}

func getModulePage(module string) (io.ReadCloser, error) {
	url := fmt.Sprintf("https://www.mediawiki.org/w/api.php?action=help&modules=%s", module)

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return r.Body, nil
}

func parseModulePage(r io.ReadCloser) (Module, error) {
	defer r.Close()

	module := Module{}
	mode := None
	z := html.NewTokenizer(r)
	b := strings.Builder{}
	p := &Param{}

	for i := 1; ; i++ {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return module, nil

		case html.DoctypeToken:
		case html.SelfClosingTagToken:
			// No-op

		case html.EndTagToken:
			t := z.Token()

			if mode == None {
				continue
			}

			switch {
			case mode == ModuleFlag && t.Data == "li":
				mode = ModuleFlags
			case mode == ModuleFlags && t.Data == "div":
				b.Reset()
				mode = ModuleDescription
			case mode == Var && t.Data == "dt":
				p = &Param{Name: strings.TrimSpace(b.String()), Type: String}
				module.Parameters = append(module.Parameters, p)
				b.Reset()
				mode = Parameters
			case mode == Description && t.Data == "dd":
				p.Description = strings.TrimSpace(b.String())
				b.Reset()
				mode = Parameters
			case mode == Info && t.Data == "dd":
				s := b.String()

				switch {
				case s == "Deprecated.":
					p.Deprecated = true
				case strings.Contains(s, "equired"):
					p.Required = true
				case strings.HasPrefix(s, "Type: "):
					s := s[6:]
					if i := strings.Index(s, " ("); i != -1 {
						s = s[:i]
					}

					switch s {
					case "boolean":
						p.Type = Boolean
					case "expiry":
						p.Type = Expiry
					case "integer":
						p.Type = Integer
					default:
						return Module{}, fmt.Errorf("unhandled type for parameter %s: %s", p.Name, s)
					}
				default:
					b.Reset()
					b.WriteString(p.Description)
					b.WriteRune('\n')
					b.WriteString(s)
					p.Description = b.String()
				}

				b.Reset()
				mode = Parameters
			}

		case html.StartTagToken:
			t := z.Token()
			m := attrMap(t.Attr)

			switch {
			case t.Data == "h2" && m["class"] == "apihelp-header apihelp-module-name":
				module.Name = m["id"]
			case t.Data == "span" && strings.HasPrefix(m["class"], "apihelp-flag"):
				mode = ModuleFlag
			case m["class"] == "apihelp-block apihelp-help-urls":
				module.Description = strings.TrimSpace(b.String())
				b.Reset()
				mode = ModuleHelpUrl
			case m["class"] == "apihelp-block apihelp-parameters":
				mode = Parameters
			case m["class"] == "apihelp-block apihelp-examples":
				mode = None
			case mode != None && t.Data == "dt":
				b.Reset()
				mode = Var
			case mode != None && t.Data == "dd" && m["class"] == "description":
				b.Reset()
				mode = Description
			case mode != None && t.Data == "dd" && m["class"] == "info":
				b.Reset()
				mode = Info
			}

		case html.TextToken:
			text := string(z.Text())

			if mode == ModuleFlag {
				module.Flags = append(module.Flags, text)
			} else if mode != None {
				b.WriteString(text)
			}
		}
	}
}
