package str_case

import (
	"strings"
	"unicode"
)

func UcFirst(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}

	val := strings.NewReader(value)
	result := new(strings.Builder)
	for val.Len() > 0 {
		s, _, _ := val.ReadRune()

		if unicode.IsLetter(s) {
			result.WriteRune(unicode.ToUpper(s))
			break
		}
		
		result.WriteRune(s)
	}

	val.WriteTo(result)

	return result.String()
}