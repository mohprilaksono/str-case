package str_case

import (
	"strings"
	"unicode"
)

func Title(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}

	s := strings.NewReader(value)
	result := new(strings.Builder)
	isNotLetter := false
	for s.Len() > 0 {
		r, _, _ := s.ReadRune()

		if !unicode.IsLetter(r) {
			result.WriteRune(r)
			if !isNotLetter {
				isNotLetter = true
			}
			continue
		}

		if isNotLetter {
			result.WriteRune(unicode.ToUpper(r))
			isNotLetter = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return UcFirst(result.String())
}