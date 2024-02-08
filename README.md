str_case
==========

a Go library for string case manipulation

Install
-------

```bash
go get -u github.com/mohprilaksono/str-case
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

Convert the given string to Apa case, following the [APA Guidelines](https://apastyle.apa.org/style-grammar-guidelines/capitalization/title-case)

```go
import str_case "github.com/mohprilaksono/str-case"

var value = str_case.Apa("Thomas And Friends") 

//Thomas and Friends
```