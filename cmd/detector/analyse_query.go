package detector

import "regexp"

func analyseSQLInjection(uri string, channel chan bool) {
	var re = regexp.MustCompile(model)

	if len(re.FindStringIndex(uri)) > 0 {
		channel <- true
		return
	}

	channel <- false
}
