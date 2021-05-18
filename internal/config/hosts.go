package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type HostConfig struct {
	User   string
	Token  string
	Cookie string
}

func Hosts() (config map[string]HostConfig, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(fmt.Sprintf("%v/.config/gj/hosts.yaml", home))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &config)
	return
}
