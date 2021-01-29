package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v2 "github.com/yildizozan/gandalf/cmd/config/v2"
	"github.com/yildizozan/gandalf/cmd/fasthttp"
	"github.com/yildizozan/gandalf/cmd/nethttp"
	"os"
)

var flagVerbose bool

var cmdRoot = &cobra.Command{
	Use:   "gandalf",
	Short: "Gandalf is web application firewall",
	//Run: func(cmd *cobra.Command, args []string) {
	//	nethttp.Start()
	//},
}

var cmdNetHttp = &cobra.Command{
	Use:   "nethttp",
	Short: "Nethttp anything to the screen",
	Run: func(cmd *cobra.Command, args []string) {
		nethttp.Start()
	},
}

var cmdFastHttp = &cobra.Command{
	Use:   "fasthttp",
	Short: "Fasthttp anything to the screen",
	Run: func(cmd *cobra.Command, args []string) {
		fasthttp.Start()
	},
}

func Execute() {
	cmdRoot.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "verbose output")
	cmdRoot.AddCommand(cmdNetHttp)
	cmdRoot.AddCommand(cmdFastHttp)
	cmdRoot.MarkPersistentFlagRequired("port")
	if err := cmdRoot.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Flags
	cmdRoot.PersistentFlags().Int16P("port", "p", 8080, "listening port")

	viper.Set("Verbose", true)

	// Set the file name of the configurations file
	viper.SetConfigName("gandalf")

	// Set the path to look for the configurations file
	viper.AddConfigPath("/etc/gandalf.yml")
	viper.AddConfigPath("$HOME/gandalf.yml")
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	// Set flags
	viper.BindPFlag("port", cmdRoot.PersistentFlags().Lookup("port"))

	// Set undefined variables
	viper.SetDefault("version", "1.0.0")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	err := viper.Unmarshal(&v2.Config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	//fmt.Println("Reading variables using the model..")
	//fmt.Println("Version is\t", configuration.Version)
	//fmt.Println("Name is\t", configuration.App.Name)
	//fmt.Println("Host is\t", configuration.App.Host)
	//fmt.Println("Log is\t", configuration.App.Logger)
	//fmt.Println("Rules is\t", configuration.App.Rules)
	//fmt.Println("Header is\t", detector.Config.App.Rules.Header)
	//fmt.Println("Ip is\t", configuration.App.Rules.Ip)
	//fmt.Println("Path is\t", configuration.App.Rules.Path)
	//fmt.Println("Http is\t", configuration.Spec.Http)
}
