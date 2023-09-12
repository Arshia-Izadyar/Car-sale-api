package common

import "unicode"

func ToSnakeCase(str string) string {
	var result []rune
	lastCharWasUpper := false

	for _, char := range str {
		if unicode.IsUpper(char) {
			if !lastCharWasUpper {
				if len(result) > 0 {
					result = append(result, '_')
				}
			}
			lastCharWasUpper = true
		} else {
			lastCharWasUpper = false
		}
		result = append(result, unicode.ToLower(char))
	}

	return string(result)
}
