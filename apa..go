package str_case

import (
	"bufio"
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Apa(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}

	minorWords := []string{"and", "as", "but", "if", "nor", "or", "so", "yet", "a", "an", "the", "at", "by", "for", "in", "of", "off", "on", "per", "to", "up", "via"}
	endPunct := []rune{'.', '!', '?', ':', '-', ','}

	c := cases.Title(language.English)
	scanner := bufio.NewScanner(strings.NewReader(c.String(value)))
	scanner.Split(bufio.ScanWords)
	result := new(strings.Builder)
	foundEndPunct := false
	for scanner.Scan() {
		s := scanner.Text()
		
		if result.Len() > 0 {
			result.WriteRune(' ')
		}
		
		for _, ep := range endPunct {
			if strings.ContainsRune(s, ep) {
				foundEndPunct = true
				break
			}
		}

		if slices.Contains[[]string, string](minorWords, strings.ToLower(s)) {
			if foundEndPunct {
				result.WriteString(s)
				foundEndPunct = false
			} else {
				result.WriteString(strings.ToLower(s))
			}

			continue
		}

		result.WriteString(s)
	}

	return UcFirst(result.String())
}