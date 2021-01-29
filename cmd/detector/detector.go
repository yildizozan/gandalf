package detector

import (
	"fmt"
	"github.com/spf13/viper"
	config "github.com/yildizozan/gandalf/cmd/config/v2"
	"github.com/yildizozan/gandalf/cmd/metrics"
	"net/http"
)

func Analyse(req *http.Request) bool {

	chanXSS := make(chan bool)
	chanSQL := make(chan bool)
	chanHeader := make(chan bool)
	chanIp := make(chan bool)
	chanPath := make(chan bool)

	go analyseXSS(req.URL.RawQuery, chanXSS)
	go analyseSQLInjection(req.URL.RawQuery, chanSQL)
	go analyseHeaders(&config.Config.Header, &req.Header, chanHeader)
	go analyseIp(&config.Config.Ip, &req.RemoteAddr, chanIp)
	go analysePath(&config.Config.Path, &req.URL.Path, chanPath)

	xss := <-chanXSS
	if xss {
		metrics.HttpRequestsXSSVulnerable.WithLabelValues(config.Config.Name, req.Proto, "400")
	}

	sql := <-chanSQL
	if sql {
		metrics.HttpRequestsSQLInjectionVulnerable.WithLabelValues(config.Config.Name, req.Proto, "400")
	}

	header := <-chanHeader
	if header {
		metrics.HttpRequestsHeaderFilter.WithLabelValues(config.Config.Name, req.Proto, "400")
	}

	ip := <-chanIp
	if ip {
		metrics.HttpRequestsIpBlacklist.WithLabelValues(config.Config.Name, req.Proto, "400")
	}

	path := <-chanPath
	if path {
		metrics.HttpRequestsPathFiler.WithLabelValues(config.Config.Name, req.Proto, "400")
	}

	if viper.GetBool("verbose") {
		fmt.Printf("App: %s => XSS: %t, SQL: %t, Header: %t, IP: %t, Path %t\n",
			viper.GetString("app.name"), xss, sql, header, ip, path)
	}

	return xss || sql || header || ip || path
}

/*
// For presentation template
func XAnalyse(req *http.Request) bool {

	chanXSS := make(chan bool)
	chanSQL := make(chan bool)
	chanHeader := make(chan bool)
	chanIp := make(chan bool)
	chanPath := make(chan bool)

	go analyseXSS(req.URL.RawQuery, chanXSS)
	go analyseSQLInjection(req.URL.RawQuery, chanSQL)
	go analyseHeaders(&config.Config.Header, &req.Header, chanHeader)
	go analyseIp(&config.Config.Ip, &req.RemoteAddr, chanIp)
	go analysePath(&config.Config.Path, &req.URL.Path, chanPath)

	xss, sql, header, ip, path := <-chanXSS, <-chanSQL, <-chanHeader, <-chanIp, <-chanPath

	metrics.Collect(xss, sql, header, ip, path)

	if viper.GetBool("verbose") {
		fmt.Printf("App: %s => XSS: %t, SQL: %t, Header: %t, IP: %t, Path %t\n",
			viper.GetString("app.name"), xss, sql, header, ip, path)
	}

	return xss || sql || header || ip || path
}
*/
