package detector

import (
	config "github.com/yildizozan/gandalf/cmd/config/v2"
	"strings"
)

func analyseIp(rules *config.Ip, remoteAddr *string, c chan bool) {
	ip := strings.Split(*remoteAddr, ":")[0]

	//fmt.Println(ruleIpWhite)
	//fmt.Println(ruleIpBlack)
	for _, k := range rules.Whitelist {
		if k == ip {
			c <- false
			return
		}
	}

	for _, k := range rules.Blacklist {
		if k == ip {
			c <- true
			return
		}
	}

	c <- false
}
