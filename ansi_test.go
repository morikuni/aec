package aec_test

import (
	"testing"

	"github.com/morikuni/aec"
)

func TestANSIWith(t *testing.T) {
	t.Parallel()

	tests := []struct {
		doc  string
		got  aec.ANSI
		want string
	}{
		{
			doc:  "single",
			got:  aec.Bold.With(),
			want: "\x1b[1m",
		},
		{
			doc:  "multiple",
			got:  aec.Bold.With(aec.RedF, aec.Right(2)),
			want: "\x1b[1m\x1b[31m\x1b[2C",
		},
		{
			doc:  "cursor",
			got:  aec.Up(2).With(),
			want: "\x1b[2A",
		},
		{
			doc:  "cursor then sgr",
			got:  aec.Up(2).With(aec.Bold),
			want: aec.Up(2).String() + aec.Bold.String(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.doc, func(t *testing.T) {
			if got := tc.got.String(); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestANSIApply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		doc  string
		ansi aec.ANSI
		in   string
		want string
	}{
		{
			doc:  "single",
			ansi: aec.Bold,
			in:   "hello",
			want: "\x1b[1mhello\x1b[0m",
		},
		{
			doc:  "combined",
			ansi: aec.Bold.With(aec.RedF),
			in:   "hello",
			want: "\x1b[1m\x1b[31mhello\x1b[0m",
		},
		{
			doc:  "combined multiple",
			ansi: aec.Bold.With(aec.RedF, aec.Right(2)),
			in:   "hello",
			want: "\x1b[1m\x1b[31m\x1b[2Chello\x1b[0m",
		},
		{
			doc:  "empty string",
			ansi: aec.Bold,
			in:   "",
			want: "\x1b[1m\x1b[0m",
		},
	}

	for _, tc := range tests {
		t.Run(tc.doc, func(t *testing.T) {
			if got := tc.ansi.Apply(tc.in); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		doc  string
		got  string
		want string
	}{
		{
			doc:  "no ansi",
			got:  aec.Apply("hello"),
			want: "hello",
		},
		{
			doc:  "empty string",
			got:  aec.Apply("", aec.Bold),
			want: "\x1b[1m\x1b[0m",
		},
		{
			doc:  "single",
			got:  aec.Apply("hello", aec.Bold),
			want: "\x1b[1mhello\x1b[0m",
		},
		{
			doc:  "multiple",
			got:  aec.Apply("hello", aec.Bold, aec.RedF, aec.Right(2)),
			want: "\x1b[1m\x1b[31m\x1b[2Chello\x1b[0m",
		},
	}

	for _, tc := range tests {
		t.Run(tc.doc, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("got %q, want %q", tc.got, tc.want)
			}
		})
	}
}
