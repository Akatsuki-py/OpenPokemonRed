package text

import (
	"pokered/pkg/util"
	"strings"
)

func preprocess(src string) string {
	s := strings.Split(src, util.LF())[1:]
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
	result = clean(result)
	return result
}

func clean(str string) string {
	str = strings.ReplaceAll(str, "\\n\\p", "\\p")
	str = strings.ReplaceAll(str, "\\c\\p", "\\p")
	str = strings.ReplaceAll(str, "\\n\\▼", "\\▼")
	str = strings.ReplaceAll(str, "\\c\\▼", "\\▼")
	str = strings.TrimSuffix(str, "\\c")
	str = strings.TrimSuffix(str, "\\n")
	str = strings.TrimSuffix(str, "\\p")
	return str
}
