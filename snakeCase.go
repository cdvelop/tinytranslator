package lang

import (
	"strings"
	"unicode"
)

// snakeCase converts a string to snake_case format with optional separator.
// If no separator is provided, underscore "_" is used as default.
// Example:
//
//	Input: "camelCase" -> Output: "camel_case"
//	Input: "PascalCase", "-" -> Output: "pascal-case"
//	Input: "APIResponse" -> Output: "api_response"
//	Input: "user123Name", "." -> Output: "user123.name"
func SnakeCase(str string, sep ...string) string {
	separator := "_"
	if len(sep) > 0 {
		separator = sep[0]
	}
	var out string
	for i, r := range str {
		if unicode.IsUpper(r) {
			// If it's uppercase and not the first character, add separator
			if i > 0 && (unicode.IsLower(rune(str[i-1])) || unicode.IsDigit(rune(str[i-1]))) {
				out += separator
			}
			// Convert uppercase to lowercase
			out += strings.ToLower(string(r))
		} else {
			// If it's not uppercase, simply add it
			out += string(r)
		}
	}
	return out
}
