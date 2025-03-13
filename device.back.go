//go:build !wasm
// +build !wasm

package tinytranslator

import (
	"os"
	"strings"
)

// getCurrentSystemLanguage returns the current system language code
func getCurrentSystemLanguage() (string, error) {
	// Common environment variables that contain language information
	langVars := []string{"LANG", "LANGUAGE", "LC_ALL", "LC_MESSAGES"}

	// Try to get language from environment variables
	for _, envVar := range langVars {
		if envValue := os.Getenv(envVar); envValue != "" {
			// Parse language code from environment variable
			// Formats can be: en_US.UTF-8, en-US, en_US, en
			code := strings.Split(envValue, ".")[0] // Remove encoding part
			code = strings.Split(code, "_")[0]      // Get language part, ignore region
			code = strings.Split(code, "-")[0]      // Handle dash format

			if code != "" {
				return strings.ToLower(code), nil
			}
		}
	}

	// Default to English if cannot detect
	return "en", nil
}
