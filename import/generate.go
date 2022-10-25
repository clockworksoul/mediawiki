package main

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

func Generate(m Module) string {
	b := &bytes.Buffer{}

	b.WriteString(`package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

`)

	for _, d := range strings.Split(m.Description, "\n") {
		fmt.Fprintf(b, "// %s\n", d)
	}

	if len(m.Flags) > 0 {
		b.WriteString("//\n// Flags:\n")
		for _, f := range m.Flags {
			b.WriteString("// * ")
			b.WriteString(f)
			b.WriteRune('\n')
		}
	}

	b.WriteRune('\n')

	caser := cases.Title(language.English)
	name := caser.String(m.Name)

	fmt.Fprintln(b, "//", name)
	fmt.Fprintf(b, "type %sOption func(map[string]string)\n\n", name)

	for _, p := range m.Parameters {
		if p.Name == "*" {
			fmt.Fprintln(b, writeAdditionalParameter(m, p))
			continue
		}

		switch p.Type {
		case Boolean:
			fmt.Fprintln(b, writeBooleanParameter(m, p))
		case String:
			fmt.Fprintln(b, writeStringParameter(m, p))
		}
	}

	return b.String()
}

func writeHeaders(m Module, p *Param) string {
	b := &bytes.Buffer{}

	fmt.Fprintf(b, "// With%s%s\n", caser.String(m.Name), caser.String(p.Name))
	for _, d := range strings.Split(p.Description, "\n") {
		fmt.Fprintf(b, "// %s\n", d)
	}

	return b.String()
}

func writeBooleanParameter(m Module, p *Param) string {
	b := &bytes.Buffer{}
	mn := caser.String(m.Name)
	pn := caser.String(p.Name)

	fmt.Fprint(b, writeHeaders(m, p))

	fmt.Fprintf(b, `func (w *Client) With%s%s(b bool) %sOption {
	return func(m map[string]string) {
		m["%s"] = strconv.FormatBool(b)
	}
}
`, mn, pn, mn, p.Name)

	return b.String()
}

func writeStringParameter(m Module, p *Param) string {
	b := &bytes.Buffer{}
	mn := caser.String(m.Name)
	pn := caser.String(p.Name)

	fmt.Fprint(b, writeHeaders(m, p))

	fmt.Fprintf(b, `func (w *Client) With%s%s(s string) %sOption {
	return func(m map[string]string) {
		m["%s"] = s
	}
}
`, mn, pn, mn, p.Name)

	return b.String()
}

func writeAdditionalParameter(m Module, p *Param) string {
	b := &bytes.Buffer{}
	mn := caser.String(m.Name)

	fmt.Fprintf(b, "// With%sAdditionalParam\n", caser.String(m.Name))
	for _, d := range strings.Split(p.Description, "\n") {
		fmt.Fprintf(b, "// %s\n", d)
	}

	fmt.Fprintf(b, `func (w *Client) With%sAdditionalParam(key, s string) %sOption {
	return func(m map[string]string) {
		m[key] = s
	}
}
`, mn, mn)

	return b.String()
}
