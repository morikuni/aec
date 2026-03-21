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

// Cursor control and reporting sequences.
//
// Includes CSI sequences defined by ECMA-48 / ISO 6429, DEC private modes,
// and legacy DEC/SCO save and restore sequences.
var (
	// Cursor save and restore sequences.
	//
	// These use both SCO ("ESC[s"/"ESC[u") and DEC ("ESC 7"/"ESC 8")
	// sequences for compatibility. They are not standardized in ECMA-48.
	Save    ANSI = newAnsi(esc + "s" + "\x1b7") // saves the cursor position.
	Restore ANSI = newAnsi(esc + "u" + "\x1b8") // restores the cursor position.

	// Cursor visibility (DEC private mode 25).
	Hide ANSI = newAnsi(esc + "?25l") // hides the cursor.
	Show ANSI = newAnsi(esc + "?25h") // shows the cursor.

	// Cursor position report (DSR 6, ECMA-48).
	Report ANSI = newAnsi(esc + "6n") // requests the cursor position.
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
