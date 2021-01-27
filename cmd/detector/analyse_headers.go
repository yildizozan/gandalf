package detector

import (
	config "github.com/yildizozan/gandalf/cmd/config/v2"
	"net/http"
	"strings"
)

func analyseHeaders(rules *config.Header, header *http.Header, c chan bool) {
	for k, _ := range *header {
		for rk, _ := range *rules {
			if strings.ToLower(k) == strings.ToLower(rk) {
				c <- true
				return
			}
		}
	}
	c <- false
}
