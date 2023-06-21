package lib

import (
	"strings"
	"unicode"
)

func SnakeCase(text string) string {
	words := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})
	snake_case_text := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		snake_case_text += "_" + strings.ToLower(words[i])
	}
	return snake_case_text
}
