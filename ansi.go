package aec

import (
	"fmt"
	"strings"
)

// Esc is the ANSI escape sequence introducer (ESC).
const Esc = "\x1b["

// Reset clears all SGR attributes.
const Reset = "\x1b[0m"

var empty = newAnsi("")

// ANSI represents an ANSI escape sequence (SGR).
type ANSI interface {
	fmt.Stringer

	// With returns a new sequence composed with the given ones.
	With(...ANSI) ANSI

	// Apply wraps s with the sequence and appends a reset.
	Apply(string) string
}

// ansiImpl represents an ANSI escape sequence (SGR).
type ansiImpl string

func newAnsi(s string) *ansiImpl {
	r := ansiImpl(s)
	return &r
}

// With returns a new ANSI sequence composed of this and the provided ones.
func (a *ansiImpl) With(ansi ...ANSI) ANSI {
	return concat(append([]ANSI{a}, ansi...))
}

// Apply wraps s with the ANSI sequence and a reset.
func (a *ansiImpl) Apply(s string) string {
	return a.String() + s + Reset
}

// String returns the ANSI escape sequence.
func (a *ansiImpl) String() string {
	return string(*a)
}

// Apply wraps s with all provided ANSI sequences and a reset.
func Apply(s string, ansi ...ANSI) string {
	if len(ansi) == 0 {
		return s
	}
	return concat(ansi).Apply(s)
}

func concat(ansi []ANSI) ANSI {
	strs := make([]string, 0, len(ansi))
	for _, p := range ansi {
		strs = append(strs, p.String())
	}
	return newAnsi(strings.Join(strs, ""))
}
