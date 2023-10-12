package color

import "os"

type color interface {
	escape(bg bool) string
}

type ansi struct {
	code uint8
}

func (c *ansi) escape(back bool) string {
	return AnsiToEscape(c.code, back)
}

type rgb struct {
	r uint8
	g uint8
	b uint8
}

func (c *rgb) escape(back bool) string {
	colorterm := os.Getenv("COLORTERM")

	if colorterm == "truecolor" || colorterm == "24bit" {
		return RgbToEscape(c.r, c.b, c.g, back)
	}

	return AnsiToEscape(RgbToAnsi(c.r, c.b, c.g), back)
}
