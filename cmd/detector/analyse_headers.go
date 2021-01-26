package detector

import (
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func analyseHeaders(header *http.Header, c chan bool) {
	rules := viper.GetStringMapStringSlice("app.rules.header")
	//fmt.Println(len(*header), len(rules))
	for k, _ := range *header {
		for rk, _ := range rules {
			if strings.ToLower(k) == strings.ToLower(rk) {
				c <- true
				return
			}
		}
	}
	c <- false
}
