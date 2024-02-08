package str_case

import (
	"strings"
	"unicode"
)

func Studly(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}
	
	r := strings.NewReplacer("_", " ", "-", " ")
	values := strings.NewReader(r.Replace(strings.TrimSpace(value)))
	result := new(strings.Builder)
	isSpace := false
	for values.Len() > 0 {
		v, _, _ := values.ReadRune()
		
		if unicode.IsSpace(v) {
			isSpace = true
			continue
		}

		if unicode.IsLetter(v) {
			if isSpace {
				isSpace = false
				result.WriteRune(unicode.ToUpper(v))
				continue
			}
		}

		result.WriteRune(v)
	}

	return UcFirst(result.String())
}