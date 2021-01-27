package detector

import (
	"fmt"
	"github.com/spf13/viper"
	v2 "github.com/yildizozan/gandalf/cmd/config/v2"
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
	go analyseHeaders(&v2.Config.Header, &req.Header, chanHeader)
	go analyseIp(&v2.Config.Ip, &req.RemoteAddr, chanIp)
	go analysePath(&v2.Config.Path, &req.URL.Path, chanPath)

	xss := <-chanXSS
	if xss {
		metrics.HttpRequestsXSSVulnerable.WithLabelValues(v2.Config.Name, req.Proto, "400")
	}

	sql := <-chanSQL
	if sql {
		metrics.HttpRequestsSQLInjectionVulnerable.WithLabelValues(v2.Config.Name, req.Proto, "400")
	}

	header := <-chanHeader
	if header {
		metrics.HttpRequestsHeaderFilter.WithLabelValues(v2.Config.Name, req.Proto, "400")
	}

	ip := <-chanIp
	if ip {
		metrics.HttpRequestsIpBlacklist.WithLabelValues(v2.Config.Name, req.Proto, "400")
	}

	path := <-chanPath
	if path {
		metrics.HttpRequestsPathFiler.WithLabelValues(v2.Config.Name, req.Proto, "400")
	}

	if viper.GetBool("verbose") {
		fmt.Printf("App: %s => XSS: %t, SQL: %t, Header: %t, IP: %t, Path %t\n",
			viper.GetString("app.name"), xss, sql, header, ip, path)
	}

	return xss || sql || header || ip || path
}
