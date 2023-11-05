package color

import (
	"regexp"
	"strings"
)

// String function recognizes certain patterns in the entered text and formats the text according to these patterns.
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
		back := len(matchLower) >= 3 && matchLower[1:3] == "bg"

		if back {
			target = matchLower[3:]
		} else {
			target = matchLower[1:]
		}

		escape, exists := escape(target, back)
		if exists {
			replace(match, escape)
		}
	}

	return text
}
