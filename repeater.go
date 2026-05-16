package noregex

import (
	"fmt"
)

type Repeater struct{ value string }

func Repeated(p Pattern) Repeater { return Repeater{fmt.Sprintf("(?:%s)", p.Token())} }

func (r Repeater) Exactly(n uint) Pattern { return pattern{fmt.Sprintf("%s{%d}", r.value, n)} }

func (r Repeater) AtLeast(n uint) Pattern { return pattern{fmt.Sprintf("%s{%d,}", r.value, n)} }

func (r Repeater) AtMost(n uint) Pattern { return pattern{fmt.Sprintf("%s{,%d}", r.value, n)} }

func (r Repeater) Between(n, m uint) Pattern {
	return pattern{fmt.Sprintf("%s{%d,%d}", r.value, min(n, m), max(n, m))}
}
