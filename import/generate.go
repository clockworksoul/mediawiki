package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

func Generate(m Module) (string, error) {
	b := &bytes.Buffer{}

	imports, err := gatherImports(m)
	if err != nil {
		return "", err
	}

	fmt.Fprintln(b, `package mediawiki

import (`)
	for _, i := range imports {
		fmt.Fprintf(b, "\t%q\n", i)
	}

	fmt.Fprint(b, ")\n\n")

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
		case Integer:
			fmt.Fprintln(b, writeIntegerParameter(m, p))
		case Expiry, String:
			fmt.Fprintln(b, writeStringParameter(m, p))
		default:
			return "", fmt.Errorf("unsupported parameter type: %s", p.Type)
		}
	}

	fmt.Fprintf(b,
		`func (w *Client) %s(ctx context.Context, options ...%sOption) (Response, error) {
	if err := w.checkKeepAlive(ctx); err != nil {
		return Response{}, err
	}

	token, err := w.GetToken(ctx, CSRFToken)
	if err != nil {
		return Response{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "%s",
		"token":  token,
	}

	for _, o := range options {
		o(parameters)
	}

	// Make the request.
	r, err := w.Post(ctx, parameters)
	if err != nil {
		return r, fmt.Errorf("failed to post: %%w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%%s: %%s", e.Code, e.Info)
	} else if r.%s == nil {
		return r, fmt.Errorf("unexpected error in %s")
	}

	return r, nil
}
`, name, name, m.Name, name, m.Name)

	return b.String(), nil
}

func gatherImports(mod Module) ([]string, error) {
	m := map[string]interface{}{
		"context": true,
		"fmt":     true,
	}

	for _, p := range mod.Parameters {
		switch p.Type {
		case Boolean:
			m["strconv"] = true
		case Integer:
			m["strconv"] = true
		case Expiry, String:
		default:
			return nil, fmt.Errorf("unsupported parameter type: %s", p.Type)
		}
	}

	var imps []string
	for k, _ := range m {
		imps = append(imps, k)
	}
	sort.Strings(imps)

	return imps, nil
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

func writeIntegerParameter(m Module, p *Param) string {
	b := &bytes.Buffer{}
	mn := caser.String(m.Name)
	pn := caser.String(p.Name)

	fmt.Fprint(b, writeHeaders(m, p))

	fmt.Fprintf(b, `func (w *Client) With%s%s(i int) %sOption {
	return func(m map[string]string) {
		m["%s"] = strconv.FormatInt(int64(i), 10)
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
