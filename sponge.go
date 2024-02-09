package str_case

import (
	"math/rand/v2"
	"strings"
	"unicode"
)

func Sponge(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}

	s := strings.NewReader(strings.TrimSpace(value))
	result := new(strings.Builder)
	for s.Len() > 0 {
		r, _, _ := s.ReadRune()

		if rand.Float32() > 0.5 {
			result.WriteRune(unicode.ToUpper(r))
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}