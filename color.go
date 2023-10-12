package color

import (
	"fmt"
	"regexp"
	"strings"
)

var colors = map[string]Color{
	// Red tones
	"red":    &Rgb{255, 0, 0}, // Red
	"red10":  &Rgb{240, 0, 0}, // Light Red
	"red20":  &Rgb{224, 0, 0}, // Light Red
	"red30":  &Rgb{208, 0, 0}, // Light Red
	"red40":  &Rgb{192, 0, 0}, // Light Red
	"red50":  &Rgb{176, 0, 0}, // Light Red
	"red60":  &Rgb{160, 0, 0}, // Dark Red
	"red70":  &Rgb{144, 0, 0}, // Dark Red
	"red80":  &Rgb{128, 0, 0}, // Dark Red
	"red90":  &Rgb{112, 0, 0}, // Dark Red
	"red100": &Rgb{96, 0, 0},  // Dark Red

	// Green tones
	"green":    &Rgb{0, 255, 0}, // Green
	"green10":  &Rgb{0, 240, 0}, // Light Green
	"green20":  &Rgb{0, 224, 0}, // Light Green
	"green30":  &Rgb{0, 208, 0}, // Light Green
	"green40":  &Rgb{0, 192, 0}, // Light Green
	"green50":  &Rgb{0, 176, 0}, // Light Green
	"green60":  &Rgb{0, 160, 0}, // Dark Green
	"green70":  &Rgb{0, 144, 0}, // Dark Green
	"green80":  &Rgb{0, 128, 0}, // Dark Green
	"green90":  &Rgb{0, 112, 0}, // Dark Green
	"green100": &Rgb{0, 96, 0},  // Dark Green

	// Yellow tones
	"yellow":    &Rgb{255, 255, 0}, // Yellow
	"yellow10":  &Rgb{240, 240, 0}, // Light Yellow
	"yellow20":  &Rgb{224, 224, 0}, // Light Yellow
	"yellow30":  &Rgb{208, 208, 0}, // Light Yellow
	"yellow40":  &Rgb{192, 192, 0}, // Light Yellow
	"yellow50":  &Rgb{176, 176, 0}, // Light Yellow
	"yellow60":  &Rgb{160, 160, 0}, // Dark Yellow
	"yellow70":  &Rgb{144, 144, 0}, // Dark Yellow
	"yellow80":  &Rgb{128, 128, 0}, // Dark Yellow
	"yellow90":  &Rgb{112, 112, 0}, // Dark Yellow
	"yellow100": &Rgb{96, 96, 0},   // Dark Yellow

	// Cyan tones
	"cyan":    &Rgb{0, 255, 255}, // Cyan
	"cyan10":  &Rgb{0, 240, 240}, // Light Cyan
	"cyan20":  &Rgb{0, 224, 224}, // Light Cyan
	"cyan30":  &Rgb{0, 208, 208}, // Light Cyan
	"cyan40":  &Rgb{0, 192, 192}, // Light Cyan
	"cyan50":  &Rgb{0, 176, 176}, // Light Cyan
	"cyan60":  &Rgb{0, 160, 160}, // Dark Cyan
	"cyan70":  &Rgb{0, 144, 144}, // Dark Cyan
	"cyan80":  &Rgb{0, 128, 128}, // Dark Cyan
	"cyan90":  &Rgb{0, 112, 112}, // Dark Cyan
	"cyan100": &Rgb{0, 96, 96},   // Dark Cyan

	// Blue tones
	"blue":    &Rgb{0, 0, 255}, // Blue
	"blue10":  &Rgb{0, 0, 240}, // Light Blue
	"blue20":  &Rgb{0, 0, 224}, // Light Blue
	"blue30":  &Rgb{0, 0, 208}, // Light Blue
	"blue40":  &Rgb{0, 0, 192}, // Light Blue
	"blue50":  &Rgb{0, 0, 176}, // Light Blue
	"blue60":  &Rgb{0, 0, 160}, // Dark Blue
	"blue70":  &Rgb{0, 0, 144}, // Dark Blue
	"blue80":  &Rgb{0, 0, 128}, // Dark Blue
	"blue90":  &Rgb{0, 0, 112}, // Dark Blue
	"blue100": &Rgb{0, 0, 96},  // Dark Blue

	// Magenta tones
	"magenta":    &Rgb{255, 0, 255}, // Magenta
	"magenta10":  &Rgb{240, 0, 240}, // Light Magenta
	"magenta20":  &Rgb{224, 0, 224}, // Light Magenta
	"magenta30":  &Rgb{208, 0, 208}, // Light Magenta
	"magenta40":  &Rgb{192, 0, 192}, // Light Magenta
	"magenta50":  &Rgb{176, 0, 176}, // Light Magenta
	"magenta60":  &Rgb{160, 0, 160}, // Dark Magenta
	"magenta70":  &Rgb{144, 0, 144}, // Dark Magenta
	"magenta80":  &Rgb{128, 0, 128}, // Dark Magenta
	"magenta90":  &Rgb{112, 0, 112}, // Dark Magenta
	"magenta100": &Rgb{96, 0, 96},   // Dark Magenta

	// White tones
	"white":    &Rgb{255, 255, 255}, // White
	"white10":  &Rgb{240, 240, 240}, // Light White
	"white20":  &Rgb{224, 224, 224}, // Light White
	"white30":  &Rgb{208, 208, 208}, // Light White
	"white40":  &Rgb{192, 192, 192}, // Light White
	"white50":  &Rgb{176, 176, 176}, // Light White
	"white60":  &Rgb{160, 160, 160}, // Dark White
	"white70":  &Rgb{144, 144, 144}, // Dark White
	"white80":  &Rgb{128, 128, 128}, // Dark White
	"white90":  &Rgb{112, 112, 112}, // Dark White
	"white100": &Rgb{96, 96, 96},    // Dark White

	// Black tones
	"black":    &Rgb{0, 0, 0},       // Black
	"black10":  &Rgb{16, 16, 16},    // Dark Gray
	"black20":  &Rgb{32, 32, 32},    // Dark Gray
	"black30":  &Rgb{48, 48, 48},    // Dark Gray
	"black40":  &Rgb{64, 64, 64},    // Gray
	"black50":  &Rgb{80, 80, 80},    // Gray
	"black60":  &Rgb{96, 96, 96},    // Light Gray
	"black70":  &Rgb{112, 112, 112}, // Light Gray
	"black80":  &Rgb{128, 128, 128}, // Light Gray
	"black90":  &Rgb{144, 144, 144}, // Light Gray
	"black100": &Rgb{160, 160, 160}, // Light Gray
}

var additional = make(map[string]Color)

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

func Register(name string, color Color) {
	if _, exist := colors[strings.ToLower(name)]; exist {
		fmt.Printf("Color: Error occurred color name '%s' already exists.", name)
		return
	}

	additional[strings.ToLower(name)] = color
}

func useAdditional(text, match string) string {
	value, exist := additional[strings.ToLower(match[1:])]
	pattern := match + `(\s{1})?`

	if exist {
		return regexp.MustCompile(pattern).ReplaceAllString(
			text, value.String(strings.ToLower(match[1:3]) == "bg"),
		)
	}

	return text
}

func Stringf(format string, a ...any) string {
	return String(fmt.Sprintf(format, a...))
}

func String(text string) string {
	pattern := `%[a-zA-Z0-9]+`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)

	for _, match := range matches {
		pattern := match + `(\s{1})?`

		text = useAdditional(text, match)

		if sqr, exists := sgrParams[strings.ToLower(match[1:])]; exists {
			text = regexp.MustCompile(pattern).ReplaceAllString(
				text, fmt.Sprintf("\x1b[%dm", sqr),
			)
			continue
		}

		bg := strings.ToLower(match[1:3]) == "bg"

		var color string
		if bg {
			color = strings.ToLower(match[3:])
		} else {
			color = strings.ToLower(match[1:])
		}

		if color, exists := colors[color]; exists {
			text = regexp.MustCompile(pattern).ReplaceAllString(
				text, color.String(bg),
			)
		}
	}

	return text
}
