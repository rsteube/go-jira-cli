package config

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// TODO read/write config file functions
// TODO current impletementation fragile
type HostConfig struct {
	User   string            `yaml:"user,omitempty"`
	Token  string            `yaml:"token,omitempty"`
	Cookie map[string]string `yaml:"cookie,omitempty"`
}

// TODO handle missing file/directory
func Hosts() (config map[string]*HostConfig, err error) {
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

// TODO handle missing file/directory
func AddHost(host string, config *HostConfig) error {
	current, err := Hosts()
	if err != nil {
		return err
	}

	if current == nil { // TODO fix in Hosts
		current = make(map[string]*HostConfig)
	}

	current[host] = config

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	marshalled, err := yaml.Marshal(current)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fmt.Sprintf("%v/.config/gj/hosts.yaml", home), marshalled, fs.ModePerm)
}

func RemoveHost(host string) error {
	current, err := Hosts()
	if err != nil {
		return err
	}

	if current == nil { // TODO fix in Hosts
		current = make(map[string]*HostConfig)
	}

	delete(current, host)

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	marshalled, err := yaml.Marshal(current)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fmt.Sprintf("%v/.config/gj/hosts.yaml", home), marshalled, fs.ModePerm)
}
