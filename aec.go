package aec

import "fmt"

// EraseMode is a input for EraceXXX function.
type EraseMode uint

var (
	// EraseModes is enumeration of EraseMode.
	EraseModes struct {
		// All erase all.
		All EraseMode

		// Head erase to head.
		Head EraseMode

		// Tail erase to tail.
		Tail EraseMode
	}

	// Save saves the cursor position.
	Save ANSI

	// Restore restores the cursor position.
	Restore ANSI
)

// Up moves up the cursor.
func Up(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dA", n))
}

// Down moves down the cursor.
func Down(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dB", n))
}

// Right moves right the cursor.
func Right(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dC", n))
}

// Left moves left the cursor.
func Left(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dD", n))
}

// NextLine moves down the cursor to head of a line.
func NextLine(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dE", n))
}

// PreviousLine moves up the cursor to head of a line.
func PreviousLine(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dF", n))
}

// Column set the cursor position to a given column.
func Column(col uint) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dG", col))
}

// Position set the cursor position to a given absolute position.
func Position(row, col uint) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%d;%dH", row, col))
}

// EraseDisplay erases display by given EraseMode.
func EraseDisplay(m EraseMode) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dJ", m))
}

// EraseLine erases lines by given EraseMode.
func EraseLine(m EraseMode) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dK", m))
}

// ScrollUp scrolls up the page.
func ScrollUp(n int) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dS", n))
}

// ScrollDown scrolls down the page.
func ScrollDown(n int) ANSI {
	if n == 0 {
		return empty
	}
	return newAnsi(fmt.Sprintf(esc+"%dT", n))
}

func init() {
	EraseModes = struct {
		All  EraseMode
		Head EraseMode
		Tail EraseMode
	}{
		Tail: 0,
		Head: 1,
		All:  2,
	}

	Save = newAnsi(esc + "s")
	Restore = newAnsi(esc + "u")
}
