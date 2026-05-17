package noregex_test

import (
	"testing"

	no "github.com/uyo9/noregex"
)

func TestRaw(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"literal", `\f`, `\f`},
		{"complex", `(?:\d+)`, `(?:\d+)`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Raw(tt.input).Token()
			if got != tt.want {
				t.Errorf("Raw(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestLiterally(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"foo", "foo", "foo"},
		{"dot", "foo.bar", `foo\.bar`},
		{"plus", "a+b", `a\+b`},
		{"star", "a*b", `a\*b`},
		{"question mark", "a?b", `a\?b`},
		{"caret", "a^b", `a\^b`},
		{"dollar", "a$b", `a\$b`},
		{"round brackets", "a(b)c", `a\(b\)c`},
		{"square brackets", "a[b]c", `a\[b\]c`},
		{"curly brackets", "a{b}c", `a\{b\}c`},
		{"backslash", `a\b`, `a\\b`},
		{"pipe", "a|b", `a\|b`},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Literally(tt.input).Token()
			if got != tt.want {
				t.Errorf("Literally(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestASCII(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"alpha", "alpha", "[[:alpha:]]"},
		{"digit", "digit", "[[:digit:]]"},
		{"space", "space", "[[:space:]]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.ASCII(tt.input).Token()
			if got != tt.want {
				t.Errorf("ASCII(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNotASCII(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"alpha", "alpha", "[[:^alpha:]]"},
		{"digit", "digit", "[[:^digit:]]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.NotASCII(tt.input).Token()
			if got != tt.want {
				t.Errorf("NotASCII(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestUnicode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"L", "L", `[\p{L}]`},
		{"Han", "Han", `[\p{Han}]`},
		{"Greek", "Greek", `[\p{Greek}]`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Unicode(tt.input).Token()
			if got != tt.want {
				t.Errorf("Unicode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNotUnicode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"L", "L", `[\P{L}]`},
		{"Han", "Han", `[\P{Han}]`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.NotUnicode(tt.input).Token()
			if got != tt.want {
				t.Errorf("NotUnicode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestOneOf(t *testing.T) {
	tests := []struct {
		name  string
		input []rune
		want  string
	}{
		{"digits", []rune{'1', '2', '3'}, "[123]"},
		{"character", []rune{'a'}, "[a]"},
		{"vowels", []rune{'a', 'e', 'i', 'o', 'u'}, "[aeiou]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.OneOf(tt.input...).Token()
			if got != tt.want {
				t.Errorf("OneOf(%q) = %q, want %q", string(tt.input), got, tt.want)
			}
		})
	}
}

func TestMaybe(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		want  string
	}{
		{"digit", no.Digit(), `(?:\d)?`},
		{"letter", no.Letter(), `(?:[a-zA-Z])?`},
		{"literal", no.Literally("foo"), `(?:foo)?`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Maybe(tt.input).Token()
			if got != tt.want {
				t.Errorf("Maybe(%q) = %q, want %q", tt.input.Token(), got, tt.want)
			}
		})
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		want  string
	}{
		{"digit", no.Digit(), `(?:\d)+`},
		{"letter", no.Letter(), `(?:[a-zA-Z])+`},
		{"literal", no.Literally("foo"), `(?:foo)+`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Some(tt.input).Token()
			if got != tt.want {
				t.Errorf("Some(%q) = %q, want %q", tt.input.Token(), got, tt.want)
			}
		})
	}
}

func TestMaybeSome(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		want  string
	}{
		{"digit", no.Digit(), `(?:\d)*`},
		{"letter", no.Letter(), `(?:[a-zA-Z])*`},
		{"literal", no.Literally("foo"), `(?:foo)*`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.MaybeSome(tt.input).Token()
			if got != tt.want {
				t.Errorf("MaybeSome(%q) = %q, want %q", tt.input.Token(), got, tt.want)
			}
		})
	}
}

func TestEither(t *testing.T) {
	tests := []struct {
		name   string
		inputs []no.Pattern
		want   string
	}{
		{"two literals", []no.Pattern{no.Literally("foo"), no.Literally("bar")}, `(?:foo|bar)`},
		{"three literals", []no.Pattern{no.Literally("foo"), no.Literally("bar"), no.Literally("baz")}, `(?:foo|bar|baz)`},
		{"digit or letter", []no.Pattern{no.Digit(), no.Letter()}, `(?:\d|[a-zA-Z])`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Either(tt.inputs...).Token()
			if got != tt.want {
				t.Errorf("Either(...) = %q, want %q", got, tt.want)
			}
		})
	}
}
