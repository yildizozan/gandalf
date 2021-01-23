package main

import (
	"bufio"
	"fmt"
	"github.com/yildizozan/gandalf/cmd/config"
	"github.com/yildizozan/gandalf/cmd/detector"
	"github.com/yildizozan/gandalf/cmd/log"
	"github.com/yildizozan/gandalf/cmd/metrics"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var configFile config.Configurations
var rules []string

// TODO: Handler recursive çalışıyor reverse proxy 8080->8001
func handler(res http.ResponseWriter, req *http.Request) {
	metrics.HttpRequestsTotal.WithLabelValues("http", "200").Inc()

	str, _ := url.QueryUnescape(req.URL.RawQuery)

	// Detector
	result := detector.AnalyseRawQuery(str)

	if result {
		metrics.HttpRequestsTotalVulnerable.WithLabelValues("http", "400").Inc()

		res.WriteHeader(http.StatusBadRequest)
		_, err := res.Write([]byte("Vulnerable founded! You shall not pass!"))
		if err != nil {
			log.Error("Response not sending!")
		}
		return
	}

	// TODO: Get target from file
	uri := fmt.Sprintf("http://%s/%s", "rematriks.com:80", req.RequestURI)
	target, err := url.Parse(uri)
	if err != nil {
		log.Log(fmt.Sprintf("%s\n", err))
	}
	for k, v := range target.Query() {
		fmt.Printf("%s	%s\n", k, v)
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

func main() {
	var c config.Configurations
	configFile = *c.GetConf()
	fmt.Println(*c.GetConf())

	file, err := os.Open("rules.txt")
	if err != nil {
		log.Log(fmt.Sprintf("%s\n", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}
	log.Log(fmt.Sprintf("%d rules loaded!\n", len(rules)))

	if err := scanner.Err(); err != nil {
		log.Log(fmt.Sprintf("%s\n", err))
	}

	// Metrics
	metrics.Init()

	// Gandalf
	var cw metrics.ConnectionWatcher
	http.HandleFunc("/", handler)
	server := &http.Server{
		Addr:      ":8080",
		ConnState: cw.OnStateChange,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
