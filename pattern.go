package noregex

import (
	"fmt"
)

type Pattern interface {
	Token
	With(...Flag) Pattern
	Or(Pattern) Pattern
	Then(Pattern) Pattern
	Grouped() Pattern
	GroupedAs(name string) Pattern
	Lazily() Pattern
}

type pattern struct{ value string }

func (p pattern) Token() string { return p.value }

func (p pattern) With(fs ...Flag) Pattern {
	raw := ""
	for _, f := range fs {
		raw += f.Raw()
	}

	return pattern{fmt.Sprintf("(?%s:%s)", raw, p.value)}
}

func (p pattern) Or(pp Pattern) Pattern {
	return pattern{fmt.Sprintf("%s|%s", p.value, pp.Token())}
}

func (p pattern) Then(pp Pattern) Pattern {
	return pattern{fmt.Sprintf("%s%s", p.value, pp.Token())}
}

func (p pattern) Grouped() Pattern { return pattern{fmt.Sprintf("(%s)", p.value)} }

func (p pattern) GroupedAs(name string) Pattern {
	return pattern{fmt.Sprintf("(?<%s>%s)", name, p.value)}
}

func (p pattern) Lazily() Pattern { return pattern{fmt.Sprintf("%s?", p.value)} }
