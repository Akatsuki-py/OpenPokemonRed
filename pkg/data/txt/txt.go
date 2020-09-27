package txt

import (
	"pokered/pkg/util"
	"strings"
)

// Compile raw text data
func Compile(src string) string {
	s := strings.Split(src, util.LF())[1:]
	result := ""
	cout := false
	for _, elm := range s {
		switch elm {
		case "":
			result += "\\p"
			cout = false
		case "▼":
			result += "\\▼"
		default:
			if cout {
				result += elm + "\\c"
			} else {
				result += elm + "\\n"
				cout = true
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
