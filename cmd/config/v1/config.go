package v1

import (
	"fmt"
	"github.com/prometheus/common/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Configurations struct {
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
