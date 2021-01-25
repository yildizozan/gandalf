package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yildizozan/gandalf/cmd/config"
	"github.com/yildizozan/gandalf/cmd/detector"
	"github.com/yildizozan/gandalf/cmd/log"
	"github.com/yildizozan/gandalf/cmd/metrics"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func handler(res http.ResponseWriter, req *http.Request) {
	appName := viper.GetString("app.name")
	metrics.HttpRequestsTotal.WithLabelValues(appName, "http", "200").Inc()

	// Detector
	result := detector.Analyse(req)
	if result {
		metrics.HttpRequestsTotalVulnerable.WithLabelValues(appName, "http", "400").Inc()

		res.WriteHeader(http.StatusBadRequest)
		_, err := res.Write([]byte("You shall not pass! - Gandalf"))
		if err != nil {
			log.Error("Response not sending!")
		}
		return
	}

	// Target
	uri := fmt.Sprintf("http://%s:%s%s", "localhost", viper.GetString("app.port"), req.RequestURI)
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

func main() {
	viper.Set("Verbose", true)

	// Set the file name of the configurations file
	viper.SetConfigName("gandalf")

	// Set the path to look for the configurations file
	viper.AddConfigPath("/etc/gandalf.yml")
	viper.AddConfigPath("$HOME/.gandalf.yml")
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	var configuration config.MyConfig

	// Set undefined variables
	viper.SetDefault("version", "1.1.1")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Version is\t", configuration.Version)
	fmt.Println("App is\t", configuration.App)
	fmt.Println("Rules is\t", configuration.App.Rules)
	fmt.Println("Path is\t", configuration.App.Rules.Path)
	//fmt.Println("Hosts is\t", configuration.Spec.Hosts)
	//fmt.Println("Http is\t", configuration.Spec.Http)
	fmt.Println()

	// Metrics
	metrics.Init()
	metrics.NetHttpHandle()

	// Gandalf
	var cw metrics.ConnectionWatcher
	http.HandleFunc("/", handler)
	server := &http.Server{
		Addr:      ":8080",
		ConnState: cw.OnStateChangeForNetHttp,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
