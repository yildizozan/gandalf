package detector

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func analyseHeaders(header http.Header, c chan bool) {
	rule := viper.GetStringMapStringSlice("app.rules.header")
	for k, v := range header {
		for rk, _ := range rule {
			if k == rk {
				//fmt.Println("k:", k, "v:", v, "rk:", rk, "rv:", rv)
			}
			//fmt.Println("k:", k, "v:", v, "rk:", rk, "rv:", rv)
		}
		fmt.Println("k:", k, "v:", v)
		if k == "Accept-Encoding" {
			for idx, vv := range v {
				fmt.Println(idx, vv)
			}
		}
	}
	c <- false
}
