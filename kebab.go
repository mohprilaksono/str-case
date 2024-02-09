package str_case

import (
	"strings"
	"unicode"
)

func Kebab(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}
	
	s := strings.NewReader(strings.TrimSpace(value))
	result := new(strings.Builder)
	isSpace := false
	isFirstLetter := false
	for s.Len() > 0 {
		r, _, _ := s.ReadRune()

		if !isFirstLetter {
			if unicode.IsLetter(r) {
				result.WriteRune(unicode.ToLower(r))
				isFirstLetter = true
			} else {
				result.WriteRune(r)
			}

			continue
		}

		if unicode.IsSpace(r) {
			isSpace = true
			continue
		}

		if unicode.IsLetter(r) {
			if isSpace {
				isSpace = false
				result.WriteRune('-')
				result.WriteRune(unicode.ToLower(r))
				continue
			}

			if unicode.IsUpper(r) {
				result.WriteRune('-')
				result.WriteRune(unicode.ToLower(r))
				continue
			}
		}
		
		result.WriteRune(r)
	}

	return result.String()
}