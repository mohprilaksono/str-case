package str_case

import (
	"strings"
	"unicode"
)

func LcFirst(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}

	s := strings.NewReader(value)
	result := new(strings.Builder)
	result.Grow(len(value))
	for s.Len() > 0 {
		r, _, _ := s.ReadRune()

		if unicode.IsLetter(r) {
			result.WriteRune(unicode.ToLower(r))
			break
		}
		
		result.WriteRune(r)
	}

	s.WriteTo(result)

	return result.String()
}