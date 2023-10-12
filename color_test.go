package color

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorize(t *testing.T) {
	asrt := assert.New(t)

	scenarios := []struct {
		input  string
		output string
	}{
		{
			input:  "%red Hello, World!",
			output: fmt.Sprintf("%vHello, World!", colors["red"].String(false)),
		}, {
			input:  "%BgRed100 Hello, World!",
			output: fmt.Sprintf("%vHello, World!", colors["red100"].String(true)),
		}, {
			input:  "%italic %red Hello, World!%reset",
			output: fmt.Sprintf("\x1b[%dm%vHello, World!\x1b[%dm", sgrParams["italic"], colors["red"].String(false), sgrParams["reset"]),
		},
	}

	for _, s := range scenarios {
		asrt.Equal(String(s.input), s.output)
	}
}

func TestRegister(t *testing.T) {
	asrt := assert.New(t)

	scenarios := []struct {
		name  string
		value Color
	}{
		{
			name:  "TestRed",
			value: &Rgb{255, 0, 0},
		}, {
			name:  "TestGreen",
			value: &Ansi{10},
		},
	}

	for _, s := range scenarios {
		// create new color
		Register(s.name, s.value)

		_, exist := additional[strings.ToLower(s.name)]
		asrt.Equal(exist, true)
	}
}

func TestUseAdditional(t *testing.T) {
	asrt := assert.New(t)

	Register("TestRed", &Ansi{9})
	Register("BgTestRed", &Ansi{9})

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
		asrt.Equal(useAdditional(s.text, s.match), s.output)
	}
}
