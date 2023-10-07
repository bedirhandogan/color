package color

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorize(t *testing.T) {
	asrt := assert.New(t)

	scenarios := []struct {
		input  string
		output string
	}{
		{
			input:  "%red Hello, World!",
			output: fmt.Sprintf("\x1b[38;2;%d;%d;%dmHello, World!", colors["red"][0], colors["red"][2], colors["red"][2]),
		}, {
			input:  "%BgRed100 Hello, World!",
			output: fmt.Sprintf("\x1b[48;2;%d;%d;%dmHello, World!", colors["red100"][0], colors["red100"][2], colors["red100"][2]),
		}, {
			input:  "%italic %red Hello, World! %reset",
			output: fmt.Sprintf("\x1b[%dm\x1b[38;2;%d;%d;%dmHello, World!\x1b[%dm", sgrParams["italic"], colors["red"][0], colors["red"][2], colors["red"][2], sgrParams["reset"]),
		},
	}

	for _, s := range scenarios {
		asrt.Equal(Colorize(s.input), s.output)
	}
}
