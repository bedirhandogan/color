package color

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
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
			output: fmt.Sprintf("\x1b[38;2;%d;%d;%dmHello, World!", colors["red"][0], colors["red"][1], colors["red"][2]),
		}, {
			input:  "%BgRed100 Hello, World!",
			output: fmt.Sprintf("\x1b[48;2;%d;%d;%dmHello, World!", colors["red100"][0], colors["red100"][1], colors["red100"][2]),
		}, {
			input:  "%italic %red Hello, World! %reset",
			output: fmt.Sprintf("\x1b[%dm\x1b[38;2;%d;%d;%dmHello, World!\x1b[%dm", sgrParams["italic"], colors["red"][0], colors["red"][1], colors["red"][2], sgrParams["reset"]),
		},
	}

	for _, s := range scenarios {
		asrt.Equal(Colorize(s.input), s.output)
	}
}

func TestNewColor(t *testing.T) {
	asrt := assert.New(t)

	scenarios := []struct {
		name  string
		value []int
	}{
		{
			name:  "TestRed",
			value: []int{255, 0, 0},
		}, {
			name:  "TestGreen",
			value: []int{10},
		},
	}

	for _, s := range scenarios {
		// create new color
		NewColor(s.name, s.value...)

		_, exist := newColors[strings.ToLower(s.name)]
		asrt.Equal(exist, true)
	}
}

func TestUseNewColor(t *testing.T) {
	asrt := assert.New(t)

	NewColor("TestRed", 9)
	NewColor("BgTestRed", 9)

	scenarios := []struct {
		text, match, output string
	}{
		{
			text:   "%TestRed Hello, World!",
			match:  "%TestRed",
			output: fmt.Sprintf("\x1b[38;5;%dmHello, World!", 9),
		}, {
			text:   "%BgTestRed Hello, World!",
			match:  "%BgTestRed",
			output: fmt.Sprintf("\x1b[48;5;%dmHello, World!", 9),
		},
	}

	for _, s := range scenarios {
		asrt.Equal(useNewColor(s.text, s.match), s.output)
	}
}
