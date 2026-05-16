package noregex_test

import (
	"regexp"
	"testing"

	no "github.com/uyo9/noregex"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		tokens  []no.Token
		input   string
		matched bool
	}{
		{
			"simple literal",
			[]no.Token{no.Literally("foo")},
			"foo",
			true,
		},
		{
			"literal no match",
			[]no.Token{no.Literally("foo")},
			"bar",
			false,
		},
		{
			"digit pattern",
			[]no.Token{no.Some(no.Digit())},
			"123",
			true,
		},
		{
			"case ignored",
			[]no.Token{no.CaseIgnored(), no.Literally("foo")},
			"FOO",
			true,
		},
		{
			"multilined",
			[]no.Token{no.Multilined(), no.LineStart(), no.Literally("foo")},
			"bar\nfoo",
			true,
		},
		{
			"singlelined",
			[]no.Token{no.Singlelined(), no.Some(no.Character())},
			"foo\nbar",
			true,
		},
		{
			"anchored line literal",
			[]no.Token{no.LineStart(), no.Literally("foo"), no.LineEnd()},
			"foo",
			true,
		},
		{
			"anchored line no match",
			[]no.Token{no.LineStart(), no.Literally("foo"), no.LineEnd()},
			"foobar",
			false,
		},
		{
			"anchored word boundary",
			[]no.Token{no.Literally("foo"), no.WordBoundary()},
			"foo bar",
			true,
		},
		{
			"anchored not word boundary no match",
			[]no.Token{no.Literally("foo"), no.NotWordBoundary()},
			"foo bar",
			false,
		},
		{
			"either",
			[]no.Token{no.Either(no.Literally("foo"), no.Literally("bar"))},
			"bar",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := no.New(tt.tokens...)
			re := regexp.MustCompile(p)
			got := re.MatchString(tt.input)
			if got != tt.matched {
				t.Errorf("New(...) = %q, MatchString(%q) = %v, want %v", p, tt.input, got, tt.matched)
			}
		})
	}
}
