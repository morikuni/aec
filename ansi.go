package aec

import (
	"fmt"
	"strings"
)

const esc = "\x1b["

// Reset resets SGR effect.
const Reset = "\x1b[0m"

var empty = newAnsi("")

// ANSI represents ANSI escape code.
type ANSI interface {
	fmt.Stringer

	// With adapts given ANSIs.
	With(...ANSI) ANSI

	// Apply wraps given string in ANSI.
	Apply(string) string
}

// ansiImpl represents an ANSI escape code.
type ansiImpl string

func newAnsi(s string) *ansiImpl {
	r := ansiImpl(s)
	return &r
}

// With returns a new ANSICode sequence composed of this and the provided ANSI codes.
func (a *ansiImpl) With(ansi ...ANSI) ANSI {
	return concat(append([]ANSI{a}, ansi...))
}

// Apply wraps the given string with the ANSI sequence and a reset code.
func (a *ansiImpl) Apply(s string) string {
	return a.String() + s + Reset
}

// String returns the ANSICode escape code as a string.
func (a *ansiImpl) String() string {
	return string(*a)
}

// Apply wraps the given string with all provided ANSI sequences.
func Apply(s string, ansi ...ANSI) string {
	if len(ansi) == 0 {
		return s
	}
	return concat(ansi).Apply(s)
}

// concat combines multiple ANSI codes into a single ANSI sequence.
func concat(ansi []ANSI) ANSI {
	strs := make([]string, 0, len(ansi))
	for _, p := range ansi {
		strs = append(strs, p.String())
	}
	return newAnsi(strings.Join(strs, ""))
}
