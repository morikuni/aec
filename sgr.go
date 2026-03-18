package aec

import (
	"fmt"
)

// RGB3Bit is a 3-bit RGB color.
type RGB3Bit uint8

// RGB8Bit is an 8-bit indexed color.
type RGB8Bit uint8

func newSGR(n uint) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dm", n))
}

// NewRGB3Bit returns a 3-bit color from the given RGB values.
func NewRGB3Bit(r, g, b uint8) RGB3Bit {
	return RGB3Bit((r >> 7) | ((g >> 6) & 0x2) | ((b >> 5) & 0x4))
}

// NewRGB8Bit returns an 8-bit indexed color from the given RGB values.
func NewRGB8Bit(r, g, b uint8) RGB8Bit {
	return RGB8Bit(16 + 36*(r/43) + 6*(g/43) + b/43)
}

// Color3BitF returns an SGR sequence that sets the foreground color.
func Color3BitF(c RGB3Bit) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dm", c+30))
}

// Color3BitB returns an SGR sequence that sets the background color.
func Color3BitB(c RGB3Bit) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dm", c+40))
}

// Color8BitF returns an SGR sequence that sets the 8-bit foreground color.
func Color8BitF(c RGB8Bit) ANSI {
	return newAnsi(fmt.Sprintf(esc+"38;5;%dm", c))
}

// Color8BitB returns an SGR sequence that sets the 8-bit background color.
func Color8BitB(c RGB8Bit) ANSI {
	return newAnsi(fmt.Sprintf(esc+"48;5;%dm", c))
}

// FullColorF returns an SGR sequence that sets the 24-bit foreground color.
func FullColorF(r, g, b uint8) ANSI {
	return newAnsi(fmt.Sprintf(esc+"38;2;%d;%d;%dm", r, g, b))
}

// FullColorB returns an SGR sequence that sets the 24-bit background color.
func FullColorB(r, g, b uint8) ANSI {
	return newAnsi(fmt.Sprintf(esc+"48;2;%d;%d;%dm", r, g, b))
}

// Text styles and decorations, represented as Select Graphic Rendition (SGR)
// parameters as defined by ECMA-48 / ISO 6429.
var (
	// Bold sets the text style to bold or increased intensity.
	Bold ANSI = newSGR(1)

	// Faint sets the text style to faint or decreased intensity.
	Faint ANSI = newSGR(2)

	// Italic sets the text style to italic.
	Italic ANSI = newSGR(3)

	// Underline sets the text style to underlined.
	Underline ANSI = newSGR(4)

	// BlinkSlow sets the text style to slow blinking.
	BlinkSlow ANSI = newSGR(5)

	// BlinkRapid sets the text style to rapid blinking.
	BlinkRapid ANSI = newSGR(6)

	// Inverse swaps the foreground and background colors.
	Inverse ANSI = newSGR(7)

	// Conceal sets the text style to concealed.
	Conceal ANSI = newSGR(8)

	// CrossOut sets the text style to crossed-out.
	CrossOut ANSI = newSGR(9)

	// Frame sets the text style to framed.
	Frame ANSI = newSGR(51)

	// Encircle sets the text style to encircled.
	Encircle ANSI = newSGR(52)

	// Overline sets the text style to overlined.
	Overline ANSI = newSGR(53)
)

// Foreground colors, represented as SGR parameters.
var (
	// DefaultF restores the default foreground color.
	DefaultF ANSI = newSGR(39)

	// Standard foreground colors.
	BlackF   ANSI = newSGR(30)
	RedF     ANSI = newSGR(31)
	GreenF   ANSI = newSGR(32)
	YellowF  ANSI = newSGR(33)
	BlueF    ANSI = newSGR(34)
	MagentaF ANSI = newSGR(35)
	CyanF    ANSI = newSGR(36)
	WhiteF   ANSI = newSGR(37)

	// Bright foreground colors.
	LightBlackF   ANSI = newSGR(90)
	LightRedF     ANSI = newSGR(91)
	LightGreenF   ANSI = newSGR(92)
	LightYellowF  ANSI = newSGR(93)
	LightBlueF    ANSI = newSGR(94)
	LightMagentaF ANSI = newSGR(95)
	LightCyanF    ANSI = newSGR(96)
	LightWhiteF   ANSI = newSGR(97)
)

// Background colors, represented as SGR parameters.
var (
	// DefaultB restores the default background color.
	DefaultB ANSI = newSGR(49)

	// Standard background colors.
	BlackB   ANSI = newSGR(40)
	RedB     ANSI = newSGR(41)
	GreenB   ANSI = newSGR(42)
	YellowB  ANSI = newSGR(43)
	BlueB    ANSI = newSGR(44)
	MagentaB ANSI = newSGR(45)
	CyanB    ANSI = newSGR(46)
	WhiteB   ANSI = newSGR(47)

	// Bright background colors.
	LightBlackB   ANSI = newSGR(100)
	LightRedB     ANSI = newSGR(101)
	LightGreenB   ANSI = newSGR(102)
	LightYellowB  ANSI = newSGR(103)
	LightBlueB    ANSI = newSGR(104)
	LightMagentaB ANSI = newSGR(105)
	LightCyanB    ANSI = newSGR(106)
	LightWhiteB   ANSI = newSGR(107)
)
