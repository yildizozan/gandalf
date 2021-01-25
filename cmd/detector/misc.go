package detector

import "regexp"

func regex(str string, regex string) bool {
	var re = regexp.MustCompile(regex)

	if len(re.FindStringIndex(str)) > 0 {
		return true
	}

	return false
}
