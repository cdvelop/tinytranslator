package tinytranslator

import (
	"bytes"
	"reflect"
	"strconv"
)

type writer interface {
	Write(p []byte) (n int, err error)
}

// defaultWriter implements writer interface
type defaultWriter struct{}

func (d defaultWriter) Write(p []byte) (n int, err error) { return len(p), nil }

// translation represents a single translation entry
type translation struct {
	Key    string   // Original snake case key
	Values []string // Values in different languages
}

// language represents a supported language
type language struct {
	Code  string // eg: "en", "es"
	Index int    // Index in the translations array
}

type Translator struct {
	defaultLang   string
	langSupported []language
	translations  []translation
	err           errMessage
	writer
}

type errMessage struct {
	message string
}

// global dictionary of translations
var D dictionary

// NewTranslationEngine creates and initializes a new translation engine.
//
// It analyzes a global dictionary structure to build translations and supported languages.
// English is always included as the default language, and additional languages are
// extracted from struct tags in the dictionary.
//
// Parameters:
//   - params: Optional variadic parameters that can include:
//   - string: Sets the default language code (e.g., "es", "fr")
//   - writer: A custom writer implementation for outputting translations
//
// Returns:
//   - *Translator: A configured translator instance ready for use
//
// Example usage:
//
//	// Create with default settings (English as default language)
//	translator := NewTranslationEngine()
//
//	// Create with Spanish as default language
//	translator := NewTranslationEngine("es")
//
//	// Create with custom writer
//	customWriter := MyCustomWriter{}
//	translator := NewTranslationEngine(customWriter)
//
//	// Create with both custom language and writer
//	translator := NewTranslationEngine("fr", customWriter)
func NewTranslationEngine(params ...any) *Translator {
	// Define supported languages
	supportedLangs := []language{
		{Code: "en", Index: 0},
	}

	l := Translator{
		defaultLang:   "en",
		langSupported: supportedLangs,
		translations:  make([]translation, 0, 100), // Pre-allocate space
		err:           errMessage{message: ""},
		writer:        defaultWriter{},
	}

	// Process dictionary tags to extract supported languages
	v := reflect.ValueOf(&D).Elem()
	t := v.Type()

	// Parse tags from first field to get language codes
	if t.NumField() > 0 {
		field := t.Field(0)
		langTags := parseTag(field.Tag)

		// Initialize language support
		for langCode := range langTags {
			if langCode != "en" { // "en" already added
				l.langSupported = append(l.langSupported, language{
					Code:  langCode,
					Index: len(l.langSupported),
				})
			}
		}
	}

	// Process dictionary fields and build translations
	for i := range v.NumField() {
		field := v.Field(i)
		dbFieldType := t.Field(i)

		if field.CanSet() {
			// Convert field name to: snake case
			snakeCaseName := snakeCase(dbFieldType.Name)
			// Assign field name to dictionary structure
			field.SetString(snakeCaseName)

			// Create new translation entry
			trans := translation{
				Key:    snakeCaseName,
				Values: make([]string, len(l.langSupported)),
			}

			// Set default translation (English)
			separateName := snakeCase(dbFieldType.Name, " ")
			trans.Values[0] = separateName

			// Add translations for other languages
			for _, lang := range l.langSupported[1:] {
				value := dbFieldType.Tag.Get(lang.Code)
				trans.Values[lang.Index] = value
			}

			l.translations = append(l.translations, trans)
		}
	}

	// Process variadic parameters
	for _, param := range params {
		switch v := param.(type) {
		case string:
			l.setDefaultLanguage(v)
		case writer:
			l.writer = v
		}
	}

	return &l
}

// WithCurrentDeviceLanguage sets the translator to use the system's current language.
//
// It automatically detects the user's preferred language from the operating system
// (in backend) or from the browser (in WebAssembly/frontend) and sets it as the
// default language if it's supported.
//
// If the detected language is not supported, it falls back to English.
//
// Returns:
//   - *Translator: The same translator instance to allow method chaining
//   - error: An error if the language detection failed
//
// Example usage:
//
//	// Create a translator with system language
//	translator := NewTranslationEngine().WithCurrentDeviceLanguage()
func (l *Translator) WithCurrentDeviceLanguage() (*Translator, error) {
	// Get current language from OS or browser
	langCode, err := getCurrentSystemLanguage()
	if err != nil {
		return l, err
	}

	// Try to set the detected language
	err = l.setDefaultLanguage(langCode)
	if err != nil {
		// If not supported, silently fall back to English
		l.defaultLang = "en"
	}

	return l, nil
}

// setDefaultLanguage sets the default language
func (l *Translator) setDefaultLanguage(language string) error {

	langIndex := l.findLanguageIndex(language)
	if langIndex < 0 {
		return l.Err(D.Language, language, D.NotSupported)
	}

	l.defaultLang = language
	return nil
}

// T returns the translation of the given arguments.
func (l Translator) T(args ...any) string {

	var out bytes.Buffer
	var space string

	// Check if we have at least one argument
	if len(args) == 0 {
		return ""
	}

	// Check if first argument is a string and a supported language
	targetLangIndex := l.findLanguageIndex(l.defaultLang)
	if firstArg, ok := args[0].(string); ok {
		// Check if it's a supported language code
		langIndex := l.findLanguageIndex(firstArg)
		if langIndex >= 0 {
			targetLangIndex = langIndex
			args = args[1:] // Remove the language argument
		}
	}

	// Process remaining arguments
	for argNumber, arg := range args {
		switch v := arg.(type) {
		case string:
			if v == "" {
				continue
			}
			out.WriteString(space + l.findTranslation(v, targetLangIndex))
		case []string:
			for _, s := range v {
				if s == "" {
					continue
				}
				out.WriteString(space + l.findTranslation(s, targetLangIndex))
				space = " "
			}
		// Other cases remain the same
		case rune:
			if v == ':' {
				out.WriteString(":")
				continue
			}
			out.WriteString(space + string(v))
		case int:
			out.WriteString(space + strconv.Itoa(v))
		case float64:
			out.WriteString(space + strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			out.WriteString(space + strconv.FormatBool(v))
		case error:
			out.WriteString(space + v.Error())
		default:
			out.WriteString(space + l.findTranslation("argument", targetLangIndex) +
				": " + strconv.Itoa(argNumber) + " " +
				l.findTranslation("unknown", targetLangIndex))
		}
		space = " "
	}
	return out.String()
}

func (l Translator) Err(args ...any) error {
	l.err.message = l.T(args...)
	return l.err
}

func (e errMessage) Error() string {
	return e.message
}

func (l Translator) Print(args ...any) {
	if l.writer != nil {
		l.writer.Write([]byte(l.T(args...)))
	}
}

// findLanguageIndex returns the index of a language or -1 if not found
func (l *Translator) findLanguageIndex(code string) int {
	for _, lang := range l.langSupported {
		if lang.Code == code {
			return lang.Index
		}
	}
	return -1
}

// findTranslation returns the translation for a key in the specified language
func (l *Translator) findTranslation(key string, langIndex int) string {
	for _, trans := range l.translations {
		if trans.Key == key {
			if langIndex >= 0 && langIndex < len(trans.Values) {
				if trans.Values[langIndex] != "" {
					return trans.Values[langIndex]
				}
				// Fallback to default language if translation is empty
				defIndex := l.findLanguageIndex(l.defaultLang)
				return trans.Values[defIndex]
			}
			break
		}
	}
	return key
}
