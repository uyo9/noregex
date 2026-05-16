package noregex_test

import (
	"testing"

	no "github.com/uyo9/noregex"
)

func TestOr(t *testing.T) {
	tests := []struct {
		name string
		a, b no.Pattern
		want string
	}{
		{"two literals", no.Literally("foo"), no.Literally("bar"), "foo|bar"},
		{"literal or digit", no.Literally("foo"), no.Digit(), `foo|\d`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Or(tt.b).Token()
			if got != tt.want {
				t.Errorf("Or(%q, %q) = %q, want %q", tt.a.Token(), tt.b.Token(), got, tt.want)
			}
		})
	}
}

func TestThen(t *testing.T) {
	tests := []struct {
		name string
		a, b no.Pattern
		want string
	}{
		{"two literals", no.Literally("foo"), no.Literally("bar"), "foobar"},
		{"literal then digit", no.Literally("foo"), no.Digit(), `foo\d`},
		{"digit then letter", no.Digit(), no.Letter(), `\d[a-zA-Z]`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Then(tt.b).Token()
			if got != tt.want {
				t.Errorf("Then(%q, %q) = %q, want %q", tt.a.Token(), tt.b.Token(), got, tt.want)
			}
		})
	}
}

func TestGrouped(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		want  string
	}{
		{"literal", no.Literally("foo"), "(foo)"},
		{"digit", no.Digit(), `(\d)`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Grouped().Token()
			if got != tt.want {
				t.Errorf("Grouped(%q) = %q, want %q", tt.input.Token(), got, tt.want)
			}
		})
	}
}

func TestGroupedAs(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		group string
		want  string
	}{
		{"literal", no.Literally("foo"), "bar", "(?<bar>foo)"},
		{"digit", no.Digit(), "foo", `(?<foo>\d)`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.GroupedAs(tt.group).Token()
			if got != tt.want {
				t.Errorf("GroupedAs(%q, %q) = %q, want %q", tt.input.Token(), tt.group, got, tt.want)
			}
		})
	}
}

func TestLazily(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		want  string
	}{
		{"some digit", no.Some(no.Digit()), `(?:\d)+?`},
		{"maybe letter", no.Maybe(no.Letter()), `(?:[a-zA-Z])??`},
		{"maybe some literal", no.MaybeSome(no.Literally("foo")), `(?:foo)*?`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Lazily().Token()
			if got != tt.want {
				t.Errorf("Lazily(%q) = %q, want %q", tt.input.Token(), got, tt.want)
			}
		})
	}
}
