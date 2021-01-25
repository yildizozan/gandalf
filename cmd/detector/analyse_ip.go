package detector

import (
	"github.com/spf13/viper"
	"strings"
)

func analyseIp(remoteAddr *string, c chan bool) {
	ip := strings.Split(*remoteAddr, ":")[0]

	ruleIpWhite := viper.GetStringSlice("app.rules.ip.whitelist")
	ruleIpBlack := viper.GetStringSlice("app.rules.ip.blacklist")
	//fmt.Println(ruleIpWhite)
	//fmt.Println(ruleIpBlack)
	for _, k := range ruleIpWhite {
		if k == ip {
			c <- false
			return
		}
	}

	for _, k := range ruleIpBlack {
		if k == ip {
			c <- true
			return
		}
	}

	c <- false
}
