package tinytranslator

import "reflect"

// parseTag extrae todas las parejas clave:"valor" de un StructTag sin usar regexp.
func parseTag(tag reflect.StructTag) map[string]string {
	result := make(map[string]string)
	tagStr := string(tag)

	// Estado del parser
	var i int
	for i < len(tagStr) {
		// Encontrar inicio de clave (debe ser una letra o guion bajo)
		for i < len(tagStr) && !isValidKeyStart(tagStr[i]) {
			i++
		}
		if i >= len(tagStr) {
			break
		}

		// Leer clave
		keyStart := i
		for i < len(tagStr) && (isAlphaNum(tagStr[i]) || tagStr[i] == '_') {
			i++
		}
		if i >= len(tagStr) {
			break
		}
		key := tagStr[keyStart:i]

		// Buscar los dos puntos
		if i >= len(tagStr) || tagStr[i] != ':' {
			i++
			continue
		}
		i++ // Saltar los dos puntos

		// Buscar comillas de apertura
		if i >= len(tagStr) || tagStr[i] != '"' {
			i++
			continue
		}
		i++ // Saltar comillas de apertura

		// Leer valor (hasta comillas de cierre)
		valueStart := i
		for i < len(tagStr) && tagStr[i] != '"' {
			i++
		}
		if i >= len(tagStr) {
			break
		}

		value := tagStr[valueStart:i]
		result[key] = value
		i++ // Saltar comillas de cierre
	}

	return result
}

// isValidKeyStart comprueba si el carácter puede ser el inicio de una clave
func isValidKeyStart(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// isAlphaNum comprueba si el carácter es alfanumérico
func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
