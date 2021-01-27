package nethttp

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yildizozan/gandalf/cmd/logger"
	"github.com/yildizozan/gandalf/cmd/metrics"
	"github.com/yildizozan/gandalf/cmd/proxy"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Start() {
	// Metrics
	metrics.Init()
	metrics.NetHttpHandle()

	// Proxy
	uri, err := url.Parse(viper.GetString("app.host"))
	if viper.GetBool("verbose") {
		fmt.Println(uri)
	}
	if err != nil {
		logger.Log("Could not parse downstream url: %s", viper.GetString("app.name"))
	}

	proxy.Proxy = httputil.NewSingleHostReverseProxy(uri)

	director := proxy.Proxy.Director
	proxy.Proxy.Director = func(req *http.Request) {
		director(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = req.URL.Host
	}

	// Gandalf
	var cw metrics.ConnectionWatcher
	http.HandleFunc("/", proxy.Handler)
	addr := fmt.Sprintf(":%d", viper.GetInt("port"))
	server := &http.Server{
		Addr:      addr,
		ConnState: cw.OnStateChangeForNetHttp,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
