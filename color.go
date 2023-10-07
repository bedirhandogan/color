package color

import (
	"fmt"
	"regexp"
	"strings"
)

var colors = map[string][3]int{
	// Red tones
	"red":    {255, 0, 0}, // Red
	"red10":  {240, 0, 0}, // Light Red
	"red20":  {224, 0, 0}, // Light Red
	"red30":  {208, 0, 0}, // Light Red
	"red40":  {192, 0, 0}, // Light Red
	"red50":  {176, 0, 0}, // Light Red
	"red60":  {160, 0, 0}, // Dark Red
	"red70":  {144, 0, 0}, // Dark Red
	"red80":  {128, 0, 0}, // Dark Red
	"red90":  {112, 0, 0}, // Dark Red
	"red100": {96, 0, 0},  // Dark Red

	// Green tones
	"green":    {0, 255, 0}, // Green
	"green10":  {0, 240, 0}, // Light Green
	"green20":  {0, 224, 0}, // Light Green
	"green30":  {0, 208, 0}, // Light Green
	"green40":  {0, 192, 0}, // Light Green
	"green50":  {0, 176, 0}, // Light Green
	"green60":  {0, 160, 0}, // Dark Green
	"green70":  {0, 144, 0}, // Dark Green
	"green80":  {0, 128, 0}, // Dark Green
	"green90":  {0, 112, 0}, // Dark Green
	"green100": {0, 96, 0},  // Dark Green

	// Yellow tones
	"yellow":    {255, 255, 0}, // Yellow
	"yellow10":  {240, 240, 0}, // Light Yellow
	"yellow20":  {224, 224, 0}, // Light Yellow
	"yellow30":  {208, 208, 0}, // Light Yellow
	"yellow40":  {192, 192, 0}, // Light Yellow
	"yellow50":  {176, 176, 0}, // Light Yellow
	"yellow60":  {160, 160, 0}, // Dark Yellow
	"yellow70":  {144, 144, 0}, // Dark Yellow
	"yellow80":  {128, 128, 0}, // Dark Yellow
	"yellow90":  {112, 112, 0}, // Dark Yellow
	"yellow100": {96, 96, 0},   // Dark Yellow

	// Cyan tones
	"cyan":    {0, 255, 255}, // Cyan
	"cyan10":  {0, 240, 240}, // Light Cyan
	"cyan20":  {0, 224, 224}, // Light Cyan
	"cyan30":  {0, 208, 208}, // Light Cyan
	"cyan40":  {0, 192, 192}, // Light Cyan
	"cyan50":  {0, 176, 176}, // Light Cyan
	"cyan60":  {0, 160, 160}, // Dark Cyan
	"cyan70":  {0, 144, 144}, // Dark Cyan
	"cyan80":  {0, 128, 128}, // Dark Cyan
	"cyan90":  {0, 112, 112}, // Dark Cyan
	"cyan100": {0, 96, 96},   // Dark Cyan

	// Blue tones
	"blue":    {0, 0, 255}, // Blue
	"blue10":  {0, 0, 240}, // Light Blue
	"blue20":  {0, 0, 224}, // Light Blue
	"blue30":  {0, 0, 208}, // Light Blue
	"blue40":  {0, 0, 192}, // Light Blue
	"blue50":  {0, 0, 176}, // Light Blue
	"blue60":  {0, 0, 160}, // Dark Blue
	"blue70":  {0, 0, 144}, // Dark Blue
	"blue80":  {0, 0, 128}, // Dark Blue
	"blue90":  {0, 0, 112}, // Dark Blue
	"blue100": {0, 0, 96},  // Dark Blue

	// Magenta tones
	"magenta":    {255, 0, 255}, // Magenta
	"magenta10":  {240, 0, 240}, // Light Magenta
	"magenta20":  {224, 0, 224}, // Light Magenta
	"magenta30":  {208, 0, 208}, // Light Magenta
	"magenta40":  {192, 0, 192}, // Light Magenta
	"magenta50":  {176, 0, 176}, // Light Magenta
	"magenta60":  {160, 0, 160}, // Dark Magenta
	"magenta70":  {144, 0, 144}, // Dark Magenta
	"magenta80":  {128, 0, 128}, // Dark Magenta
	"magenta90":  {112, 0, 112}, // Dark Magenta
	"magenta100": {96, 0, 96},   // Dark Magenta

	// White tones
	"white":    {255, 255, 255}, // White
	"white10":  {240, 240, 240}, // Light White
	"white20":  {224, 224, 224}, // Light White
	"white30":  {208, 208, 208}, // Light White
	"white40":  {192, 192, 192}, // Light White
	"white50":  {176, 176, 176}, // Light White
	"white60":  {160, 160, 160}, // Dark White
	"white70":  {144, 144, 144}, // Dark White
	"white80":  {128, 128, 128}, // Dark White
	"white90":  {112, 112, 112}, // Dark White
	"white100": {96, 96, 96},    // Dark White

	// Black tones
	"black":    {0, 0, 0},       // Black
	"black10":  {16, 16, 16},    // Dark Gray
	"black20":  {32, 32, 32},    // Dark Gray
	"black30":  {48, 48, 48},    // Dark Gray
	"black40":  {64, 64, 64},    // Gray
	"black50":  {80, 80, 80},    // Gray
	"black60":  {96, 96, 96},    // Light Gray
	"black70":  {112, 112, 112}, // Light Gray
	"black80":  {128, 128, 128}, // Light Gray
	"black90":  {144, 144, 144}, // Light Gray
	"black100": {160, 160, 160}, // Light Gray
}

var newColors = make(map[string]interface{})

// SGR Parameters https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
var sgrParams = map[string]int{
	"reset":     0,
	"bold":      1,
	"italic":    3,
	"blink":     5,
	"underline": 4,
	"overline":  53,
	"invert":    7,
	"strike":    9,
}

// NewColor function retrieves the color name and color code associated with specific color patterns and add new color
//
//	Parameters:
//	- Name: If you use the bg prefix when naming the color, e.g (BgCRed) creates a background color.
//	- Value: If it takes a 1-digit value, for example (10) 256, it will be perceived as ANSI color index.
//	If it takes a 3-digit value, for example (0, 255, 0), it will be perceived as rgb channels.
func NewColor(name string, value ...int) {
	if _, exist := colors[strings.ToLower(name)]; exist {
		fmt.Printf(Colorize("Color: Error occurred color name '%s' already exists."), name)
		return
	}

	switch len(value) {
	case 1:
		newColors[strings.ToLower(name)] = []int{value[0]}
	case 3:
		newColors[strings.ToLower(name)] = []int{value[0], value[1], value[2]}
	}
}

func useNewColors(text, match, pattern string) string {
	value, exist := newColors[strings.ToLower(match[1:])]
	color, isInt := value.([]int)

	if exist && isInt {
		if len(color) == 1 {
			if len(match) >= 3 && strings.ToLower(match[1:3]) == "bg" {
				return regexp.MustCompile(pattern).ReplaceAllString(text, fmt.Sprintf("\x1b[48;5;%dm", color[0]))
			}

			return regexp.MustCompile(pattern).ReplaceAllString(text, fmt.Sprintf("\x1b[38;5;%dm", color[0]))
		}

		if len(match) >= 3 && strings.ToLower(match[1:3]) == "bg" {
			return regexp.MustCompile(pattern).ReplaceAllString(text, fmt.Sprintf("\x1b[48;2;%d;%d;%dm", color[0], color[1], color[2]))
		}

		return regexp.MustCompile(pattern).ReplaceAllString(text, fmt.Sprintf("\x1b[38;2;%d;%d;%dm", color[0], color[1], color[2]))
	}

	return text
}

// Colorize function parses the style formatter in text and stylizes it with ANSI codes and RGB.
func Colorize(text string) string {
	pattern := `%[a-zA-Z0-9]+`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)

	for _, match := range matches {
		pattern := match + `(\s{1})?`

		text = useNewColors(text, match, pattern)

		if sqr, exists := sgrParams[strings.ToLower(match[1:])]; exists {
			text = regexp.MustCompile(`(\s{1})?`+pattern).ReplaceAllString(text, fmt.Sprintf("\x1b[%dm", sqr))
		}

		chans, exists := colors[strings.ToLower(match[3:])]
		if len(match) == 3 && strings.ToLower(match[1:3]) == "bg" && exists {
			text = regexp.MustCompile(pattern).ReplaceAllString(text, fmt.Sprintf("\x1b[48;2;%d;%d;%dm", chans[0], chans[1], chans[2]))
		}

		if chans, exists := colors[strings.ToLower(match[1:])]; exists {
			text = regexp.MustCompile(pattern).ReplaceAllString(text, fmt.Sprintf("\x1b[38;2;%d;%d;%dm", chans[0], chans[1], chans[2]))
		}
	}

	return text
}
