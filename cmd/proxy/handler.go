package proxy

import (
	"github.com/spf13/viper"
	"github.com/yildizozan/gandalf/cmd/detector"
	"github.com/yildizozan/gandalf/cmd/logger"
	"github.com/yildizozan/gandalf/cmd/metrics"
	"net/http"
	"net/http/httputil"
)

var Proxy *httputil.ReverseProxy

func Handler(res http.ResponseWriter, req *http.Request) {
	metrics.HttpRequestsTotal.WithLabelValues(viper.GetString("app.name"), "http", "200").Inc()

	// Detector
	result := detector.Analyse(req)
	if result {
		metrics.HttpRequestsTotalVulnerable.WithLabelValues(viper.GetString("app.name"), "http", "400").Inc()

		res.WriteHeader(http.StatusBadRequest)
		_, err := res.Write([]byte("You shall not pass! - Gandalf"))
		if err != nil {
			logger.Error("Response not sending!")
		}

		l := logger.CreateLog(req.RemoteAddr, req.Method, req.RequestURI, req.Proto, "400")
		logger.Logger(viper.GetString("app.name"), true, "", l)
		return
	}

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	Proxy.ServeHTTP(res, req)

	l := logger.CreateLog(req.RemoteAddr, req.Method, req.RequestURI, req.Proto, "200")
	logger.Logger(viper.GetString("app.name"), true, "", l)
}
