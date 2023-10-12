package color

import (
	"fmt"
	"math"
)

func AnsiToRgb(ansi uint8) (r, g, b uint8) {
	if ansi < 16 {
		switch ansi {
		case 0: // Black
			r, g, b = 0, 0, 0
		case 1: // Red
			r, g, b = 128, 0, 0
		case 2: // Green
			r, g, b = 0, 128, 0
		case 3: // Yellow
			r, g, b = 128, 128, 0
		case 4: // Blue
			r, g, b = 0, 0, 128
		case 5: // Magenta
			r, g, b = 128, 0, 128
		case 6: // Cyan
			r, g, b = 0, 128, 128
		case 7: // White
			r, g, b = 192, 192, 192
		case 8: // Gray
			r, g, b = 128, 128, 128
		case 9: // Bright Red
			r, g, b = 255, 0, 0
		case 10: // Bright Green
			r, g, b = 0, 255, 0
		case 11: // Bright Yellow
			r, g, b = 255, 255, 0
		case 12: // Bright Blue
			r, g, b = 0, 0, 255
		case 13: // Bright Magenta
			r, g, b = 255, 0, 255
		case 14: // Bright Cyan
			r, g, b = 0, 255, 255
		case 15: // Bright White
			r, g, b = 255, 255, 255
		}
	} else if ansi >= 16 && ansi <= 231 {
		sub := ansi - 16
		r = uint8(((sub / 36) % 6) * 255 / 5)
		g = uint8(((sub / 6) % 6) * 255 / 5)
		b = uint8((sub % 6) * 255 / 5)
	} else if ansi >= 232 && ansi <= 255 {
		val := uint8((ansi-232)*10 + 8)
		r, g, b = val, val, val
	}
	return
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
