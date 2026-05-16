package noregex

import (
	"fmt"
	"regexp"
	"strings"
)

func Literally(literal string) Pattern { return pattern{regexp.QuoteMeta(literal)} }

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
