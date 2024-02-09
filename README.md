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
### Snake

Converts the given string to snake_case.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Snake("helloWorld") 

// hello_world
```
### Sponge

Converts the given string to sPoNGeCAse, is is usually used for creating a "Mocking SpongeBob" meme.

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
### UcFirst

Converts the given string with the first-character uppercased.

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.UcFirst("helloWorld") 

// HelloWorld
```
Running Tests
-------

```bash
go test -v
```
