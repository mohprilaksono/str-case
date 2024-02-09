package str_case

import (
	"strings"
	"unicode"
)

func Swap(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}

	s := strings.NewReader(value)
	result := new(strings.Builder)
	for s.Len() > 0 {
		r, _, _ := s.ReadRune()

		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				result.WriteRune(unicode.ToLower(r))
			} else {
				result.WriteRune(unicode.ToUpper(r))
			}
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}