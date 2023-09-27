package color

import (
	"regexp"
	"strconv"
	"strings"
)

// SCode Struct for style code.
type SCode int

// Regular Colors
const (
	Black SCode = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Background Colors
const (
	BackgroundBlack SCode = iota + 40
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

// Bright Colors
const (
	BrightBlack SCode = iota + 90
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

// Bright Background Colors
const (
	BackgroundBrightBlack SCode = iota + 100
	BackgroundBrightRed
	BackgroundBrightGreen
	BackgroundBrightYellow
	BackgroundBrightBlue
	BackgroundBrightMagenta
	BackgroundBrightCyan
	BackgroundBrightWhite
)

// Select Graphic Rendition https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
const (
	Reset      = 0
	Bold       = 1
	Italic     = 3
	Underline  = 4
	Reverse    = 7
	CrossedOut = 9
)

// Match value for styles
var styles = map[string]SCode{
	"black":       Black,
	"red":         Red,
	"green":       Green,
	"yellow":      Yellow,
	"blue":        Blue,
	"magenta":     Magenta,
	"cyan":        Cyan,
	"white":       White,
	"bgblack":     BackgroundBlack,
	"bgred":       BackgroundRed,
	"bggreen":     BackgroundGreen,
	"bgyellow":    BackgroundYellow,
	"bgblue":      BackgroundBlue,
	"bgmagenta":   BackgroundMagenta,
	"bgcyan":      BackgroundCyan,
	"bgwhite":     BackgroundWhite,
	"brblack":     BrightBlack,
	"brred":       BrightRed,
	"brgreen":     BrightGreen,
	"bryellow":    BrightYellow,
	"brblue":      BrightBlue,
	"brmagenta":   BrightMagenta,
	"brcyan":      BrightCyan,
	"brwhite":     BrightWhite,
	"bgbrblack":   BackgroundBrightBlack,
	"bgbrred":     BackgroundBrightRed,
	"bgbrgreen":   BackgroundBrightGreen,
	"bgbryellow":  BackgroundBrightYellow,
	"bgbrblue":    BackgroundBrightBlue,
	"bgbrmagenta": BackgroundBrightMagenta,
	"bgbrcyan":    BackgroundBrightCyan,
	"bgbrwhite":   BackgroundBrightWhite,
	"rt":          Reset,
	"bd":          Bold,
	"ic":          Italic,
	"un":          Underline,
	"re":          Reverse,
	"ct":          CrossedOut,
}

// Colorize function parses the style formatter in text and stylizes it with ANSI codes.
func Colorize(text string) string {
	pattern := `%[a-zA-Z]+`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)

	for _, match := range matches {
		formatter := strings.TrimLeft(match, "%")
		code, exists := styles[strings.ToLower(formatter)]

		pattern := match + `(\s{1})?`

		if exists {
			text = regexp.MustCompile(pattern).ReplaceAllString(text, "\x1b["+strconv.Itoa(int(code))+"m")
		}
	}

	return text
}
