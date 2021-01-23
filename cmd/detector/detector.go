package detector

import (
	"fmt"
	"regexp"
)

func Analyse(uri map[string][]string) bool {
	var re = regexp.MustCompile(`(?m)['"]`)

	for key, values := range uri {
		fmt.Println("Key:", key, "Value:", values)
		for _, value := range values {
			for i, match := range re.FindAllString(value, -1) {
				fmt.Println(match, "found at index", i)
				return true
			}
		}

	}

	return false
}

func AnalyseRawQuery(uri string) bool {
	var re = regexp.MustCompile(`(?i)(\%27)|(\')|(\-\-)|(\%23)|(\#)`)

	if len(re.FindStringIndex(uri)) > 0 {
		return true
	}

	return false
}
