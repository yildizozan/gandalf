package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var ActiveHttpRequests = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "gandalf_http_requests_active",
		Help: "Number of active http requests.",
	}, []string{})

var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gandalf_http_requests_total",
		Help: "Number of total http requests.",
	}, []string{"http_type", "code"})

var HttpRequestsTotalVulnerable = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gandalf_http_requests_total_vulnerable",
		Help: "Number of total vulnerable http requests.",
	}, []string{"http_type", "code"})

func Init() {
	r := prometheus.NewRegistry()
	r.MustRegister(HttpRequestsTotal)
	r.MustRegister(HttpRequestsTotalVulnerable)
	r.MustRegister(ActiveHttpRequests)
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

}
