package lang

import (
	"reflect"
	"regexp"
)

// parseTag extrae todas las parejas clave:"valor" de un StructTag.
func parseTag(tag reflect.StructTag) map[string]string {
	result := make(map[string]string)
	// Expresión regular que busca: clave:"valor"
	re := regexp.MustCompile(`(\w+):"([^"]*)"`)
	matches := re.FindAllStringSubmatch(string(tag), -1)
	for _, m := range matches {
		if len(m) == 3 {
			result[m[1]] = m[2]
		}
	}
	return result
}
