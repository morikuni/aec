# aec
Go wrapper for ANSI escape code.

## Features

ANSI escape codes depend on terminal environment.
Some these features may not work.

Check [Wikipedia](https://en.wikipedia.org/wiki/ANSI_escape_code) for more detail.

### Moving Cursor

- `Up(n)`
- `Down(n)`
- `Right(n)`
- `Left(n)`
- `NextLine(n)`
- `PreviousLine(n)`
- `Column(col)`
- `Position(row, col)`
- `Save`
- `Restore`

### Erase

- `EraseDisplay(mode)`
- `EraseLine(mode)`

### Scroll

- `ScrollUp(n)`
- `ScrollDown(n)`

### Style

- `Bold`
- `Faint`
- `Italic`
- `Underline`
- `BlinkSlow`
- `BlinkRapid`
- `Inverse`
- `Conceal`
- `CrossOut`
- `Frame`
- `Encircle`
- `Overline`

### Color

Foreground color.

- `DefaultF`
- `BlackF`
- `RedF`
- `GreenF`
- `YellowF`
- `BlueF`
- `MagentaF`
- `CyanF`
- `WhiteF`
- `LightBlackF`
- `LightRedF`
- `LightGreenF`
- `LightYellowF`
- `LightBlueF`
- `LightMagentaF`
- `LightCyanF`
- `LightWhiteF`
- `Color3BitF(color)`
- `Color8BitF(color)`
- `FullColorF(r, g, b)`

Background color.

- `DefaultB`
- `BlackB`
- `RedB`
- `GreenB`
- `YellowB`
- `BlueB`
- `MagentaB`
- `CyanB`
- `WhiteB`
- `LightBlackB`
- `LightRedB`
- `LightGreenB`
- `LightYellowB`
- `LightBlueB`
- `LightMagentaB`
- `LightCyanB`
- `LightWhiteB`
- `Color3BitB(color)`
- `Color8BitB(color)`
- `FullColorB(r, g, b)`

### Color Converter

24bit RGB color to ANSI color.

- `NewRGB3Bit(r, g, b)`
- `NewRGB8Bit(r, g, b)`

### Builder

To mix these features.

```go
custom := aec.EmptyBuilder.Right(2).RGB8BitF(128, 255, 64).RedB().ANSI
custom.Apply("Hello World")
```

## Example

Simple progressbar.

![sample](./sample.gif)

```go
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/morikuni/aec"
)

func main() {
	const n = 20
	builder := aec.EmptyBuilder

	up2 := aec.PreviousLine(2)
	col := aec.Column(n + 1)
	bar := aec.Color8BitF(aec.NewRGB8Bit(64, 255, 64))
	label := builder.LightRedF().Underline().With(col).Right(1).ANSI

	// for up2
	fmt.Println()
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Print(up2)
		fmt.Println(label.Apply(fmt.Sprint(i+1, "/", n)))
		fmt.Print("[")
		fmt.Print(bar.Apply(strings.Repeat("=", i)))
		fmt.Println(col.Apply("]"))
		time.Sleep(100 * time.Millisecond)
	}
}
```

## License

[MIT](./LICENSE)


