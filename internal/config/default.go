package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type DefaultConfig struct {
	Editor string
	Host   string
	Pager  string
}

func Default() DefaultConfig {
	if config, err := loadDefault(); err == nil {
		return *config
	}
	return DefaultConfig{}
}

func loadDefault() (config *DefaultConfig, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(fmt.Sprintf("%v/.config/gj/default.yaml", home))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &config)
	return
}
