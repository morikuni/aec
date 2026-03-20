package aec

import "strconv"

// RGB3Bit is a 3bit RGB color.
type RGB3Bit uint8

// RGB8Bit is a 8bit RGB color.
type RGB8Bit uint8

func newSGR(n uint8) ANSI {
	return sgr1("", n)
}

func sgr1(prefix string, n uint8) ANSI {
	buf := make([]byte, 0, len(esc)+len(prefix)+3+1)
	buf = append(buf, esc...)
	buf = append(buf, prefix...)
	buf = strconv.AppendUint(buf, uint64(n), 10)
	buf = append(buf, 'm')
	return newAnsi(string(buf))
}

func newFullColor(prefix string, a, b, c uint8) ANSI {
	buf := make([]byte, 0, len(esc)+len(prefix)+3+1+3+1+3+1)
	buf = append(buf, esc...)
	buf = append(buf, prefix...)
	buf = strconv.AppendUint(buf, uint64(a), 10)
	buf = append(buf, ';')
	buf = strconv.AppendUint(buf, uint64(b), 10)
	buf = append(buf, ';')
	buf = strconv.AppendUint(buf, uint64(c), 10)
	buf = append(buf, 'm')
	return newAnsi(string(buf))
}

// NewRGB3Bit create a RGB3Bit from given RGB.
func NewRGB3Bit(r, g, b uint8) RGB3Bit {
	return RGB3Bit((r >> 7) | ((g >> 6) & 0x2) | ((b >> 5) & 0x4))
}

// NewRGB8Bit create a RGB8Bit from given RGB.
func NewRGB8Bit(r, g, b uint8) RGB8Bit {
	return RGB8Bit(16 + 36*(r/43) + 6*(g/43) + b/43)
}

// Color3BitF set the foreground color of text.
func Color3BitF(c RGB3Bit) ANSI {
	return newSGR(uint8(c + 30))
}

// Color3BitB set the background color of text.
func Color3BitB(c RGB3Bit) ANSI {
	return newSGR(uint8(c + 40))
}

// Color8BitF set the foreground color of text.
func Color8BitF(c RGB8Bit) ANSI {
	return sgr1("38;5;", uint8(c))
}

// Color8BitB set the background color of text.
func Color8BitB(c RGB8Bit) ANSI {
	return sgr1("48;5;", uint8(c))
}

// FullColorF set the foreground color of text.
func FullColorF(r, g, b uint8) ANSI {
	return newFullColor("38;2;", r, g, b)
}

// FullColorB set the background color of text.
func FullColorB(r, g, b uint8) ANSI {
	return newFullColor("48;2;", r, g, b)
}

// Style
var (
	// Bold set the text style to bold or increased intensity.
	Bold ANSI = newSGR(1)

	// Faint set the text style to faint.
	Faint ANSI = newSGR(2)

	// Italic set the text style to italic.
	Italic ANSI = newSGR(3)

	// Underline set the text style to underline.
	Underline ANSI = newSGR(4)

	// BlinkSlow set the text style to slow blink.
	BlinkSlow ANSI = newSGR(5)

	// BlinkRapid set the text style to rapid blink.
	BlinkRapid ANSI = newSGR(6)

	// Inverse swap the foreground color and background color.
	Inverse ANSI = newSGR(7)

	// Conceal set the text style to conceal.
	Conceal ANSI = newSGR(8)

	// CrossOut set the text style to crossed out.
	CrossOut ANSI = newSGR(9)

	// Frame set the text style to framed.
	Frame ANSI = newSGR(51)

	// Encircle set the text style to encircled.
	Encircle ANSI = newSGR(52)

	// Overline set the text style to overlined.
	Overline ANSI = newSGR(53)
)

// Foreground color of text.
var (
	// DefaultF is the default foreground color.
	DefaultF ANSI = newSGR(39)

	// Normal color
	BlackF   ANSI = newSGR(30)
	RedF     ANSI = newSGR(31)
	GreenF   ANSI = newSGR(32)
	YellowF  ANSI = newSGR(33)
	BlueF    ANSI = newSGR(34)
	MagentaF ANSI = newSGR(35)
	CyanF    ANSI = newSGR(36)
	WhiteF   ANSI = newSGR(37)

	// Light color
	LightBlackF   ANSI = newSGR(90)
	LightRedF     ANSI = newSGR(91)
	LightGreenF   ANSI = newSGR(92)
	LightYellowF  ANSI = newSGR(93)
	LightBlueF    ANSI = newSGR(94)
	LightMagentaF ANSI = newSGR(95)
	LightCyanF    ANSI = newSGR(96)
	LightWhiteF   ANSI = newSGR(97)
)

// Background color of text.
var (
	// DefaultB is the default background color.
	DefaultB ANSI = newSGR(49)

	// Normal color
	BlackB   ANSI = newSGR(40)
	RedB     ANSI = newSGR(41)
	GreenB   ANSI = newSGR(42)
	YellowB  ANSI = newSGR(43)
	BlueB    ANSI = newSGR(44)
	MagentaB ANSI = newSGR(45)
	CyanB    ANSI = newSGR(46)
	WhiteB   ANSI = newSGR(47)

	// Light color
	LightBlackB   ANSI = newSGR(100)
	LightRedB     ANSI = newSGR(101)
	LightGreenB   ANSI = newSGR(102)
	LightYellowB  ANSI = newSGR(103)
	LightBlueB    ANSI = newSGR(104)
	LightMagentaB ANSI = newSGR(105)
	LightCyanB    ANSI = newSGR(106)
	LightWhiteB   ANSI = newSGR(107)
)
