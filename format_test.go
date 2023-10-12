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
			input:  "&red Hello, World!",
			output: fmt.Sprintf("%vHello, World!", defaultColors["red"].escape(false)),
		}, {
			input:  "&BgRed100 Hello, World!",
			output: fmt.Sprintf("%vHello, World!", defaultColors["red100"].escape(true)),
		}, {
			input:  "&italic &red Hello, World!&reset",
			output: fmt.Sprintf("\x1b[%dm%vHello, World!\x1b[%dm", sgrParams["italic"], defaultColors["red"].escape(false), sgrParams["reset"]),
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
		value color
	}{
		{
			name:  "TestRed",
			value: &rgb{255, 0, 0},
		}, {
			name:  "TestGreen",
			value: &ansi{10},
		},
	}

	for _, s := range scenarios {
		// create new color
		register(s.name, s.value)

		_, exist := additionalColors[strings.ToLower(s.name)]
		asrt.Equal(exist, true)
	}
}

func TestUseAdditional(t *testing.T) {
	asrt := assert.New(t)

	register("TestRed", &ansi{9})
	register("BgTestRed", &ansi{9})

	scenarios := []struct {
		text, output string
	}{
		{
			text:   "&TestRed Hello, World!",
			output: fmt.Sprintf("\x1b[38;5;%dmHello, World!", 9),
		}, {
			text:   "&BgTestRed Hello, World!",
			output: fmt.Sprintf("\x1b[48;5;%dmHello, World!", 9),
		},
	}

	for _, s := range scenarios {
		asrt.Equal(String(s.text), s.output)
	}
}
