package detector

import "regexp"

const model = `(?i)(\%27)|(\')|(\-\-)|(\%23)|(\#)`

func analyseRawQuery(uri string, channel chan bool) {
	var re = regexp.MustCompile(model)

	if len(re.FindStringIndex(uri)) > 0 {
		channel <- true
		return
	}

	channel <- false
}