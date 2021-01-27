package detector

import "regexp"

func analyseXSS(uri string, channel chan bool) {
	var re = regexp.MustCompile(model)

	if len(re.FindStringIndex(uri)) > 0 {
		channel <- true
		return
	}

	channel <- false
}
