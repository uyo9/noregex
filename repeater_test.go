package noregex_test

import (
	"testing"

	no "github.com/uyo9/noregex"
)

func TestExactly(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		n     uint
		want  string
	}{
		{"digit 3 times", no.Digit(), 3, `(?:\d){3}`},
		{"letter 5 times", no.Letter(), 5, `(?:[a-zA-Z]){5}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Repeated(tt.input).Exactly(tt.n).Token()
			if got != tt.want {
				t.Errorf("Repeated(%q).Exactly(%d) = %q, want %q", tt.input.Token(), tt.n, got, tt.want)
			}
		})
	}
}

func TestAtLeast(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		n     uint
		want  string
	}{
		{"digit at least 2 times", no.Digit(), 2, `(?:\d){2,}`},
		{"letter at least 1 time", no.Letter(), 1, `(?:[a-zA-Z]){1,}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Repeated(tt.input).AtLeast(tt.n).Token()
			if got != tt.want {
				t.Errorf("Repeated(%q).AtLeast(%d) = %q, want %q", tt.input.Token(), tt.n, got, tt.want)
			}
		})
	}
}

func TestAtMost(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		n     uint
		want  string
	}{
		{"digit at most 3 times", no.Digit(), 3, `(?:\d){,3}`},
		{"letter at most 5 times", no.Letter(), 5, `(?:[a-zA-Z]){,5}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Repeated(tt.input).AtMost(tt.n).Token()
			if got != tt.want {
				t.Errorf("Repeated(%q).AtMost(%d) = %q, want %q", tt.input.Token(), tt.n, got, tt.want)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		name  string
		input no.Pattern
		n, m  uint
		want  string
	}{
		{"digit 2 to 5 times", no.Digit(), 2, 5, `(?:\d){2,5}`},
		{"letter 1 to 3 times", no.Letter(), 1, 3, `(?:[a-zA-Z]){1,3}`},
		{"digit 5 to 2 times", no.Digit(), 5, 2, `(?:\d){2,5}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := no.Repeated(tt.input).Between(tt.n, tt.m).Token()
			if got != tt.want {
				t.Errorf("Repeated(%q).Between(%d, %d) = %q, want %q", tt.input.Token(), tt.n, tt.m, got, tt.want)
			}
		})
	}
}
