package color

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestColor(t *testing.T) {
	asrt := assert.New(t)

	scenarios := []struct {
		input  string
		output string
	}{
		{
			input:  "%red Hello, World!",
			output: "\x1b[" + strconv.Itoa(int(Red)) + "mHello, World!",
		},
		{
			input:  "%ic %red Hello, World!",
			output: "\x1b[" + strconv.Itoa(Italic) + "m\x1b[" + strconv.Itoa(int(Red)) + "mHello, World!",
		},
		{
			input:  "%red Hello, %cyan World!",
			output: "\x1b[" + strconv.Itoa(int(Red)) + "mHello, \x1b[" + strconv.Itoa(int(Cyan)) + "mWorld!",
		},
		{
			input:  "%red Hello, %cyan World! %rt \n",
			output: "\x1b[" + strconv.Itoa(int(Red)) + "mHello, \x1b[" + strconv.Itoa(int(Cyan)) + "mWorld! \x1b[0m\n",
		},
	}

	for _, s := range scenarios {
		asrt.Equal(Colorize(s.input), s.output)
	}
}
