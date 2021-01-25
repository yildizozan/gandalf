package config

import (
	"fmt"
	"github.com/yildizozan/gandalf/cmd/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

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
	Name string
	Port int16
	Rules
}

type MyConfig struct {
	Version string
	App
}

/////////////////////////////

type Uri struct {
	Exact  string
	Prefix string
}

type Match []struct {
	Uri  Uri
	Host string
}

type Http []struct {
	Name        string
	Match       Match
	Destination string
}

type Spec struct {
	Hosts []string
	Http  Http
}

type Configurations struct {
	Version string
	Spec    Spec
}

type Config struct {
	Version string `yaml:"version"`
	Spec    struct {
		Hosts []string `yaml:"hosts"`
		Http  []struct {
			Name  string `yaml:"name"`
			Match []struct {
				Uri struct {
					Exact  string `yaml:"exact"`
					Prefix string `yaml:"prefix"`
				} `yaml:"uri"`
				Host string `yaml:"host"`
			} `yaml:"match"`
			Destination string `yaml:"destination"`
		} `yaml:"http"`
	} `yaml:"spec"`
}

func (c *Configurations) GetConf() *Configurations {
	var err error
	yamlFile, err := ioutil.ReadFile("nethttp.yml")
	if err != nil {
		message := fmt.Sprintf("nethttp.yml err   #%v \n", err)
		log.Error(message)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		message := fmt.Sprintf("Unmarshal: %v\n", err)
		log.Error(message)
	}

	return c
}