# Go Translation Library

## Overview
The Go Translation Library provides a simple way to handle multilingual text in Go applications. It supports:
- Concurrent-safe translations
- Multiple languages
- String formatting with various data types
- Custom error messages

## Installation
```bash
go get github.com/cdvelop/lang
```

## Basic Usage

```go
package main

import (
    "fmt"
    "github.com/cdvelop/lang"
)

func main() {
    // Initialize translator
    translator := lang.New()
    
    // Set default language to Spanish
    translator.SetDefaultLanguage("es")
    
    // Get translation
    msg := translator.T(lang.D.Language, "test", lang.D.NotSupported)
    fmt.Println(msg) // Output: idioma test no soportado
}
```

## Core Components

### Lang Type
The `Lang` type is the main interface for translations. It provides methods for:
- Setting default language
- Formatting translated strings
- Handling errors
- Writing output

### Dictionary
The package includes a pre-defined dictionary with common terms in English and Spanish. New terms can be added by extending the `dictionary` struct.

## API Reference

### New()
```go
func New(params ...any) *Lang
```
Creates a new Lang instance. Optional parameters:
- sync.Mutex for concurrent safety
- io.Writer for custom output

### SetDefaultLanguage()
```go
func (l *Lang) SetDefaultLanguage(language string) error
```
Sets the default language for translations. Returns error if language is not supported.

### T()
```go
func (l Lang) T(args ...any) string
```
Returns translated string. Accepts multiple arguments of different types:
- Strings (translated if in dictionary)
- Numbers
- Booleans
- Errors
- Custom types (converted to string)

### Err()
```go
func (l Lang) Err(args ...any) error
```
Returns a translated error message.

### Print()
```go
func (l Lang) Print(args ...any)
```
Writes translated string to configured writer.

## Examples

### Concurrent Usage
```go
var wg sync.WaitGroup
translator := lang.New(&sync.Mutex{})

wg.Add(2)
go func() {
    defer wg.Done()
    msg := translator.T("es", lang.D.Language)
    fmt.Println(msg)
}()

go func() {
    defer wg.Done()
    msg := translator.T("en", lang.D.Language)
    fmt.Println(msg)
}()
wg.Wait()
```

### Custom Writer
```go
translator := lang.New(os.Stdout)
translator.Print("Hello", " ", "World")
```

## Supported Languages
- English (en)
- Spanish (es)

## Contributing
To add new languages or terms:
1. Extend the dictionary struct in dictionary.go
2. Add language tags with translations
3. Update langSupported in New() function
