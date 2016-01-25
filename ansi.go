package aec

import (
	"fmt"
)

const esc = "\x1b["

// Reset resets SGR effect.
const Reset string = "\x1b[0m"

var empty = newAnsi("")

// ANSI represents ANSI escape code.
type ANSI interface {
	fmt.Stringer

	// With adapts a given ANSI.
	With(ANSI) ANSI

	// Apply wraps given string in ANSI.
	Apply(string) string
}

type ansiImpl string

func newAnsi(s string) *ansiImpl {
	r := ansiImpl(s)
	return &r
}

func (a *ansiImpl) With(ansi ANSI) ANSI {
	s := a.String() + ansi.String()
	r := ansiImpl(s)
	return &r
}

func (a *ansiImpl) Apply(s string) string {
	return a.String() + s + Reset
}

func (a *ansiImpl) String() string {
	return string(*a)
}
