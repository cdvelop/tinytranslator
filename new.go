package lang

import (
	"bytes"
	"reflect"
	"slices"
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

type Lang struct {
	defaultLang   string
	langSupported []string
	translations  map[string]map[string]string
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
	l := Lang{
		defaultLang:   "en",
		langSupported: []string{"en"},
		translations: map[string]map[string]string{
			"en": {}, // do not complete manually!
		},
		err:    errMessage{message: ""},
		sync:   defaultSync{},
		writer: defaultWriter{},
	}

	v := reflect.ValueOf(&D).Elem()
	t := v.Type()

	// Parser las etiquetas de un campo.
	field := t.Field(0)
	for key := range parseTag(field.Tag) {
		// fmt.Printf("  Etiqueta %s: %s\n", key, value)
		l.langSupported = append(l.langSupported, key)
	}

	// initialize translations map
	for _, lang := range l.langSupported {
		l.translations[lang] = map[string]string{}
	}

	for i := range v.NumField() {
		field := v.Field(i)
		dbFieldType := t.Field(i)

		if field.CanSet() {
			// Convert field name to: snake case
			snakeCaseName := SnakeCase(dbFieldType.Name)
			// Assign field name to dictionary structure
			field.SetString(snakeCaseName)
			// Separate words
			separateName := SnakeCase(dbFieldType.Name, " ")
			// Update translations map to default language "en"
			l.translations[l.defaultLang][snakeCaseName] = separateName

			// update translations map to other languages
			for _, langSupport := range l.langSupported {
				// Get tags for other languages "es","pt"
				transSupported := dbFieldType.Tag.Get(langSupport)
				if transSupported != "" {
					l.translations[langSupport][snakeCaseName] = transSupported
				}
			}
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

	// fmt.Println("dictionary initialized", R)
	return &l
}

// Set set the language eg: "es", "en", "pt", "fr"
func (l *Lang) SetDefaultLanguage(language string) error {
	l.sync.Lock()
	defer l.sync.Unlock()

	if _, ok := l.translations[language]; !ok {
		return l.Err(D.Language, language, D.NotSupported)
	}

	l.defaultLang = language
	return nil
}

// T returns the translation of the given arguments.
// h (handler reference) eg: h.T("hello", "world", ":", 2021) returns "hello world: 2021"
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
	targetLang := l.defaultLang
	if firstArg, ok := args[0].(string); ok {
		// fmt.Println("firs arg", firstArg, "l.langSupported", l.langSupported)
		// Check if it's a supported language
		if slices.Contains(l.langSupported, firstArg) {
			// fmt.Println("contiene?", firstArg)
			targetLang = firstArg
			args = args[1:] // Remove the language argument from args
			// fmt.Println("currents args", args)
		}
	}

	// Process remaining arguments
	for argNumber, arg := range args {
		switch v := arg.(type) {
		case string:
			if v == "" {
				continue
			}

			if trans, ok := l.translations[targetLang][v]; ok {
				out.WriteString(space + trans)
			} else {
				out.WriteString(space + v)
			}
		case []string:
			for _, s := range v {
				if s == "" {
					continue
				}
				if trans, ok := l.translations[targetLang][s]; ok {
					out.WriteString(space + trans)
				} else {
					out.WriteString(space + s)
				}
				space = " "
			}
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
			out.WriteString(space + D.Argument + ": " + strconv.Itoa(argNumber) + " " + D.Unknown)
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
	return l.langSupported
}
