package v2

type Path struct {
	Prefix string
	Exact  string
	Match  string
}

type Header map[string][]string

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
	Name   string
	Host   string
	Logger string
	Rules
}

type Config struct {
	Version string
	App
}
