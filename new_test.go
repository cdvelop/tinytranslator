package tinytranslator_test

import (
	"sync"
	"testing"

	. "github.com/cdvelop/tinytranslator"
)

// Constantes para los mensajes esperados
var (
	expectedEnLanguage     = "language"
	expectedEsLanguage     = "idioma"
	expectedEnNotSupported = "not supported"
	expectedEsNotSupported = "no soportado"
)

// Test that the dictionary D has been initialized with snake_case values.
func TestConcurrentAccess(t *testing.T) {
	var wg sync.WaitGroup
	iterations := 100
	translator := NewTranslationEngine()

	type result struct {
		got, expected, lang string
	}
	resultChan := make(chan result, iterations*2)

	for range iterations {
		wg.Add(2)

		go func() {
			defer wg.Done()
			got := translator.T("en", D.Language, "test", D.NotSupported)
			expectedMsg := expectedEnLanguage + " test " + expectedEnNotSupported
			resultChan <- result{got, expectedMsg, "English"}
		}()

		go func() {
			defer wg.Done()
			got := translator.T("es", D.Language, "test", D.NotSupported)
			expectedMsg := expectedEsLanguage + " test " + expectedEsNotSupported
			resultChan <- result{got, expectedMsg, "Spanish"}
		}()
	}

	wg.Wait()
	close(resultChan)

	for r := range resultChan {
		if r.got != r.expected {
			t.Errorf("\n%s translation \ngot: %q\nwant: %q", r.lang, r.got, r.expected)
		}
	}
}

// testWriter implements the writer interface for testing
type testWriter struct {
	write func(p []byte) (n int, err error)
}

func (w testWriter) Write(p []byte) (n int, err error) {
	return w.write(p)
}

func TestConcurrentPrint(t *testing.T) {
	var wg sync.WaitGroup
	iterations := 100

	// Custom writer to capture output
	var mu sync.Mutex
	var outputs []string
	customWriter := testWriter{
		write: func(p []byte) (int, error) {
			mu.Lock()
			defer mu.Unlock()
			outputs = append(outputs, string(p))
			return len(p), nil
		},
	}

	translator := NewTranslationEngine("en", customWriter)

	for range iterations {
		wg.Add(2)

		go func() {
			defer wg.Done()
			translator.Print(D.Language, "test", D.NotSupported)
		}()

		go func() {
			defer wg.Done()
			translator.Print("es", D.Language, "test", D.NotSupported)
		}()
	}

	wg.Wait()

	// Ensure we have the expected number of outputs
	if len(outputs) != iterations*2 {
		t.Errorf("Expected %d outputs, got %d", iterations*2, len(outputs))
	}
}

func TestConcurrentIndividualInstances(t *testing.T) {
	var wg sync.WaitGroup
	iterations := 100

	langInst1 := NewTranslationEngine("en")
	langInst2 := NewTranslationEngine("es")

	for range iterations {
		wg.Add(2)

		go func() {
			defer wg.Done()
			got := langInst1.T(D.Language, "test", D.NotSupported)
			expectedMsg := expectedEnLanguage + " test " + expectedEnNotSupported
			if got != expectedMsg {
				t.Errorf("\nEnglish translation \ngot: %q\nwant: %q", got, expectedMsg)
			}
		}()

		go func() {
			defer wg.Done()
			got := langInst2.T(D.Language, "test", D.NotSupported)
			expectedMsg := expectedEsLanguage + " test " + expectedEsNotSupported
			if got != expectedMsg {
				t.Errorf("\nSpanish translation \ngot: %q\nwant: %q", got, expectedMsg)
			}
		}()
	}

	wg.Wait()
}
