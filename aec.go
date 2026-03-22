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

// EraseDisplayMode is the parameter for the ED (Erase in Display, CSI Ps J)
// control sequence. It defines which portion of the display should be erased.
//
// This is an alias of [EraseMode] for backward compatibility.
type EraseDisplayMode = EraseMode

const (
	EraseDisplayTail       EraseDisplayMode = 0 // erase from cursor to end of display
	EraseDisplayHead       EraseDisplayMode = 1 // erase from start of display to cursor
	EraseDisplayAll        EraseDisplayMode = 2 // erase entire display
	EraseDisplaySavedLines EraseDisplayMode = 3 // erase saved lines (scrollback), non-standard but widely supported
)

// EraseLineMode is the parameter for the EL (Erase in Line, CSI Ps K)
// control sequence. It defines which portion of the current line should be erased.
//
// This is an alias of [EraseMode] for backward compatibility.
type EraseLineMode = EraseMode

const (
	EraseLineTail EraseLineMode = 0 // erase from cursor to end of line
	EraseLineHead EraseLineMode = 1 // erase from start of line to cursor
	EraseLineAll  EraseLineMode = 2 // erase entire line
)

// Cursor control and reporting sequences.
//
// Includes CSI sequences defined by ECMA-48 / ISO 6429, DEC private modes,
// and legacy DEC/SCO save and restore sequences.
var (
	// Cursor save and restore sequences.
	//
	// These use both SCO ("ESC[s"/"ESC[u") and DEC ("ESC 7"/"ESC 8")
	// sequences for compatibility. They are not standardized in ECMA-48.
	Save    ANSI = newAnsi(Esc + "s" + "\x1b7") // saves the cursor position.
	Restore ANSI = newAnsi(Esc + "u" + "\x1b8") // restores the cursor position.

	// Cursor visibility (DEC private mode 25).
	Hide ANSI = newAnsi(Esc + "?25l") // hides the cursor.
	Show ANSI = newAnsi(Esc + "?25h") // shows the cursor.

	// Cursor position report (DSR 6, ECMA-48).
	Report ANSI = newAnsi(Esc + "6n") // requests the cursor position.
)

// Up moves the cursor up by n positions (CUU).
func Up(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(Esc+"%dA", n))
}

// Down moves the cursor down by n positions (CUD).
func Down(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(Esc+"%dB", n))
}

// Right moves the cursor right by n positions (CUF).
func Right(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(Esc+"%dC", n))
}

// Left moves the cursor left by n positions (CUB).
func Left(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(Esc+"%dD", n))
}

// NextLine moves the cursor down n lines and to the beginning of the line (CNL).
func NextLine(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(Esc+"%dE", n))
}

// PreviousLine moves the cursor up n lines and to the beginning of the line (CPL).
func PreviousLine(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(Esc+"%dF", n))
}

// Column sets the cursor position to a given column (CHA).
func Column(col uint) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"%dG", col))
}

// Position sets the cursor position to a given absolute position (CUP).
func Position(row, col uint) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"%d;%dH", row, col))
}

// EraseDisplay erases the display using the given [EraseDisplayMode] (ED).
func EraseDisplay(m EraseDisplayMode) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"%dJ", m))
}

// EraseLine erases the line using the given [EraseLineMode] (EL).
func EraseLine(m EraseLineMode) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"%dK", m))
}

// ScrollUp scrolls the display up by n lines (SU).
func ScrollUp(n int) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(Esc+"%dS", n))
}

// ScrollDown scrolls the display down by n lines (SD).
func ScrollDown(n int) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(Esc+"%dT", n))
}
