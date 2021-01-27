package v2

// Parsed and filled with data by Viper.
var Config Configuration

type Logger struct {
	loki     string
	fluentd  string
	logstash string
}

type Path struct {
	Prefix string
	Exact  string
	Match  string
}

type Header map[string]string

type Ip struct {
	Whitelist []string
	Blacklist []string
}

type Rules struct {
	Ip
	Header
	Path
}

type App struct {
	Name string
	Host string
	Logger
	Rules
}

type Configuration struct {
	Version string
	App
}
