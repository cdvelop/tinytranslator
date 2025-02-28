package lang_test

import (
	"sync"
	"testing"

	. "github.com/cesarsoliscaro/lang"
)

// Test that the dictionary D has been initialized with snake_case values.
func TestDictionaryInitialization(t *testing.T) {
	// Create a Lang instance to force dictionary initialization.
	_ = New()

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

const expectedEnLanguage = "language"
const expectedEnNotSupported = "not supported"

const expectedEsLanguage = "idioma"
const expectedEsNotSupported = "no soportado"

// Test concurrent use of Lang with translations to English and Spanish.
func TestConcurrentSetDefaultLanguage(t *testing.T) {

	iterations := 100 // Aumentamos el número de iteraciones
	var wg sync.WaitGroup
	langInst1 := New(&sync.Mutex{})
	langInst2 := New(&sync.Mutex{})

	for range iterations {
		// Test English translation concurrently
		wg.Add(2) // Agregamos 2 porque lanzamos 2 goroutines

		go func() {
			defer wg.Done()
			if err := langInst1.SetDefaultLanguage("en"); err != nil {
				t.Errorf("langInst1 SetDefaultLanguage(en) error: %v", err)
				return
			}
			got := langInst1.T(D.Language, "test", D.NotSupported)
			expectedMsg := expectedEnLanguage + " test " + expectedEnNotSupported
			if got != expectedMsg {
				assertTranslation(t, got, expectedMsg, "English")
			}
		}()

		go func() {
			defer wg.Done()
			if err := langInst2.SetDefaultLanguage("es"); err != nil {
				t.Errorf("langInst2 SetDefaultLanguage(es) error: %v", err)
				return
			}
			got := langInst2.T(D.Language, "test", D.NotSupported)
			expectedMsg := expectedEsLanguage + " test " + expectedEsNotSupported
			if got != expectedMsg {
				assertTranslation(t, got, expectedMsg, "Spanish")
			}
		}()
	}

	wg.Wait() // Esperamos a que todas las goroutines terminen
}

func TestConcurrentIndividualSetLanguage(t *testing.T) {
	iterations := 1000
	var wg sync.WaitGroup
	langInst := New(&sync.Mutex{})

	for range iterations {
		wg.Add(2)

		go func() {
			defer wg.Done()
			got := langInst.T("en", D.Language, "test", D.NotSupported)
			expectedMsg := expectedEnLanguage + " test " + expectedEnNotSupported
			assertTranslation(t, got, expectedMsg, "English")
		}()

		go func() {
			defer wg.Done()
			got := langInst.T("es", D.Language, "test", D.NotSupported)
			expectedMsg := expectedEsLanguage + " test " + expectedEsNotSupported
			assertTranslation(t, got, expectedMsg, "Spanish")
		}()
	}

	wg.Wait()
}

// Helper function to create error messages for translation tests
func assertTranslation(t *testing.T, got, expected, lang string) {
	t.Helper()
	if got != expected {
		t.Errorf("\n%s translation \ngot: %q\nwant: %q", lang, got, expected)
	}
}
