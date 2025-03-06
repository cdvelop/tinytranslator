package lang

// snakeCase converts a string to snake_case format with optional separator.
// If no separator is provided, underscore "_" is used as default.
// Example:
//
//	Input: "camelCase" -> Output: "camel_case"
//	Input: "PascalCase", "-" -> Output: "pascal-case"
//	Input: "APIResponse" -> Output: "api_response"
//	Input: "user123Name", "." -> Output: "user123.name"
func snakeCase(str string, sep ...string) string {
	separator := "_"
	if len(sep) > 0 {
		separator = sep[0]
	}

	var out []byte
	for i, r := range str {
		if r >= 'A' && r <= 'Z' { // IsUpper - solo ASCII
			// Si es mayúscula y no es el primer carácter
			if i > 0 {
				prev := rune(str[i-1])
				// Si el carácter anterior es minúscula o dígito
				if (prev >= 'a' && prev <= 'z') || (prev >= '0' && prev <= '9') {
					out = append(out, separator...)
				}
			}
			// Convertir a minúscula (diferencia ASCII entre 'A' y 'a' es 32)
			out = append(out, byte(r+32))
		} else {
			// Si no es mayúscula, simplemente añadirla
			out = append(out, byte(r))
		}
	}

	return string(out)
}
