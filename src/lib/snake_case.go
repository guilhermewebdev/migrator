package lib

import (
	"strings"
	"unicode"
)

func SnakeCase(text string) string {
	var result strings.Builder
	for i, char := range text {
		if unicode.IsSpace(char) {
			result.WriteRune('_')
		} else if unicode.IsUpper(char) {
			if i > 0 && unicode.IsLower(rune(text[i-1])) {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(char))
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
