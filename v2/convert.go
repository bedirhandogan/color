package color

import (
	"fmt"
	"math"
	"strconv"
)

// RGB equivalent of the first 15 ANSI values.
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

// AnsiToRgb function, converts a number specified as an ANSI color index to an RGB color value.
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

// RgbToAnsi converts a color in the RGB color space to an ANSI color code.
func RgbToAnsi(r, g, b uint8) uint8 {
	ratio := func(x uint8) uint8 {
		return uint8(math.Round(float64(x) / 255 * 5))
	}

	if r == b && g == b {
		for i := 0; i <= 240; i += 10 {
			if r <= uint8(i) && r <= 230 {
				return uint8(232 + (i / 10))
			}
		}

		return 255
	}

	return uint8(16 + 36*int(ratio(r)) + 6*int(ratio(g)) + int(ratio(b)))
}

// HexToRgb converts the given hex color code to RGB (Red, Green, Blue) components.
func HexToRgb(hex string) rgb {
	if len(hex[1:]) != 6 {
		return rgb{}
	}

	if hex[0] == '#' {
		hex = hex[1:]
	}

	toR, err := strconv.ParseInt(hex[0:2], 16, 0)
	if err != nil {
		return rgb{}
	}

	toG, err := strconv.ParseInt(hex[2:4], 16, 0)
	if err != nil {
		return rgb{}
	}

	toB, err := strconv.ParseInt(hex[4:6], 16, 0)
	if err != nil {
		return rgb{}
	}

	return rgb{
		r: uint8(toR),
		g: uint8(toG),
		b: uint8(toB),
	}
}

// HexToAnsi converts the given hex color code to its ANSI color code representation.
func HexToAnsi(hex string) uint8 {
	to := HexToRgb(hex)
	return RgbToAnsi(to.r, to.g, to.b)
}

// RgbToHex converts RGB (Red, Green, Blue) color components to a hexadecimal color code.
func RgbToHex(r, g, b uint8) string {
	hexR := strconv.FormatInt(int64(r), 16)
	hexG := strconv.FormatInt(int64(g), 16)
	hexB := strconv.FormatInt(int64(b), 16)

	if len(hexR) == 1 {
		hexR = "0" + hexR
	}

	if len(hexG) == 1 {
		hexG = "0" + hexG
	}

	if len(hexB) == 1 {
		hexB = "0" + hexB
	}

	return "#" + hexR + hexG + hexB
}

// AnsiToHex takes an ANSI color index as input and converts it to a hexadecimal color representation.
func AnsiToHex(index uint8) string {
	to := AnsiToRgb(index)
	return RgbToHex(to.r, to.g, to.b)
}

// AnsiToEscape function, formats a given ANSI color code as a terminal escape sequence.
func AnsiToEscape(code uint8, back bool) string {
	var format string

	if back {
		format = "\x1b[48;5;%dm"
	} else {
		format = "\x1b[38;5;%dm"
	}

	return fmt.Sprintf(format, code)
}

// RgbToEscape function, converts RGB color components to an escape string that can be used in the terminal.
func RgbToEscape(r, g, b uint8, back bool) string {
	var format string

	if back {
		format = "\x1b[48;2;%d;%d;%dm"
	} else {
		format = "\x1b[38;2;%d;%d;%dm"
	}

	return fmt.Sprintf(format, r, g, b)
}

// SgrToEscape function, converts to escape sequence shown in terminal using given SGR (Select Graphics Representation).
func SgrToEscape(code uint8) string {
	return fmt.Sprintf("\x1b[%dm", code)
}
