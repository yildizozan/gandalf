package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/yildizozan/gandalf/cmd/adaptor"
	"net/http"
)

var r = prometheus.NewRegistry()

var ActiveHttpRequests = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "gandalf_http_requests_active",
		Help: "Number of active http requests.",
	}, []string{})

var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gandalf_http_requests_total",
		Help: "Number of total http requests.",
	}, []string{"app", "scheme", "code"})

var HttpRequestsTotalVulnerable = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gandalf_http_requests_total_vulnerable",
		Help: "Number of total vulnerable http requests.",
	}, []string{"app", "scheme", "code"})

func Init() {
	r.MustRegister(HttpRequestsTotal)
	r.MustRegister(HttpRequestsTotalVulnerable)
	r.MustRegister(ActiveHttpRequests)
}

func NetHttpHandle() {
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
}

func FastHttpHandle() fasthttp.RequestHandler {
	return adaptor.NewFastHTTPHandler(promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
}
