package aec_test

import (
	"testing"

	"github.com/morikuni/aec"
)

func TestNewBuilder(t *testing.T) {
	t.Parallel()

	got := aec.NewBuilder(aec.Up(2), aec.Bold).ANSI.String()
	want := aec.Up(2).With(aec.Bold).String()
	if got != want {
		t.Fatalf("builder ANSI mismatch: got %q, want %q", got, want)
	}
}

func TestEmptyBuilderChain(t *testing.T) {
	t.Parallel()

	got := aec.EmptyBuilder.Right(2).RedF().Bold().ANSI.String()
	want := aec.Right(2).With(aec.RedF, aec.Bold).String()
	if got != want {
		t.Fatalf("builder ANSI mismatch: got %q, want %q", got, want)
	}
}

func TestBuilderApply(t *testing.T) {
	t.Parallel()

	got := aec.EmptyBuilder.Bold().RedF().ANSI.Apply("hello")
	want := aec.Bold.With(aec.RedF).Apply("hello")
	if got != want {
		t.Fatalf("builder Apply mismatch: got %q, want %q", got, want)
	}
}

func TestBuilderFixtures(t *testing.T) {
	t.Parallel()

	tests := []struct {
		doc  string
		got  string
		want string
	}{
		{
			doc:  "empty",
			got:  aec.EmptyBuilder.ANSI.String(),
			want: "",
		},
		{
			doc:  "bold",
			got:  aec.EmptyBuilder.Bold().ANSI.String(),
			want: "\x1b[1m",
		},
		{
			doc:  "right then red foreground",
			got:  aec.EmptyBuilder.Right(2).RedF().ANSI.String(),
			want: "\x1b[2C\x1b[31m",
		},
		{
			doc:  "apply",
			got:  aec.EmptyBuilder.Bold().ANSI.Apply("hello"),
			want: "\x1b[1mhello\x1b[0m",
		},
	}

	for _, tc := range tests {
		t.Run(tc.doc, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("got %q, want %q", tc.got, tc.want)
			}
		})
	}
}
