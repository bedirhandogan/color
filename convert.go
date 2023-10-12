package color

import (
	"fmt"
	"math"
)

func AnsiToRgb(ansi uint8) (r, g, b uint8) {
	sub := ansi - 16
	r = uint8((sub / 36) * 255 / 5)
	g = uint8(((sub % 36) / 6) * 255 / 5)
	b = uint8((sub % 6) * 255 / 5)

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
