package aec

import (
	"fmt"
)

// RGB3Bit is a 3-bit RGB color.
type RGB3Bit uint8

// RGB8Bit is an 8-bit indexed color.
type RGB8Bit uint8

func newSGR(n uint) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"%dm", n))
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
	return newAnsi(fmt.Sprintf(Esc+"%dm", c+30))
}

// Color3BitB returns an SGR sequence that sets the background color.
func Color3BitB(c RGB3Bit) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"%dm", c+40))
}

// Color8BitF returns an SGR sequence that sets the 8-bit foreground color.
func Color8BitF(c RGB8Bit) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"38;5;%dm", c))
}

// Color8BitB returns an SGR sequence that sets the 8-bit background color.
func Color8BitB(c RGB8Bit) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"48;5;%dm", c))
}

// FullColorF returns an SGR sequence that sets the 24-bit foreground color.
func FullColorF(r, g, b uint8) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"38;2;%d;%d;%dm", r, g, b))
}

// FullColorB returns an SGR sequence that sets the 24-bit background color.
func FullColorB(r, g, b uint8) ANSI {
	return newAnsi(fmt.Sprintf(Esc+"48;2;%d;%d;%dm", r, g, b))
}

// Text styles and decorations, represented as Select Graphic Rendition (SGR)
// parameters as defined by ECMA-48 / ISO 6429.
var (
	// Intensity and text presentation.
	Bold       ANSI = newSGR(1) // enables bold or increased intensity.
	Faint      ANSI = newSGR(2) // enables faint or decreased intensity.
	Italic     ANSI = newSGR(3) // enables italic.
	Underline  ANSI = newSGR(4) // enables underline.
	BlinkSlow  ANSI = newSGR(5) // enables slow blinking.
	BlinkRapid ANSI = newSGR(6) // enables rapid blinking.
	Inverse    ANSI = newSGR(7) // swaps the foreground and background colors.
	Conceal    ANSI = newSGR(8) // conceals the text.
	CrossOut   ANSI = newSGR(9) // crosses out the text.

	// Framing and overlines.
	Frame    ANSI = newSGR(51) // enables framing.
	Encircle ANSI = newSGR(52) // enables encircling.
	Overline ANSI = newSGR(53) // enables overline.
)

// Foreground colors, represented as SGR parameters.
var (
	// DefaultF restores the default foreground color.
	DefaultF ANSI = newSGR(39)

	// Standard foreground colors.
	BlackF   ANSI = newSGR(30) // black.
	RedF     ANSI = newSGR(31) // red.
	GreenF   ANSI = newSGR(32) // green.
	YellowF  ANSI = newSGR(33) // yellow.
	BlueF    ANSI = newSGR(34) // blue.
	MagentaF ANSI = newSGR(35) // magenta.
	CyanF    ANSI = newSGR(36) // cyan.
	WhiteF   ANSI = newSGR(37) // white (gray).

	// Bright foreground colors.
	LightBlackF   ANSI = newSGR(90) // bright black.
	LightRedF     ANSI = newSGR(91) // bright red.
	LightGreenF   ANSI = newSGR(92) // bright green.
	LightYellowF  ANSI = newSGR(93) // bright yellow.
	LightBlueF    ANSI = newSGR(94) // bright blue.
	LightMagentaF ANSI = newSGR(95) // bright magenta.
	LightCyanF    ANSI = newSGR(96) // bright cyan.
	LightWhiteF   ANSI = newSGR(97) // bright white.
)

// Background colors, represented as SGR parameters.
var (
	// DefaultB restores the default background color.
	DefaultB ANSI = newSGR(49)

	// Standard background colors.
	BlackB   ANSI = newSGR(40) // black.
	RedB     ANSI = newSGR(41) // red.
	GreenB   ANSI = newSGR(42) // green.
	YellowB  ANSI = newSGR(43) // yellow.
	BlueB    ANSI = newSGR(44) // blue.
	MagentaB ANSI = newSGR(45) // magenta.
	CyanB    ANSI = newSGR(46) // cyan.
	WhiteB   ANSI = newSGR(47) // white (gray).

	// Bright background colors.
	LightBlackB   ANSI = newSGR(100) // bright black.
	LightRedB     ANSI = newSGR(101) // bright red.
	LightGreenB   ANSI = newSGR(102) // bright green.
	LightYellowB  ANSI = newSGR(103) // bright yellow.
	LightBlueB    ANSI = newSGR(104) // bright blue.
	LightMagentaB ANSI = newSGR(105) // bright magenta.
	LightCyanB    ANSI = newSGR(106) // bright cyan.
	LightWhiteB   ANSI = newSGR(107) // bright white.
)
