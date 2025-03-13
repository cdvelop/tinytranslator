//go:build wasm
// +build wasm

package tinytranslator

import (
	"syscall/js"
)

// getCurrentSystemLanguage returns the current browser language code
func getCurrentSystemLanguage() (string, error) {
	// Access navigator.language or navigator.languages from JavaScript
	navigator := js.Global().Get("navigator")

	// Try to get preferred language
	if !navigator.Get("language").IsUndefined() {
		// Get browser language (e.g. "en-US", "es-ES")
		fullLang := navigator.Get("language").String()
		// Extract just the language code part (e.g. "en", "es")
		if len(fullLang) >= 2 {
			return fullLang[:2], nil
		}
	}

	// Try to get languages array
	if !navigator.Get("languages").IsUndefined() && navigator.Get("languages").Length() > 0 {
		// Get first preferred language
		fullLang := navigator.Get("languages").Index(0).String()
		// Extract just the language code part
		if len(fullLang) >= 2 {
			return fullLang[:2], nil
		}
	}

	// Default to English if we cannot detect
	return "en", nil
}
