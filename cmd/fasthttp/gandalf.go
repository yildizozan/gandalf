package fasthttp

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/yildizozan/gandalf/cmd/logger"
	"github.com/yildizozan/gandalf/cmd/metrics"
)

var proxyClient = &fasthttp.HostClient{
	Addr: "localhost:3000",
	// set other options here if required - most notably timeouts.
}

func ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
	metrics.HttpRequestsTotal.WithLabelValues("http", "200").Inc()
	/*
		// Detector
		result := detector.analyseRawQuery(ctx.URI().QueryArgs().String())
		if result {
			metrics.HttpRequestsTotalVulnerable.WithLabelValues("http", "400").Inc()
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.SetBody([]byte("You shall not pass! - Gandalf"))
			return
		}
	*/
	req := &ctx.Request
	res := &ctx.Response
	prepareRequest(req)
	if err := proxyClient.Do(req, res); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
	}
	postprocessResponse(res)
}

func prepareRequest(req *fasthttp.Request) {
	// do not proxy "Connection" header.
	req.Header.Del("Connection")
	// strip other unneeded headers.
	req.Header.Set("X-Forwarded-Host", string(req.Header.Peek("Host")))
	// alter other request params before sending them to upstream host
}

func postprocessResponse(resp *fasthttp.Response) {
	// do not proxy "Connection" header
	resp.Header.Del("Connection")
	// strip other unneeded headers
	resp.Header.Del("X-Forwarded-Host")
	// alter other response data if needed
}

func Start() {

	// Metrics
	metrics.Init()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) == "/metrics" {
			metrics.FastHttpHandle()
		} else {
			ReverseProxyHandler(ctx)
		}
	}
	var cw metrics.ConnectionWatcher
	server := &fasthttp.Server{
		ConnState: cw.OnStateChangeForFastHttp,
		Handler:   requestHandler,
	}
	addr := fmt.Sprintf(":%d", viper.GetInt("port"))
	if err := server.ListenAndServe(addr); err != nil {
		logger.Error("error in fasthttp server: %s", err)
	}
}
