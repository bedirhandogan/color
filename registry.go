package color

import (
	"fmt"
	"strings"
)

var defaultColors = map[string]color{
	// Red tones
	"red":    &rgb{255, 0, 0}, // Red
	"red10":  &rgb{240, 0, 0}, // Light Red
	"red20":  &rgb{224, 0, 0}, // Light Red
	"red30":  &rgb{208, 0, 0}, // Light Red
	"red40":  &rgb{192, 0, 0}, // Light Red
	"red50":  &rgb{176, 0, 0}, // Light Red
	"red60":  &rgb{160, 0, 0}, // Dark Red
	"red70":  &rgb{144, 0, 0}, // Dark Red
	"red80":  &rgb{128, 0, 0}, // Dark Red
	"red90":  &rgb{112, 0, 0}, // Dark Red
	"red100": &rgb{96, 0, 0},  // Dark Red

	// Green tones
	"green":    &rgb{0, 255, 0}, // Green
	"green10":  &rgb{0, 240, 0}, // Light Green
	"green20":  &rgb{0, 224, 0}, // Light Green
	"green30":  &rgb{0, 208, 0}, // Light Green
	"green40":  &rgb{0, 192, 0}, // Light Green
	"green50":  &rgb{0, 176, 0}, // Light Green
	"green60":  &rgb{0, 160, 0}, // Dark Green
	"green70":  &rgb{0, 144, 0}, // Dark Green
	"green80":  &rgb{0, 128, 0}, // Dark Green
	"green90":  &rgb{0, 112, 0}, // Dark Green
	"green100": &rgb{0, 96, 0},  // Dark Green

	// Yellow tones
	"yellow":    &rgb{255, 255, 0}, // Yellow
	"yellow10":  &rgb{240, 240, 0}, // Light Yellow
	"yellow20":  &rgb{224, 224, 0}, // Light Yellow
	"yellow30":  &rgb{208, 208, 0}, // Light Yellow
	"yellow40":  &rgb{192, 192, 0}, // Light Yellow
	"yellow50":  &rgb{176, 176, 0}, // Light Yellow
	"yellow60":  &rgb{160, 160, 0}, // Dark Yellow
	"yellow70":  &rgb{144, 144, 0}, // Dark Yellow
	"yellow80":  &rgb{128, 128, 0}, // Dark Yellow
	"yellow90":  &rgb{112, 112, 0}, // Dark Yellow
	"yellow100": &rgb{96, 96, 0},   // Dark Yellow

	// Cyan tones
	"cyan":    &rgb{0, 255, 255}, // Cyan
	"cyan10":  &rgb{0, 240, 240}, // Light Cyan
	"cyan20":  &rgb{0, 224, 224}, // Light Cyan
	"cyan30":  &rgb{0, 208, 208}, // Light Cyan
	"cyan40":  &rgb{0, 192, 192}, // Light Cyan
	"cyan50":  &rgb{0, 176, 176}, // Light Cyan
	"cyan60":  &rgb{0, 160, 160}, // Dark Cyan
	"cyan70":  &rgb{0, 144, 144}, // Dark Cyan
	"cyan80":  &rgb{0, 128, 128}, // Dark Cyan
	"cyan90":  &rgb{0, 112, 112}, // Dark Cyan
	"cyan100": &rgb{0, 96, 96},   // Dark Cyan

	// Blue tones
	"blue":    &rgb{0, 0, 255}, // Blue
	"blue10":  &rgb{0, 0, 240}, // Light Blue
	"blue20":  &rgb{0, 0, 224}, // Light Blue
	"blue30":  &rgb{0, 0, 208}, // Light Blue
	"blue40":  &rgb{0, 0, 192}, // Light Blue
	"blue50":  &rgb{0, 0, 176}, // Light Blue
	"blue60":  &rgb{0, 0, 160}, // Dark Blue
	"blue70":  &rgb{0, 0, 144}, // Dark Blue
	"blue80":  &rgb{0, 0, 128}, // Dark Blue
	"blue90":  &rgb{0, 0, 112}, // Dark Blue
	"blue100": &rgb{0, 0, 96},  // Dark Blue

	// Magenta tones
	"magenta":    &rgb{255, 0, 255}, // Magenta
	"magenta10":  &rgb{240, 0, 240}, // Light Magenta
	"magenta20":  &rgb{224, 0, 224}, // Light Magenta
	"magenta30":  &rgb{208, 0, 208}, // Light Magenta
	"magenta40":  &rgb{192, 0, 192}, // Light Magenta
	"magenta50":  &rgb{176, 0, 176}, // Light Magenta
	"magenta60":  &rgb{160, 0, 160}, // Dark Magenta
	"magenta70":  &rgb{144, 0, 144}, // Dark Magenta
	"magenta80":  &rgb{128, 0, 128}, // Dark Magenta
	"magenta90":  &rgb{112, 0, 112}, // Dark Magenta
	"magenta100": &rgb{96, 0, 96},   // Dark Magenta

	// White tones
	"white":    &rgb{255, 255, 255}, // White
	"white10":  &rgb{240, 240, 240}, // Light White
	"white20":  &rgb{224, 224, 224}, // Light White
	"white30":  &rgb{208, 208, 208}, // Light White
	"white40":  &rgb{192, 192, 192}, // Light White
	"white50":  &rgb{176, 176, 176}, // Light White
	"white60":  &rgb{160, 160, 160}, // Dark White
	"white70":  &rgb{144, 144, 144}, // Dark White
	"white80":  &rgb{128, 128, 128}, // Dark White
	"white90":  &rgb{112, 112, 112}, // Dark White
	"white100": &rgb{96, 96, 96},    // Dark White

	// Black tones
	"black":    &rgb{0, 0, 0},       // Black
	"black10":  &rgb{16, 16, 16},    // Dark Gray
	"black20":  &rgb{32, 32, 32},    // Dark Gray
	"black30":  &rgb{48, 48, 48},    // Dark Gray
	"black40":  &rgb{64, 64, 64},    // Gray
	"black50":  &rgb{80, 80, 80},    // Gray
	"black60":  &rgb{96, 96, 96},    // Light Gray
	"black70":  &rgb{112, 112, 112}, // Light Gray
	"black80":  &rgb{128, 128, 128}, // Light Gray
	"black90":  &rgb{144, 144, 144}, // Light Gray
	"black100": &rgb{160, 160, 160}, // Light Gray
}

var additionalColors = make(map[string]color)

var sgrParams = map[string]uint8{
	"reset":     0,
	"bold":      1,
	"italic":    3,
	"blink":     5,
	"underline": 4,
	"overline":  53,
	"invert":    7,
	"strike":    9,
}

func Escape(name string, back bool) (string, bool) {
	value, exists := additionalColors[name]
	if exists {
		return value.escape(back), true
	}

	value, exists = defaultColors[name]
	if exists {
		return value.escape(back), true
	}

	return "", false
}

func RegisterAnsi(name string, code uint8) error {
	return register(name, &ansi{code})
}

func RegisterRgb(name string, r, g, b uint8) error {
	return register(name, &rgb{r, g, b})
}

func register(name string, color color) error {
	if _, exist := defaultColors[strings.ToLower(name)]; exist {
		return fmt.Errorf("color '%s' already exists", name)
	}

	additionalColors[strings.ToLower(name)] = color
	return nil
}
