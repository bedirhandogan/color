package color

import (
	"fmt"
	"math"
)

var first15Index = [][3]uint8{
	{0, 0, 0},       // Black
	{255, 0, 0},     // Red
	{0, 255, 0},     // Green
	{255, 255, 0},   // Yellow
	{0, 0, 255},     // Blue
	{255, 0, 255},   // Magenta
	{0, 255, 255},   // Cyan
	{255, 255, 255}, // White
	{128, 128, 128}, // Gray
	{255, 0, 0},     // Bright Red
	{0, 255, 0},     // Bright Green
	{255, 255, 0},   // Bright Yellow
	{0, 0, 255},     // Bright Blue
	{255, 0, 255},   // Bright Magenta
	{0, 255, 255},   // Bright Cyan
	{255, 255, 255}, // Bright White
}

func AnsiToRgb(index uint8) rgb {
	switch {
	case index < 16:
		return rgb{
			r: first15Index[index][0],
			g: first15Index[index][1],
			b: first15Index[index][2],
		}
	case index >= 16 && index <= 231:
		sub := index - 16
		return rgb{
			r: uint8(((sub / 36) % 6) * 255 / 5),
			g: uint8(((sub / 6) % 6) * 255 / 5),
			b: uint8((sub % 6) * 255 / 5),
		}
	case index >= 232:
		return rgb{
			r: uint8((index-232)*10 + 8),
			g: uint8((index-232)*10 + 8),
			b: uint8((index-232)*10 + 8),
		}
	}

	return rgb{}
}

func RgbToAnsi(r, g, b uint8) uint8 {
	ratio := func(x uint8) uint8 {
		return uint8(math.Round(float64(x) / 255.0 * 5.0))
	}

	return uint8(16 + 36*int(ratio(r)) + 6*int(ratio(g)) + int(ratio(b)))
}

func AnsiToEscape(code uint8, back bool) string {
	var format string

	if back {
		format = "\x1b[48;5;%dm"
	} else {
		format = "\x1b[38;5;%dm"
	}

	return fmt.Sprintf(format, code)
}

func RgbToEscape(r, g, b uint8, back bool) string {
	var format string

	if back {
		format = "\x1b[48;2;%d;%d;%dm"
	} else {
		format = "\x1b[38;2;%d;%d;%dm"
	}

	return fmt.Sprintf(format, r, g, b)
}

func SgrToEscape(code uint8) string {
	return fmt.Sprintf("\x1b[%dm", code)
}
