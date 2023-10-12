package color

import (
	"fmt"
	"regexp"
	"strings"
)

func Stringf(format string, a ...interface{}) string {
	return String(fmt.Sprintf(format, a...))
}

func String(text string) string {
	pattern := `&[a-zA-Z0-9]+`
	regx := regexp.MustCompile(pattern)

	matches := regx.FindAllString(text, -1)

	replace := func(match string, new string) {
		pattern := match + `(\s{1})?`
		text = regexp.MustCompile(pattern).ReplaceAllString(text, new)
	}

	for _, match := range matches {

		sgr, exists := sgrParams[strings.ToLower(match[1:])]
		if exists {
			replace(match, SgrToEscape(sgr))
			continue
		}

		matchLower := strings.ToLower(match)

		var target string
		back := matchLower[1:3] == "bg"

		if back {
			target = matchLower[3:]
		} else {
			target = matchLower[1:]
		}

		value, exists := additionalColors[target]
		if exists {
			replace(match, value.escape(back))
			continue
		}

		value, exists = defaultColors[target]
		if exists {
			replace(match, value.escape(back))
		}
	}

	return text
}
