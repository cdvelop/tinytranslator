package lang

import (
	"sync"
	"testing"
)

// Test that the dictionary D has been initialized with snake_case values.
func TestDictionaryInitialization(t *testing.T) {
	// Create a Lang instance to force dictionary initialization.
	_ = New(nil)

	// Test known fields. Assuming dictionary type has these fields.
	// If the dictionary type has more or different fields, adjust expected values as needed.

	// Use reflection here if the fields are not exported.
	// For this test we assume the dictionary fields are accessible.

	if D.Language != "language" {
		t.Errorf("D.Language = %q; want %q", D.Language, "language")
	}
	if D.NotSupported != "not_supported" {
		t.Errorf("D.NotSupported = %q; want %q", D.NotSupported, "not_supported")
	}
}

// Test concurrent use of Lang with translations to English and Spanish.
func TestConcurrentLang(t *testing.T) {
	// Expected translations.
	// For English, since New() uses SnakeCase with " " delimiter for separate names,
	// and for "Language" and "NotSupported" we expect them to be:
	//   "language" and "not supported" (because separate words are joined with space).
	// However, the English mapping is set using the separateName from SnakeCase(dbFieldType.Name, " "),
	// and then the field is set to its snake_case version.
	// So for English:
	const expectedEnLanguage = "language"
	const expectedEnNotSupported = "not supported"

	// For Spanish, translations will be updated only if dictionary struct field tag "es" is provided.
	// We'll assume that the dictionary struct has the following tags:
	//   Language    `es:"idioma"`
	//   NotSupported `es:"no soportado"`
	const expectedEsLanguage = "idioma"
	const expectedEsNotSupported = "no soportado"

	var wg sync.WaitGroup
	iterations := 10
	langInst := New(nil)

	for i := 0; i < iterations; i++ {
		// Test English translation concurrently.
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Set language to "en".
			if err := langInst.SetLanguage("en"); err != nil {
				t.Errorf("SetLanguage(en) error: %v", err)
				return
			}
			// Build a message using dictionary fields and an extra word.
			// Using D.Language and D.NotSupported.
			msg := langInst.T(D.Language, "test", D.NotSupported)
			// Expected: for "en", dictionary mapping is updated with separateName for each field.
			// For D.Language, separateName is obtained by SnakeCase("Language", " ") -> "language"
			// For D.NotSupported, separateName is SnakeCase("NotSupported", " ") -> "not supported"
			expectedMsg := expectedEnLanguage + " test " + expectedEnNotSupported
			if msg != expectedMsg {
				t.Errorf("English translation = %q; want %q", msg, expectedMsg)
			}
		}()

		// Test Spanish translation concurrently.
		wg.Add(1)
		go func() {
			defer wg.Done()
			// langInst := New()
			// Set language to "es".
			if err := langInst.SetLanguage("es"); err != nil {
				t.Errorf("SetLanguage(es) error: %v", err)
				return
			}
			// Build a message using dictionary fields and an extra word.
			got := langInst.T(D.Language, "test", D.NotSupported)
			// Expected: for "es", translation comes from struct tags.
			expectedMsg := expectedEsLanguage + " test " + expectedEsNotSupported
			if got != expectedMsg {
				t.Errorf("\nGot Spanish translation = %q\nwant %q", got, expectedMsg)
			}
		}()
	}
	wg.Wait()
}
