package aec

import "strconv"

// EraseMode is listed in a variable EraseModes.
type EraseMode uint

var (
	// EraseModes is a list of EraseMode.
	EraseModes = struct {
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
func Up(n uint) ANSI { return csiUintDefault1(n, 'A') }

// Down moves the cursor down by n positions (CUD).
func Down(n uint) ANSI { return csiUintDefault1(n, 'B') }

// Right moves the cursor right by n positions (CUF).
func Right(n uint) ANSI { return csiUintDefault1(n, 'C') }

// Left moves the cursor left by n positions (CUB).
func Left(n uint) ANSI { return csiUintDefault1(n, 'D') }

// NextLine moves the cursor down n lines and to the beginning of the line (CNL).
func NextLine(n uint) ANSI { return csiUintDefault1(n, 'E') }

// PreviousLine moves the cursor up n lines and to the beginning of the line (CPL).
func PreviousLine(n uint) ANSI { return csiUintDefault1(n, 'F') }

// Column sets the cursor position to a given column (CHA).
func Column(col uint) ANSI {
	if col <= 1 {
		return newAnsi(esc + "G")
	}
	return csiUint(col, 'G')
}

// Position sets the cursor position to a given absolute position (CUP).
func Position(row, col uint) ANSI {
	// According to ECMA-48 (ISO/IEC 6429), cursor positioning (CUP) uses 1-based
	// coordinates with default values of 1;1. Empty or zero parameters are treated
	// as defaults, so 0;0 is effectively interpreted as 1;1 (home position).
	if row == 0 {
		row = 1
	}
	if col == 0 {
		col = 1
	}
	if row == 1 && col == 1 {
		// ESC[0;0H == ESC[1;1H == ESC[H
		return newAnsi(esc + "H")
	}
	buf := make([]byte, 0, len(esc)+42)
	buf = append(buf, esc...)
	buf = strconv.AppendUint(buf, uint64(row), 10)
	buf = append(buf, ';')
	buf = strconv.AppendUint(buf, uint64(col), 10)
	buf = append(buf, 'H')
	return newAnsi(string(buf))
}

// EraseDisplay erases the display using the given EraseMode (ED).
func EraseDisplay(m EraseMode) ANSI {
	if m == 0 {
		return newAnsi(esc + "J")
	}
	return csiUint(uint(m), 'J')
}

// EraseLine erases the line using the given EraseMode (EL).
func EraseLine(m EraseMode) ANSI {
	if m == 0 {
		return newAnsi(esc + "K")
	}
	return csiUint(uint(m), 'K')
}

// ScrollUp scrolls the display up by n lines (SU). n <= 0 is a no-op.
func ScrollUp(n int) ANSI {
	return scroll(n, 'S')
}

// ScrollDown scrolls the display down by n lines (SD). n <= 0 is a no-op.
func ScrollDown(n int) ANSI {
	return scroll(n, 'T')
}

func scroll(n int, direction byte) ANSI {
	switch {
	case n <= 0:
		// Negative counts are not defined for SU/SD; treat them as a no-op.
		return empty
	case n == 1:
		return newAnsi(esc + string(direction))
	default:
		return csiInt(n, direction)
	}
}

func csiInt(n int, suffix byte) ANSI {
	buf := make([]byte, 0, len(esc)+21)
	buf = append(buf, esc...)
	buf = strconv.AppendInt(buf, int64(n), 10)
	buf = append(buf, suffix)
	return newAnsi(string(buf))
}

func csiUintDefault1(n uint, direction byte) ANSI {
	switch n {
	case 0:
		return empty
	case 1:
		return newAnsi(esc + string(direction))
	default:
		return csiUint(n, direction)
	}
}

func csiUint(n uint, suffix byte) ANSI {
	buf := make([]byte, 0, len(esc)+21)
	buf = append(buf, esc...)
	buf = strconv.AppendUint(buf, uint64(n), 10)
	buf = append(buf, suffix)
	return newAnsi(string(buf))
}
