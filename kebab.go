package str_case

import (
	"strings"
	"unicode"
)

func Kebab(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}
	
	values := strings.NewReader(LcFirst(strings.TrimSpace(value)))
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
				result.WriteRune('-')
				result.WriteRune(unicode.ToLower(v))
				continue
			}

			if unicode.IsUpper(v) {
				result.WriteRune('-')
				result.WriteRune(unicode.ToLower(v))
				continue
			}
		}
		
		result.WriteRune(v)
	}

	return result.String()
}