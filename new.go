package lang

import (
	"bytes"
	"reflect"
	"strconv"
)

type syncMutex interface {
	Lock()
	Unlock()
}

type writer interface {
	Write(p []byte) (n int, err error)
}

// defaultSync implements syncMutex interface
type defaultSync struct{}

func (d defaultSync) Lock()   {}
func (d defaultSync) Unlock() {}

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

type Lang struct {
	defaultLang   string
	langSupported []language
	translations  []translation
	err           errMessage
	sync          syncMutex
	writer
}

type errMessage struct {
	message string
}

// global dictionary of translations
var D dictionary

func New(params ...any) *Lang {
	// Define supported languages
	supportedLangs := []language{
		{Code: "en", Index: 0},
	}

	l := Lang{
		defaultLang:   "en",
		langSupported: supportedLangs,
		translations:  make([]translation, 0, 100), // Pre-allocate space
		err:           errMessage{message: ""},
		sync:          defaultSync{},
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
		case syncMutex:
			l.sync = v
		case writer:
			l.writer = v
		}
	}

	return &l
}

// SetDefaultLanguage sets the default language
func (l *Lang) SetDefaultLanguage(language string) error {
	l.sync.Lock()
	defer l.sync.Unlock()

	langIndex := l.FindLanguageIndex(language)
	if langIndex < 0 {
		return l.Err(D.Language, language, D.NotSupported)
	}

	l.defaultLang = language
	return nil
}

// T returns the translation of the given arguments.
func (l Lang) T(args ...any) string {
	l.sync.Lock()
	defer l.sync.Unlock()

	var out bytes.Buffer
	var space string

	// Check if we have at least one argument
	if len(args) == 0 {
		return ""
	}

	// Check if first argument is a string and a supported language
	targetLangIndex := l.FindLanguageIndex(l.defaultLang)
	if firstArg, ok := args[0].(string); ok {
		// Check if it's a supported language code
		langIndex := l.FindLanguageIndex(firstArg)
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
			out.WriteString(space + l.FindTranslation(v, targetLangIndex))
		case []string:
			for _, s := range v {
				if s == "" {
					continue
				}
				out.WriteString(space + l.FindTranslation(s, targetLangIndex))
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
			out.WriteString(space + l.FindTranslation("argument", targetLangIndex) +
				": " + strconv.Itoa(argNumber) + " " +
				l.FindTranslation("unknown", targetLangIndex))
		}
		space = " "
	}
	return out.String()
}

func (l Lang) Err(args ...any) error {
	l.err.message = l.T(args...)
	return l.err
}

func (e errMessage) Error() string {
	return e.message
}

func (l Lang) Print(args ...any) {
	if l.writer != nil {
		l.writer.Write([]byte(l.T(args...)))
	}
}

func (l Lang) GetSupportedLanguages() []string {
	var langs []string
	for _, lang := range l.langSupported {
		langs = append(langs, lang.Code)
	}
	return langs
}

// FindLanguageIndex returns the index of a language or -1 if not found
func (l *Lang) FindLanguageIndex(code string) int {
	for _, lang := range l.langSupported {
		if lang.Code == code {
			return lang.Index
		}
	}
	return -1
}

// FindTranslation returns the translation for a key in the specified language
func (l *Lang) FindTranslation(key string, langIndex int) string {
	for _, trans := range l.translations {
		if trans.Key == key {
			if langIndex >= 0 && langIndex < len(trans.Values) {
				if trans.Values[langIndex] != "" {
					return trans.Values[langIndex]
				}
				// Fallback to default language if translation is empty
				defIndex := l.FindLanguageIndex(l.defaultLang)
				return trans.Values[defIndex]
			}
			break
		}
	}
	return key
}
