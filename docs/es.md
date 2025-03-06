# Biblioteca de Traducción en Go

## Visión General
La Biblioteca de Traducción en Go proporciona una forma sencilla de manejar texto multilingüe en aplicaciones Go. Soporta:
- Traducciones seguras para uso concurrente
- Múltiples idiomas
- Formateo de cadenas con varios tipos de datos
- Mensajes de error personalizados

## para que esta pensado?
La principal idea de esta librería es poder agregar traducciones a tus aplicaciones (frontend (WebAssembly) o backend) de una manera sencilla, estructurada y segura para que puedas manejar tus aplicaciones en varios idiomas sin tener que preocuparte por la traducción de los textos, evitando errores de traducción, ademas de poder manejar errores personalizados en varios idiomas, poder responder a los usuarios en su idioma nativo sin necesidad de usar archivos externos o llamadas a api.

## Características
- Traducciones seguras para uso concurrente
- Múltiples idiomas
- Formateo de cadenas con varios tipos de datos
- interfaz de error incorporada
- sin dependencias externas

## Instalación
```bash
go get github.com/cdvelop/lang
```

## Uso Básico

```go
package main

import (
    "fmt"
   . "github.com/cdvelop/lang"
)

type Handler struct {
    *Lang
}

func main() {

    // Inicializar traductor en tu manejador
    h := Handler{
          Lang: New(),
        }
    // por defecto el lengua es inglés puedes cambiarlo el idioma de tu preferencia ej: español
    h.SetDefaultLanguage("es")
 
    // al importar la librería con . accedes directamente a la variable D que contiene las traducciones ej: D.Language
    
    msg := translator.T(D.Language, "test", lang.D.NotSupported)
    fmt.Println(msg) // Salida: idioma test no soportado
}
```

## Componentes Principales



### Diccionario
El paquete incluye un diccionario predefinido con términos comunes en dire idiomas. Se pueden agregar nuevos términos extendiendo la estructura `dictionary`.

## Referencia de API

### New()
```go
func New(params ...any) *Lang
```
Crea una nueva instancia de Lang. Parámetros opcionales:
- sync.Mutex para seguridad concurrente
- io.Writer para salida personalizada

### SetDefaultLanguage()
```go
func (l *Lang) SetDefaultLanguage(language string) error
```
Establece el idioma predeterminado para traducciones. Retorna error si el idioma no es soportado.

### T()
```go
func (l Lang) T(args ...any) string
```
Retorna cadena traducida. Acepta múltiples argumentos de diferentes tipos:
- Cadenas (traducidas si están en el diccionario)
- Números
- Booleanos
- Errores
- Tipos personalizados (convertidos a cadena)

### Err()
```go
func (l Lang) Err(args ...any) error
```
Retorna un mensaje de error traducido.

### Print()
```go
func (l Lang) Print(args ...any)
```
Escribe cadena traducida al writer configurado.

## Ejemplos

### Uso Concurrente
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

### Writer Personalizado
```go
translator := lang.New(os.Stdout)
translator.Print(D.Hello, D.World)
```

## Idiomas Soportados
- Inglés (en)
- Español (es)
- Francés (fr)
- Alemán (de)
- Italiano (it)
- Portugués (pt)
- Ruso (ru)
- Chino (zh)
- Japonés (ja)
- Coreano (ko)
- Árabe (ar)
- Hindi (hi)
- Bengalí (bn)
- Turco (tr)
- Vietnamita (vi)
- Tailandés (th)
- Indonesio (id)
- Hebreo (he)
- Polaco (pl)
- Neerlandés (nl)
- Sueco (sv)
- Danés (da)
- Noruego (no)
- Finlandés (fi)
- Checo (cs)
- Húngaro (hu)
- Rumano (ro)
- Griego (el)
- Ucraniano (uk)
- Croata (hr)
- Eslovaco (sk)
- Esloveno (sl)
- Serbio (sr)
- Búlgaro (bg)
- Macedonio (mk)
- Albanés (sq)
- Lituano (lt)
- Letón (lv)
- Estonio (et)
- Malayo (ms)

## Contribuciones
Para agregar nuevos idiomas o términos:
1. Extender la estructura dictionary en dictionary.go
2. Agregar etiquetas de idioma con traducciones
3. Actualizar langSupported en la función New()
