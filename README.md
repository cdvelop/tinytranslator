# Go Translation Library

<div align="center">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8.svg" alt="Go Version">
</div>

## Overview

The Go Translation Library provides a simple, powerful way to handle multilingual text in Go applications. It features a built-in dictionary with translations for common terms across multiple languages, making it easy to create applications that support users worldwide.

### Key Features

- 🔄 **Concurrent-safe** translations with mutex support
- 🌐 **Multiple languages** with extensive built-in translations
- 📝 **Dynamic string formatting** with various data types
- ⚠️ **Integrated error handling** with translated messages
- 🧵 **Zero external dependencies**
- 💻 **Compatible with both backend and WebAssembly** applications

## Installation

```bash
go get github.com/cdvelop/lang
```

## Quick Start

```go
package main

import (
    "fmt"
    . "github.com/cdvelop/lang"
)

func main() {
    // Create a new translator instance
    translator := New()
    
    // Set default language (optional, defaults to English)
    translator.SetDefaultLanguage("es")
    
    // Get translated text
    msg := translator.T(D.Hello, D.World)
    fmt.Println(msg) // Output: "hola mundo"
    
    // Combine translations with variables
    username := "John"
    greeting := translator.T(D.Hello, username)
    fmt.Println(greeting) // Output: "hola John"
    
    // Create translated errors
    err := translator.Err(D.Email, D.NotValid)
    fmt.Println(err.Error()) // Output: "correo electrónico no es valido"
}
```

## Core Components

### Lang Type

The `Lang` struct is the central component that handles translations:

- **Thread-safety**: Optional mutex support for concurrent environments
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
translator := New()

// With mutex for thread safety
translator := New(&sync.Mutex{})

// With custom writer
translator := New(os.Stdout)

// With both options
translator := New(&sync.Mutex{}, os.Stdout)
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

### Direct Output

```go
// Write to the configured writer
translator.Print(D.Hello, "User")
```

## Advanced Usage Examples

### Concurrent Usage

```go
package main

import (
    "fmt"
    "sync"
    . "github.com/cdvelop/lang"
)

func main() {
    var wg sync.WaitGroup
    // Create thread-safe translator
    translator := New(&sync.Mutex{})
    
    wg.Add(2)
    // Handle Spanish requests
    go func() {
        defer wg.Done()
        msg := translator.T("es", D.Language, D.NotSupported)
        fmt.Println(msg) // Output: "idioma no soportado"
    }()
    
    // Handle English requests
    go func() {
        defer wg.Done()
        msg := translator.T("en", D.Language, D.NotSupported)
        fmt.Println(msg) // Output: "language not supported"
    }()
    
    wg.Wait()
}
```

### Custom Output Writer

```go
package main

import (
    "os"
    . "github.com/cdvelop/lang"
)

func main() {
    // Create translator that writes to stdout
    translator := New(os.Stdout)
    
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
    . "github.com/cdvelop/lang"
)

type Handler struct {
    Lang *Lang
}

func NewHandler() *Handler {
    return &Handler{
        Lang: New(),
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

## License

MIT
