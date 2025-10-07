package service

import (
	"strings"

	"github.com/evgzor/go1fl-sprint6-final-tpl/pkg/morse"
)

func checkIsMorse(text string) bool {
	for _, rune := range strings.TrimSpace(text) {

		if rune != '-' && rune != '.' && rune != ' ' {
			return false
		}
	}

	return true
}

func Convert(text string) string {

	if text == "" {
		return ""
	}

	if checkIsMorse(text) {
		return morse.ToText(text)
	}

	return morse.ToMorse(text)
}
