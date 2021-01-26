package detector

import (
	"fmt"
	"github.com/spf13/viper"
	config "github.com/yildizozan/gandalf/cmd/config/v2"
	"net/http"
)

var ruleIp = config.Ip{
	Whitelist: viper.GetStringSlice("app.rules.ip.whitelist"),
	Blacklist: viper.GetStringSlice("app.rules.ip.blacklist"),
}

var ruleHeader config.Header = viper.GetStringMapStringSlice("app.rules.header")

var rulePath = config.Path{
	Prefix: viper.GetString("app.rules.path"),
	Exact:  viper.GetString("app.rules.path.exact"),
	Match:  viper.GetString("app.rules.path.match"),
}

func Analyse(req *http.Request) bool {

	chanQuery := make(chan bool)
	chanHeader := make(chan bool)
	chanIp := make(chan bool)
	chanPath := make(chan bool)

	go analyseRawQuery(req.URL.RawQuery, chanQuery)
	go analyseHeaders(&req.Header, chanHeader)
	go analyseIp(&ruleIp, &req.RemoteAddr, chanIp)
	go analysePath(&rulePath, &req.URL.Path, chanPath)

	query, header, ip, path := <-chanQuery, <-chanHeader, <-chanIp, <-chanPath
	if viper.GetBool("verbose") {
		fmt.Printf("App: %s, Query: %t, Header: %t, IP: %t, Path %t\n", viper.GetString("app.name"), query, header, ip, path)
	}

	return query || header || ip || path
}
