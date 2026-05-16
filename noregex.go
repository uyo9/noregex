package noregex

import (
	"strings"
)

func New(tokens ...Token) string {
	var builder strings.Builder
	for _, t := range tokens {
		builder.WriteString(t.Token())
	}

	return builder.String()
}
