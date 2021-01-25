package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yildizozan/gandalf/cmd/config"
	"github.com/yildizozan/gandalf/cmd/metrics"
	"github.com/yildizozan/gandalf/cmd/proxy"
	"net/http"
	"os"
)

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

	var configuration config.Config

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
	fmt.Println("Name is\t", configuration.App.Name)
	fmt.Println("Host is\t", configuration.App.Host)
	fmt.Println("Rules is\t", configuration.App.Rules)
	//fmt.Println("Hosts is\t", configuration.Spec.Hosts)
	//fmt.Println("Http is\t", configuration.Spec.Http)
	fmt.Println()

	// Metrics
	metrics.Init()
	metrics.NetHttpHandle()

	// Gandalf
	var cw metrics.ConnectionWatcher
	http.HandleFunc("/", proxy.Handler)
	server := &http.Server{
		Addr:      ":8080",
		ConnState: cw.OnStateChangeForNetHttp,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
