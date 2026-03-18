package aec

import "fmt"

// EraseMode is listed in a variable EraseModes.
type EraseMode uint

// EraseModes is a list of EraseMode.
var EraseModes = struct {
	// All erase all.
	All EraseMode

	// Head erase to head.
	Head EraseMode

	// Tail erase to tail.
	Tail EraseMode
}{
	Tail: 0,
	Head: 1,
	All:  2,
}

var (
	// Save saves the cursor position. It uses both SCO ("ESC[s") and DEC
	// ("ESC7") sequences as those were never standardized as part of the
	// ANSI.
	Save ANSI = newAnsi(esc + "s" + "\x1b7")

	// Restore restores the cursor position. It uses both SCO ("ESC[u") and
	// DEC ("ESC8") sequences as those were never standardized as part of
	// the ANSI.
	Restore ANSI = newAnsi(esc + "u" + "\x1b8")

	// Hide hides the cursor.
	Hide ANSI = newAnsi(esc + "?25l")

	// Show shows the cursor.
	Show ANSI = newAnsi(esc + "?25h")

	// Report reports the cursor position.
	Report ANSI = newAnsi(esc + "6n")
)

// Up moves the cursor up by n positions (CUU).
func Up(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dA", n))
}

// Down moves the cursor down by n positions (CUD).
func Down(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dB", n))
}

// Right moves the cursor right by n positions (CUF).
func Right(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dC", n))
}

// Left moves the cursor left by n positions (CUB).
func Left(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dD", n))
}

// NextLine moves the cursor down n lines and to the beginning of the line (CNL).
func NextLine(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dE", n))
}

// PreviousLine moves the cursor up n lines and to the beginning of the line (CPL).
func PreviousLine(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dF", n))
}

// Column sets the cursor position to a given column (CHA).
func Column(col uint) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dG", col))
}

// Position sets the cursor position to a given absolute position (CUP).
func Position(row, col uint) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%d;%dH", row, col))
}

// EraseDisplay erases the display using the given EraseMode (ED).
func EraseDisplay(m EraseMode) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dJ", m))
}

// EraseLine erases the line using the given EraseMode (EL).
func EraseLine(m EraseMode) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dK", m))
}

// ScrollUp scrolls the display up by n lines (SU).
func ScrollUp(n int) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dS", n))
}

// ScrollDown scrolls the display down by n lines (SD).
func ScrollDown(n int) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dT", n))
}
