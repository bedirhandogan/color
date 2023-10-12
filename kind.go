package color

import (
	"fmt"
)

type Color interface {
	String(bg bool) string
}

type Ansi struct {
	Code uint8
}

func (a *Ansi) String(bg bool) string {
	var format string

	if bg {
		format = "\x1b[48;5;%dm"
	} else {
		format = "\x1b[38;5;%dm"
	}

	return fmt.Sprintf(format, a.Code)
}

type Rgb struct {
	R uint8
	G uint8
	B uint8
}

func (r *Rgb) String(bg bool) string {
	var format string

	if bg {
		format = "\x1b[48;2;%d;%d;%dm"
	} else {
		format = "\x1b[38;2;%d;%d;%dm"
	}

	return fmt.Sprintf(format, r.R, r.G, r.B)
}
