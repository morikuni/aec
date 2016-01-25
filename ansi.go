package aec

const esc = "\x1b["

// Reset resets SGR effect.
const Reset string = "\x1b[0m"

var empty = newAnsi("")

// ANSI represents ANSI escape code.
type ANSI interface {
	// Code returns a ANSI escape code.
	Code() string

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

func (a *ansiImpl) Code() string {
	return string(*a)
}

func (a *ansiImpl) With(ansi ANSI) ANSI {
	s := a.Code() + ansi.Code()
	r := ansiImpl(s)
	return &r
}

func (a *ansiImpl) Apply(s string) string {
	return a.Code() + s + Reset
}

func (a *ansiImpl) String() string {
	return string(*a)
}
