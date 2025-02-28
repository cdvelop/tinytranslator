package lang

import (
	"reflect"
	"strconv"
	"strings"
)

type writer interface {
	Write(p []byte) (n int, err error)
}

type Lang struct {
	defaultLang   string   // default "en"
	langSupported []string // eg: "es", "en", "pt", "fr"
	translations  map[string]map[string]string
	out           strings.Builder
	err           errMessage

	writer
}

type errMessage struct {
	message string
}

// dictionary
var D dictionary

func New(w writer) *Lang {

	l := Lang{
		defaultLang:   "en",
		langSupported: []string{},
		translations: map[string]map[string]string{
			"en": {}, // do not complete manually!
		},
		err: errMessage{
			message: "",
		},
	}

	v := reflect.ValueOf(&D).Elem()
	t := v.Type()

	// Parsear las etiquetas de un campo.
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
	// fmt.Println("dictionary initialized", R)
	return &l
}

// Set set the language eg: "es", "en", "pt", "fr"
func (l *Lang) SetLanguage(language string) error {

	if _, ok := l.translations[language]; !ok {
		return l.Err(D.Language, language, D.NotSupported)
	}

	l.defaultLang = language
	return nil
}

// T returns the translation of the given arguments.
// eg: R.T("hello", "world") returns "hello world"
func (l Lang) T(args ...interface{}) string {
	l.out.Reset()
	var space string
	for argNumber, arg := range args {
		switch v := arg.(type) {
		case string:
			if v == "" {
				continue
			}

			if trans, ok := l.translations[l.defaultLang][v]; ok {
				l.out.WriteString(space + trans)
			} else {
				l.out.WriteString(space + v)
			}
		case []string:
			for _, s := range v {
				if s == "" {
					continue
				}
				if trans, ok := l.translations[l.defaultLang][s]; ok {
					l.out.WriteString(space + trans)
				} else {
					l.out.WriteString(space + s)
				}
				space = " "
			}
		case rune:

			if v == ':' {
				l.out.WriteString(":")
				continue
			}

			l.out.WriteString(space + string(v))
		case int:
			l.out.WriteString(space + strconv.Itoa(v))
		case float64:
			l.out.WriteString(space + strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			l.out.WriteString(space + strconv.FormatBool(v))
		case error:
			l.out.WriteString(space + v.Error())
		default:
			l.out.WriteString(space + D.Argument + ": " + strconv.Itoa(argNumber) + " " + D.Unknown)
		}
		space = " "
	}
	return l.out.String()
}

func (l Lang) Err(args ...any) error {
	l.T(args...)
	l.err.message = l.out.String()
	return l.err
}

func (l Lang) Print(args ...any) {
	if l.writer != nil {
		l.writer.Write([]byte(l.T(args...)))
	}
}

func (e errMessage) Error() string {
	return e.message
}
