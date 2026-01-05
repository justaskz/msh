package dev

import (
	"strings"
	"unicode"
)

func ParseVersion(input string) string {
	var filtered strings.Builder
	for _, r := range input {
		if unicode.IsDigit(r) || r == '.' {
			filtered.WriteRune(r)
		}
	}
	return filtered.String()
}
