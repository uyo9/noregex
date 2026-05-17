package noregex

import (
	"fmt"
	"regexp"
	"strings"
)

func Raw(s string) Pattern { return pattern{s} }

func Literally(literal string) Pattern { return pattern{regexp.QuoteMeta(literal)} }

func ASCII(name string) Pattern    { return pattern{fmt.Sprintf("[[:%s:]]", name)} }
func NotASCII(name string) Pattern { return pattern{fmt.Sprintf("[[:^%s:]]", name)} }

func Unicode(name string) Pattern    { return pattern{fmt.Sprintf(`[\p{%s}]`, name)} }
func NotUnicode(name string) Pattern { return pattern{fmt.Sprintf(`[\P{%s}]`, name)} }

func OneOf(chars ...rune) Pattern { return pattern{fmt.Sprintf("[%s]", string(chars))} }

func Maybe(p Pattern) Pattern { return pattern{fmt.Sprintf("(?:%s)?", p.Token())} }

func Some(p Pattern) Pattern { return pattern{fmt.Sprintf("(?:%s)+", p.Token())} }

func MaybeSome(p Pattern) Pattern { return pattern{fmt.Sprintf("(?:%s)*", p.Token())} }

func Either(ps ...Pattern) Pattern {
	parts := make([]string, len(ps))
	for i, p := range ps {
		parts[i] = p.Token()
	}

	return pattern{fmt.Sprintf("(?:%s)", strings.Join(parts, "|"))}
}
