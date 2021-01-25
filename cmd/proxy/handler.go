package proxy

import (
	"fmt"
	"github.com/yildizozan/gandalf/cmd/config"
	"github.com/yildizozan/gandalf/cmd/detector"
	"github.com/yildizozan/gandalf/cmd/log"
	"github.com/yildizozan/gandalf/cmd/metrics"
	"net/http"
	"net/http/httputil"
	"net/url"
)

//var conf = config.App{
//	Name:  viper.GetString("app.port"),
//	Host:  viper.GetString("app.host"),
//}

var conf = config.App{
	Name: "app",
	Host: "localhost:3000",
}

func Handler(res http.ResponseWriter, req *http.Request) {

	metrics.HttpRequestsTotal.WithLabelValues(conf.Name, "http", "200").Inc()

	// Detector
	result := detector.Analyse(req)
	if result {
		metrics.HttpRequestsTotalVulnerable.WithLabelValues(conf.Name, "http", "400").Inc()

		res.WriteHeader(http.StatusBadRequest)
		_, err := res.Write([]byte("You shall not pass! - Gandalf"))
		if err != nil {
			log.Error("Response not sending!")
		}
		return
	}

	// Target
	uri := fmt.Sprintf("http://%s%s", conf.Host, req.RequestURI)
	target, err := url.Parse(uri)
	if err != nil {
		log.Log(fmt.Sprintf("%s\n", err))
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Update the headers to allow for SSL redirection
	req.URL.Host = target.Host
	req.URL.Scheme = "http"
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = target.Host
	req.URL.Path = target.Path
	req.RequestURI = ""

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
