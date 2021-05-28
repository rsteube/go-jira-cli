package config

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Credentials struct {
	User   string            `yaml:"user,omitempty"`
	Token  string            `yaml:"token,omitempty"`
	Cookie map[string]string `yaml:"cookie,omitempty"`
}

type HostCredentials map[string]*Credentials

// TODO store locally to only load once
func Hosts() (hc HostCredentials, err error) {
	path, err := configPath("hosts.yaml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return make(HostCredentials), nil
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &hc)
	return
}

func (hc HostCredentials) Add(host string, config *Credentials) error {
	map[string]*Credentials(hc)[host] = config
	return hc.write()
}

func (hc HostCredentials) Remove(host string) error {
	delete(hc, host)
	return hc.write()
}

func (hc HostCredentials) write() error {
	path, err := configPath("hosts.yaml")
	if err != nil {
		return err
	}
	marshalled, err := yaml.Marshal(hc)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, marshalled, fs.ModePerm)
}

func configPath(filename string) (path string, err error) {
	var home string
	if home, err = os.UserHomeDir(); err == nil {
		dir := fmt.Sprintf("%v/.config/gj", home)
		if err = os.MkdirAll(dir, os.ModePerm); err == nil {
			path = fmt.Sprintf("%v/%v", dir, filename)
		}
	}
	return
}
