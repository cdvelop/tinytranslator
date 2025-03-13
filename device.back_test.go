//go:build !wasm
// +build !wasm

package tinytranslator

import (
	"os"
	"testing"
)

func TestGetCurrentSystemLanguage(t *testing.T) {
	// Guardar variables de entorno originales para restaurarlas después
	origLang := os.Getenv("LANG")
	origLanguage := os.Getenv("LANGUAGE")
	origLcAll := os.Getenv("LC_ALL")
	origLcMessages := os.Getenv("LC_MESSAGES")

	// Restaurar variables de entorno al finalizar
	defer func() {
		os.Setenv("LANG", origLang)
		os.Setenv("LANGUAGE", origLanguage)
		os.Setenv("LC_ALL", origLcAll)
		os.Setenv("LC_MESSAGES", origLcMessages)
	}()

	tests := []struct {
		name        string
		envVars     map[string]string
		expected    string
		shouldError bool
	}{
		{
			name: "LANG variable with simple value",
			envVars: map[string]string{
				"LANG":        "es",
				"LANGUAGE":    "",
				"LC_ALL":      "",
				"LC_MESSAGES": "",
			},
			expected:    "es",
			shouldError: false,
		},
		{
			name: "LANG variable with locale format",
			envVars: map[string]string{
				"LANG":        "fr_FR.UTF-8",
				"LANGUAGE":    "",
				"LC_ALL":      "",
				"LC_MESSAGES": "",
			},
			expected:    "fr",
			shouldError: false,
		},
		{
			name: "LANG variable with dash format",
			envVars: map[string]string{
				"LANG":        "de-DE",
				"LANGUAGE":    "",
				"LC_ALL":      "",
				"LC_MESSAGES": "",
			},
			expected:    "de",
			shouldError: false,
		},
		{
			name: "LANGUAGE variable has precedence over empty LANG",
			envVars: map[string]string{
				"LANG":        "",
				"LANGUAGE":    "pt_BR:pt:en",
				"LC_ALL":      "",
				"LC_MESSAGES": "",
			},
			expected:    "pt",
			shouldError: false,
		},
		{
			name: "LC_ALL variable when others are empty",
			envVars: map[string]string{
				"LANG":        "",
				"LANGUAGE":    "",
				"LC_ALL":      "it_IT.UTF-8",
				"LC_MESSAGES": "",
			},
			expected:    "it",
			shouldError: false,
		},
		{
			name: "LC_MESSAGES variable when others are empty",
			envVars: map[string]string{
				"LANG":        "",
				"LANGUAGE":    "",
				"LC_ALL":      "",
				"LC_MESSAGES": "ru_RU.UTF-8",
			},
			expected:    "ru",
			shouldError: false,
		},
		{
			name: "No environment variables set",
			envVars: map[string]string{
				"LANG":        "",
				"LANGUAGE":    "",
				"LC_ALL":      "",
				"LC_MESSAGES": "",
			},
			expected:    "en", // Default fallback
			shouldError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Configurar las variables de entorno para este caso de prueba
			for key, value := range tc.envVars {
				os.Setenv(key, value)
			}

			// Ejecutar la función a probar
			lang, err := getCurrentSystemLanguage()

			// Verificar si se espera un error
			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Verificar el resultado
			if lang != tc.expected {
				t.Errorf("Expected language code %q but got %q", tc.expected, lang)
			}
		})
	}
}
