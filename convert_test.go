package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnsiToRgb(t *testing.T) {
	asrt := assert.New(t)

	scenarios := []struct {
		index    uint8
		expected rgb
	}{
		{
			index:    uint8(5),
			expected: rgb{255, 0, 255},
		},
		{
			index:    uint8(148),
			expected: rgb{50, 50, 0},
		},
		{
			index:    uint8(240),
			expected: rgb{88, 88, 88},
		},
	}

	for _, s := range scenarios {
		got := AnsiToRgb(s.index)
		asrt.Equal(got, s.expected)
	}
}
