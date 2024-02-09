package str_case

import (
	"bufio"
	"slices"
	"strings"
)

func Apa(value string) string {
	if strings.TrimSpace(value) == "" {
		return value
	}

	minorWords := []string{"and", "as", "but", "if", "nor", "or", "so", "yet", "a", "an", "the", "at", "by", "for", "in", "of", "off", "on", "per", "to", "up", "via"}
	endPunct := []rune{'.', '!', '?', ':', '-', ','}

	scanner := bufio.NewScanner(strings.NewReader(Title(value)))
	scanner.Split(bufio.ScanWords)
	result := new(strings.Builder)
	foundEndPunct := false
	
	if scanner.Scan() {
		result.Write(scanner.Bytes())
	}

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

	return result.String()
}