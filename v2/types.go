package color

import "github.com/jwalton/go-supportscolor"

// The 'color' interface defines the necessary method for color types.
type color interface {
	escape(bg bool) string
}

// The 'ansi' struct represents an ANSI color code.
type ansi struct {
	code uint8
}

// escape method for the 'ansi' struct generates the ANSI escape sequence.
func (c *ansi) escape(back bool) string {
	return AnsiToEscape(c.code, back)
}

// The 'rgb' struct represents RGB color components.
type rgb struct {
	r uint8
	g uint8
	b uint8
}

// The escape method takes an `rgb` object representing the given RGB color values
// and is used to produce colored output for the standard error stream (`stderr`) based
// on the supported color features. If stderr supports 16 million colors, it converts
// the RGB color values to an escape sequence; otherwise, it converts the RGB color
// values to ANSI color codes and then into an escape sequence to generate colored output.
// The `back` parameter determines whether the color will be used as a background (true)
// or foreground (false) color.
func (c *rgb) escape(back bool) string {
	if supportscolor.Stderr().Has16m {
		return RgbToEscape(c.r, c.g, c.b, back)
	}

	return AnsiToEscape(RgbToAnsi(c.r, c.g, c.b), back)
}
