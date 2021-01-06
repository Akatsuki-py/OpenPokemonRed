package text

import (
	"strings"
)

func preprocess(src string) string {
	// Normalize newlines
	src = strings.Replace(src, "\r\n", "\n", -1)
	s := strings.Split(src, "\n")[1:]
	result := ""
	cont := false
	for _, line := range s {
		switch line {
		case "":
			result += "\\p"
			cont = false
		case "▼":
			result += "\\▼"
		default:
			if cont {
				result += line + "\\c"
			} else {
				result += line + "\\n"
				cont = true
			}
		}
	}
	if strings.HasSuffix(result, "\\p") {
		result += "\\d"
	}
	result = clean(result)
	return result
}

func clean(str string) string {
	str = strings.ReplaceAll(str, "\\n\\p", "\\p")
	str = strings.ReplaceAll(str, "\\p\\d", "\\d")
	str = strings.ReplaceAll(str, "\\c\\p", "\\p")
	str = strings.ReplaceAll(str, "\\c\\d", "\\d")
	str = strings.ReplaceAll(str, "\\n\\▼", "\\▼")
	str = strings.ReplaceAll(str, "\\c\\▼", "\\▼")
	str = strings.TrimSuffix(str, "\\c")
	str = strings.TrimSuffix(str, "\\n")
	str = strings.TrimSuffix(str, "\\p")
	return str
}
