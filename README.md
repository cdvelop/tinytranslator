# Go Translation Library

<div align="center">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT">
    <img src="https://img.shields.io/badge/Go-1.22+-00ADD8.svg" alt="Go Version">
    <img src="https://img.shields.io/badge/Tests-Passing-green.svg" alt="Tests: Passing">
</div>

## Overview

The Go Translation Library provides a simple, powerful way to handle multilingual text in Go applications. It features a built-in dictionary with translations for common terms across multiple languages, making it easy to create applications that support users worldwide.

### Key Features

- 🔄 **Natively concurrent-safe** translations using slices and structs, eliminating the need for mutexes.
- 🌐 **Multiple languages** with extensive built-in translations
- 📝 **Dynamic string formatting** with various data types
- ⚠️ **Integrated error handling** with translated messages
- 🧵 **Zero external dependencies**
- 💻 **Compatible with both backend and WebAssembly** applications

## Installation

```bash
go get github.com/cdvelop/tinytranslator
```

## Quick Start

```go
package main

import (
    "fmt"
    . "github.com/cdvelop/tinytranslator"
)

func main() {
    // Create a new translator instance with French as default language
    translator := NewTranslationEngine("fr")
    
    // Get translated text in French
    msg := translator.T(D.Hello, D.World)
    fmt.Println(msg) // Output: "bonjour monde"
    
    // Change language to Portuguese
    translator.SetDefaultLanguage("pt")
    
    // Get translated text in Portuguese
    msg = translator.T(D.Hello, D.World)
    fmt.Println(msg) // Output: "olá mundo"
    
    // Combine translations with variables
    username := "John"
    greeting := translator.T(D.Hello, username)
    fmt.Println(greeting) // Output: "olá John"
}
```

## Core Components

### Lang Type

The `Lang` struct is the central component that handles translations:

- **Default language**: Configurable fallback language
- **Translation lookup**: Efficient key-based translation retrieval
- **Custom output**: Configurable writer interface

### Dictionary

The package includes a pre-defined dictionary with common terms in multiple languages:

```go
// Access dictionary entries directly
D.Language    // "language" in English, "idioma" in Spanish, etc.
D.NotSupported // "not supported" in English, "no soportado" in Spanish, etc.
```

## API Reference

### Creating a Translator

```go
// Basic initialization with default settings
translator := NewTranslationEngine()

// With custom writer and default language
translator := NewTranslationEngine(os.Stdout, "es")

```

### Setting Default Language

```go
// Set default language to Spanish
err := translator.SetDefaultLanguage("es")
if err != nil {
    // Handle error - language not supported
}
```

### Translating Text

```go
// Basic translation
text := translator.T(D.Hello)

// Specify language for this translation only
text := translator.T("fr", D.Hello) // French

// Combine multiple terms
text := translator.T(D.Email, D.NotValid)

// Mix translations with custom text
text := translator.T(D.Language, "Go", D.NotSupported)

// Format with numbers
text := translator.T(D.Value, 42)

// Format with different types
text := translator.T(D.Language, ":", true, 123, 45.67)
```

### Creating Error Messages

```go
// Create error with translated message
err := translator.Err(D.Field, D.Name, D.Required)

// Use the error
if fieldIsEmpty {
    return translator.Err(D.Field, fieldName, D.Empty)
}
```

### Custom Output Writer

```go
package main

import (
    "os"
    . "github.com/cdvelop/tinytranslator"
)

func main() {
    // Create translator that writes to stdout
    translator := NewTranslationEngine(os.Stdout)
    
    // Directly print translated messages
    translator.Print(D.Hello, "User") // Writes to stdout: "Hello User"
}
```

### In Web Handlers

```go
package main

import (
    "fmt"
    "net/http"
    . "github.com/cdvelop/tinytranslator"
)

type Handler struct {
    Lang *Lang
}

func NewHandler() *Handler {
    return &Handler{
        Lang: NewTranslationEngine(),
    }
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Get language from request (query, header, cookie, etc.)
    lang := r.URL.Query().Get("lang")
    
    // Respond in requested language
    msg := h.Lang.T(lang, D.Hello, "User")
    fmt.Fprintf(w, msg)
}
```

## Supported Languages

The library currently supports the following languages:

- 🇬🇧 English (en)
- 🇪🇸 Spanish (es)
- 🇵🇹 Portuguese (pt)
- 🇫🇷 French (fr)
- 🇷🇺 Russian (ru)
- 🇩🇪 German (de)
- 🇮🇹 Italian (it)
- 🇮🇳 Hindi (hi)
- 🇧🇩 Bengali (bn)
- 🇮🇩 Indonesian (id)
- 🇸🇦 Arabic (ar)
- 🇵🇰 Urdu (ur)
- 🇨🇳 Chinese (zh)

## Contributing

To add new languages, terms, or improvements:

1. **Add new dictionary entries**: Extend the `dictionary` struct in `dictionary.go`
2. **Add translations**: Include tags for each supported language
3. **Submit a PR**: Follow the existing code style and add appropriate tests


## Contribute

If you find this project useful and would like to support it, you can make a donation [here with PayPal](https://paypal.me/cdvelop?country.x=CL&locale.x=es_XC).

Any contribution, no matter how small, is greatly appreciated. 🙌


## License

MIT
