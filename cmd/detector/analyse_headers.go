package detector

import (
	"fmt"
	"github.com/yildizozan/gandalf/cmd/config"
	"net/http"
)

func analyseHeaders(rules *config.Header, header *http.Header, c chan bool) {

	for k, v := range *header {
		for rk, _ := range *rules {
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
