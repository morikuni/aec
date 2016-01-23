package aec

import (
	"fmt"
)

func newSGR(n uint) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dm", n))
}

// RGB3Bit is a 3bit RGB color.
type RGB3Bit uint8

// RGB8Bit is a 8bit RGB color.
type RGB8Bit uint8

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
	return newAnsi(fmt.Sprintf(esc+"%dm", c+30))
}

// Color3BitB set the background color of text.
func Color3BitB(c RGB3Bit) ANSI {
	return newAnsi(fmt.Sprintf(esc+"%dm", c+40))
}

// Color8BitF set the foreground color of text.
func Color8BitF(c RGB8Bit) ANSI {
	return newAnsi(fmt.Sprintf(esc+"38;5;%dm", c))
}

// Color8BitB set the background color of text.
func Color8BitB(c RGB8Bit) ANSI {
	return newAnsi(fmt.Sprintf(esc+"48;5;%dm", c))
}

// FullColorF set the foreground color of text.
func FullColorF(r, g, b uint8) ANSI {
	return newAnsi(fmt.Sprintf(esc+"38;2;%d;%d;%dm", r, g, b))
}

// FullColorB set the foreground color of text.
func FullColorB(r, g, b uint8) ANSI {
	return newAnsi(fmt.Sprintf(esc+"48;2;%d;%d;%dm", r, g, b))
}

var (
	// Bold set the text style to bold or increased intensity.
	Bold = newSGR(1)

	// Faint set the text style to faint.
	Faint = newSGR(2)

	// Italic set the text style to italic.
	Italic = newSGR(3)

	// Underline set the text style to underline.
	Underline = newSGR(4)

	// BlinkSlow set the text style to slow blink.
	BlinkSlow = newSGR(5)

	// BlinkRapid set the text style to rapid blink.
	BlinkRapid = newSGR(6)

	// Inverse swap the foreground color and background color.
	Inverse = newSGR(7)

	// Conceal set the text style to conceal.
	Conceal = newSGR(8)

	// CrossOut set the text style to crossed out.
	CrossOut = newSGR(9)
)

// Set the foreground color of text.
var (
	BlackF   = newSGR(30)
	RedF     = newSGR(31)
	GreenF   = newSGR(32)
	YellowF  = newSGR(33)
	BlueF    = newSGR(34)
	MagentaF = newSGR(35)
	CyanF    = newSGR(36)
	WhiteF   = newSGR(37)

	// DefaultF is the default color of foreground.
	DefaultF = newSGR(39)
)

// Set the background color of text.
var (
	BlackB   = newSGR(40)
	RedB     = newSGR(41)
	GreenB   = newSGR(42)
	YellowB  = newSGR(43)
	BlueB    = newSGR(44)
	MagentaB = newSGR(45)
	CyanB    = newSGR(46)
	WhiteB   = newSGR(47)

	// DefaultB is the default color of background.
	DefaultB = newSGR(49)
)

var (
	// Frame set the text style to framed.
	Frame = newSGR(51)

	// Encircle set the text style to encircled.
	Encircle = newSGR(52)

	// Overline set the text style to overlined.
	Overline = newSGR(53)
)

// Set the foreground color of text brightly.
var (
	LightBlackF   = newSGR(90)
	LightRedF     = newSGR(91)
	LightGreenF   = newSGR(92)
	LightYellowF  = newSGR(93)
	LightBlueF    = newSGR(94)
	LightMagentaF = newSGR(95)
	LightCyanF    = newSGR(96)
	LightWhiteF   = newSGR(97)
)

// Set the background color of text brightly.
var (
	LightBlackB   = newSGR(100)
	LightRedB     = newSGR(101)
	LightGreenB   = newSGR(102)
	LightYellowB  = newSGR(103)
	LightBlueB    = newSGR(104)
	LightMagentaB = newSGR(105)
	LightCyanB    = newSGR(106)
	LightWhiteB   = newSGR(107)
)
