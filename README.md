str_case
==========

a Go library for string case manipulation.

Install
-------

```bash
go get github.com/mohprilaksono/str-case
```

Example
-------

```go
import str_case "github.com/mohprilaksono/str-case"

str_case.Apa("Thomas And Friends") //Thomas and Friends
```

Usage
-------

### Ada

Converts the given string to Ada_Case.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Ada("helloWorld") 

// Hello_World
```
### Apa

Converts the given string to Title Case, following the [APA Guidelines](https://apastyle.apa.org/style-grammar-guidelines/capitalization/title-case).

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Apa("Thomas And Friends") 

// Thomas and Friends
```
### Camel

Converts the given string to camelCase.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Camel("hello_world") 

// helloWorld
```
### Cobol

Converts the given string to COBOL-CASE.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Cobol("helloWorld") 

// HELLO-WORLD
```
### Kebab

Converts the given string to kebab-case.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Kebab("helloWorld") 

// hello-world
```
### LcFirst

Converts the given string with the first-character lowercased.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.LcFirst("HelloWorld") 

// helloWorld
```
### Macro

Converts the given string to MACRO_CASE.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Macro("helloWorld") 

// HELLO_WORLD
```
### Snake

Converts the given string to snake_case.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Snake("helloWorld") 

// hello_world
```
### Sponge

Converts the given string to sPoNGeCAse, it is usually used for creating a "Mocking SpongeBob" meme.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Sponge("HTML is not a programming language") 

// hTmL Is nOt A ProGraMmINg LanGuAgE
```
### Studly

Converts the given string to StudlyCase.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Studly("hello_world") 

// HelloWorld
```
### Swap

Swap the case value of the given string.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Swap("HelloWorlD") 

// hELLOwORLd
```
### Title

Converts the given string to Title Case.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Title("hello world") 

// Hello World
```
### Train

Converts the given string to Train-Case.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Train("helloWorld") 

// Hello-World
```
### UcFirst

Converts the given string with the first-character uppercased.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.UcFirst("helloWorld") 

// HelloWorld
```
Dependencies
------------

### Build dependencies

* none

### Test dependencies

* `github.com/stretchr/testify`

Running Tests
--------------

```bash
go test -v
```
