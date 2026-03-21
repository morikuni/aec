package aec

// Builder incrementally composes ANSI escape sequences.
//
// Builder is immutable: methods do not modify the receiver, but return a new
// Builder with the requested sequence appended. This allows chaining and
// branching from any intermediate state:
//
//	b1 := NewBuilder().RedF()
//	b2 := b1.Bold() // b1 is unchanged
//
// The resulting ANSI sequence can be obtained via b.ANSI or b.String().
type Builder struct {
	ANSI ANSI
}

// EmptyBuilder is a Builder with no sequences.
var EmptyBuilder = &Builder{empty}

// NewBuilder returns a new Builder containing the given sequences.
func NewBuilder(a ...ANSI) *Builder {
	return &Builder{concat(a)}
}

// With returns a new Builder with the given sequences appended.
func (builder *Builder) With(a ...ANSI) *Builder {
	return NewBuilder(builder.ANSI.With(a...))
}

// Cursor movement sequences.
//
// These methods reposition the cursor without modifying text.

// Up moves the cursor up by n lines.
func (builder *Builder) Up(n uint) *Builder {
	return builder.With(Up(n))
}

// Down moves the cursor down by n lines.
func (builder *Builder) Down(n uint) *Builder {
	return builder.With(Down(n))
}

// Right moves the cursor right by n columns.
func (builder *Builder) Right(n uint) *Builder {
	return builder.With(Right(n))
}

// Left moves the cursor left by n columns.
func (builder *Builder) Left(n uint) *Builder {
	return builder.With(Left(n))
}

// NextLine moves the cursor down n lines and to column 1.
func (builder *Builder) NextLine(n uint) *Builder {
	return builder.With(NextLine(n))
}

// PreviousLine moves the cursor up n lines and to column 1.
func (builder *Builder) PreviousLine(n uint) *Builder {
	return builder.With(PreviousLine(n))
}

// Column moves the cursor to the given column.
func (builder *Builder) Column(col uint) *Builder {
	return builder.With(Column(col))
}

// Position moves the cursor to the given row and column.
func (builder *Builder) Position(row, col uint) *Builder {
	return builder.With(Position(row, col))
}

// Erasing and scrolling.

// EraseDisplay erases the display using the given mode.
func (builder *Builder) EraseDisplay(m EraseMode) *Builder {
	return builder.With(EraseDisplay(m))
}

// EraseLine erases the current line using the given mode.
func (builder *Builder) EraseLine(m EraseMode) *Builder {
	return builder.With(EraseLine(m))
}

// ScrollUp scrolls the display up by n lines.
func (builder *Builder) ScrollUp(n int) *Builder {
	return builder.With(ScrollUp(n))
}

// ScrollDown scrolls the display down by n lines.
func (builder *Builder) ScrollDown(n int) *Builder {
	return builder.With(ScrollDown(n))
}

// Cursor control and reporting sequences.
//
// These methods control cursor state or request cursor information,
// including CSI sequences (ECMA-48), DEC private modes, and legacy
// DEC/SCO save and restore sequences.

// Save saves the cursor position.
func (builder *Builder) Save() *Builder {
	return builder.With(Save)
}

// Restore restores the cursor position.
func (builder *Builder) Restore() *Builder {
	return builder.With(Restore)
}

// Hide hides the cursor.
func (builder *Builder) Hide() *Builder {
	return builder.With(Hide)
}

// Show shows the cursor.
func (builder *Builder) Show() *Builder {
	return builder.With(Show)
}

// Report requests the cursor position.
func (builder *Builder) Report() *Builder {
	return builder.With(Report)
}

// Text styles and decorations.
//
// These methods modify text presentation (e.g., intensity, underline)
// using SGR parameters defined by ECMA-48 / ISO 6429.

// Bold enables bold or increased intensity.
func (builder *Builder) Bold() *Builder {
	return builder.With(Bold)
}

// Faint enables faint or decreased intensity.
func (builder *Builder) Faint() *Builder {
	return builder.With(Faint)
}

// Italic enables italic.
func (builder *Builder) Italic() *Builder {
	return builder.With(Italic)
}

// Underline enables underline.
func (builder *Builder) Underline() *Builder {
	return builder.With(Underline)
}

// BlinkSlow enables slow blinking.
func (builder *Builder) BlinkSlow() *Builder {
	return builder.With(BlinkSlow)
}

// BlinkRapid enables rapid blinking.
func (builder *Builder) BlinkRapid() *Builder {
	return builder.With(BlinkRapid)
}

// Inverse swaps foreground and background colors.
func (builder *Builder) Inverse() *Builder {
	return builder.With(Inverse)
}

// Conceal conceals the text.
func (builder *Builder) Conceal() *Builder {
	return builder.With(Conceal)
}

// CrossOut crosses out the text.
func (builder *Builder) CrossOut() *Builder {
	return builder.With(CrossOut)
}

// Foreground colors.
//
// These methods set the text (foreground) color using SGR parameters.
// Color rendering depends on the terminal's palette.

// BlackF sets the foreground color to black.
func (builder *Builder) BlackF() *Builder {
	return builder.With(BlackF)
}

// RedF sets the foreground color to red.
func (builder *Builder) RedF() *Builder {
	return builder.With(RedF)
}

// GreenF sets the foreground color to green.
func (builder *Builder) GreenF() *Builder {
	return builder.With(GreenF)
}

// YellowF sets the foreground color to yellow.
func (builder *Builder) YellowF() *Builder {
	return builder.With(YellowF)
}

// BlueF sets the foreground color to blue.
func (builder *Builder) BlueF() *Builder {
	return builder.With(BlueF)
}

// MagentaF sets the foreground color to magenta.
func (builder *Builder) MagentaF() *Builder {
	return builder.With(MagentaF)
}

// CyanF sets the foreground color to cyan.
func (builder *Builder) CyanF() *Builder {
	return builder.With(CyanF)
}

// WhiteF sets the foreground color to white.
func (builder *Builder) WhiteF() *Builder {
	return builder.With(WhiteF)
}

// DefaultF restores the default foreground color.
func (builder *Builder) DefaultF() *Builder {
	return builder.With(DefaultF)
}

// Background colors.
//
// These methods set the background color using SGR parameters.
// Color rendering depends on the terminal's palette.

// BlackB sets the background color to black.
func (builder *Builder) BlackB() *Builder {
	return builder.With(BlackB)
}

// RedB sets the background color to red.
func (builder *Builder) RedB() *Builder {
	return builder.With(RedB)
}

// GreenB sets the background color to green.
func (builder *Builder) GreenB() *Builder {
	return builder.With(GreenB)
}

// YellowB sets the background color to yellow.
func (builder *Builder) YellowB() *Builder {
	return builder.With(YellowB)
}

// BlueB sets the background color to blue.
func (builder *Builder) BlueB() *Builder {
	return builder.With(BlueB)
}

// MagentaB sets the background color to magenta.
func (builder *Builder) MagentaB() *Builder {
	return builder.With(MagentaB)
}

// CyanB sets the background color to cyan.
func (builder *Builder) CyanB() *Builder {
	return builder.With(CyanB)
}

// WhiteB sets the background color to white.
func (builder *Builder) WhiteB() *Builder {
	return builder.With(WhiteB)
}

// DefaultB restores the default background color.
func (builder *Builder) DefaultB() *Builder {
	return builder.With(DefaultB)
}

// Framing and overlines.
//
// These methods add framing or overline text decorations (SGR).

// Frame enables framing.
func (builder *Builder) Frame() *Builder {
	return builder.With(Frame)
}

// Encircle enables encircling.
func (builder *Builder) Encircle() *Builder {
	return builder.With(Encircle)
}

// Overline enables overline.
func (builder *Builder) Overline() *Builder {
	return builder.With(Overline)
}

// Bright foreground colors.

// LightBlackF sets the foreground color to bright black.
func (builder *Builder) LightBlackF() *Builder {
	return builder.With(LightBlackF)
}

// LightRedF sets the foreground color to bright red.
func (builder *Builder) LightRedF() *Builder {
	return builder.With(LightRedF)
}

// LightGreenF sets the foreground color to bright green.
func (builder *Builder) LightGreenF() *Builder {
	return builder.With(LightGreenF)
}

// LightYellowF sets the foreground color to bright yellow.
func (builder *Builder) LightYellowF() *Builder {
	return builder.With(LightYellowF)
}

// LightBlueF sets the foreground color to bright blue.
func (builder *Builder) LightBlueF() *Builder {
	return builder.With(LightBlueF)
}

// LightMagentaF sets the foreground color to bright magenta.
func (builder *Builder) LightMagentaF() *Builder {
	return builder.With(LightMagentaF)
}

// LightCyanF sets the foreground color to bright cyan.
func (builder *Builder) LightCyanF() *Builder {
	return builder.With(LightCyanF)
}

// LightWhiteF sets the foreground color to bright white.
func (builder *Builder) LightWhiteF() *Builder {
	return builder.With(LightWhiteF)
}

// Bright background colors.

// LightBlackB sets the background color to bright black.
func (builder *Builder) LightBlackB() *Builder {
	return builder.With(LightBlackB)
}

// LightRedB sets the background color to bright red.
func (builder *Builder) LightRedB() *Builder {
	return builder.With(LightRedB)
}

// LightGreenB sets the background color to bright green.
func (builder *Builder) LightGreenB() *Builder {
	return builder.With(LightGreenB)
}

// LightYellowB sets the background color to bright yellow.
func (builder *Builder) LightYellowB() *Builder {
	return builder.With(LightYellowB)
}

// LightBlueB sets the background color to bright blue.
func (builder *Builder) LightBlueB() *Builder {
	return builder.With(LightBlueB)
}

// LightMagentaB sets the background color to bright magenta.
func (builder *Builder) LightMagentaB() *Builder {
	return builder.With(LightMagentaB)
}

// LightCyanB sets the background color to bright cyan.
func (builder *Builder) LightCyanB() *Builder {
	return builder.With(LightCyanB)
}

// LightWhiteB sets the background color to bright white.
func (builder *Builder) LightWhiteB() *Builder {
	return builder.With(LightWhiteB)
}

// Extended colors.

// Color3BitF sets a 3-bit foreground color.
func (builder *Builder) Color3BitF(c RGB3Bit) *Builder {
	return builder.With(Color3BitF(c))
}

// Color3BitB sets a 3-bit background color.
func (builder *Builder) Color3BitB(c RGB3Bit) *Builder {
	return builder.With(Color3BitB(c))
}

// Color8BitF sets an 8-bit foreground color.
func (builder *Builder) Color8BitF(c RGB8Bit) *Builder {
	return builder.With(Color8BitF(c))
}

// Color8BitB sets an 8-bit background color.
func (builder *Builder) Color8BitB(c RGB8Bit) *Builder {
	return builder.With(Color8BitB(c))
}

// FullColorF sets a 24-bit foreground color.
func (builder *Builder) FullColorF(r, g, b uint8) *Builder {
	return builder.With(FullColorF(r, g, b))
}

// FullColorB sets a 24-bit background color.
func (builder *Builder) FullColorB(r, g, b uint8) *Builder {
	return builder.With(FullColorB(r, g, b))
}

// RGB3BitF sets a 3-bit foreground color from RGB components.
func (builder *Builder) RGB3BitF(r, g, b uint8) *Builder {
	return builder.Color3BitF(NewRGB3Bit(r, g, b))
}

// RGB3BitB sets a 3-bit background color from RGB components.
func (builder *Builder) RGB3BitB(r, g, b uint8) *Builder {
	return builder.Color3BitB(NewRGB3Bit(r, g, b))
}

// RGB8BitF sets an 8-bit foreground color from RGB components.
func (builder *Builder) RGB8BitF(r, g, b uint8) *Builder {
	return builder.Color8BitF(NewRGB8Bit(r, g, b))
}

// RGB8BitB sets an 8-bit background color from RGB components.
func (builder *Builder) RGB8BitB(r, g, b uint8) *Builder {
	return builder.Color8BitB(NewRGB8Bit(r, g, b))
}
