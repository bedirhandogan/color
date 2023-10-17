package color

import "os"

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

// The escape method for the 'rgb' struct generates the appropriate escape sequence
// based on the RGB color components. It checks the COLORTERM environment variable
// to determine if the terminal supports true color (24-bit color). If true color
// is supported, it uses the 'RgbToEscape' function directly. Otherwise, it converts
// the RGB values to ANSI color using 'RgbToAnsi' and then generates the escape sequence.
func (c *rgb) escape(back bool) string {
	colorterm := os.Getenv("COLORTERM")

	if colorterm == "truecolor" || colorterm == "24bit" {
		return RgbToEscape(c.r, c.b, c.g, back)
	}

	return AnsiToEscape(RgbToAnsi(c.r, c.b, c.g), back)
}
