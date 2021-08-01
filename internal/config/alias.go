package config

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"

	"gopkg.in/yaml.v2"
)

type Alias struct {
	Command     []string
	Description string
	Flags       map[string]string
}

type AliasConfig map[string]*Alias

func (c AliasConfig) TraverseSorted(f func(name string, alias *Alias) error) error {
	names := make([]string, 0)
	for name := range c {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		if err := f(name, c[name]); err != nil {
			return err
		}
	}
	return nil
}

func Aliases() AliasConfig {
	// TODO handle error
	if config, err := loadAliases(); err == nil {
		return *config
	}
	return AliasConfig{}
}

func loadAliases() (config *AliasConfig, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(fmt.Sprintf("%v/.config/gj/alias.yaml", home))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &config)
	if err == nil && config == nil { // empty file
		config = &AliasConfig{}
	}
	return
}

func AddAlias(name string, alias *Alias) error {
	aliases := Aliases()

	aliases[name] = alias

	out, err := yaml.Marshal(aliases)
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// TODO create dirs
	return ioutil.WriteFile(fmt.Sprintf("%v/.config/gj/alias.yaml", home), out, fs.ModePerm)
}

func DeleteAlias(name string) error {
	aliases := Aliases()

	delete(aliases, name)

	out, err := yaml.Marshal(aliases)
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// TODO create dirs
	return ioutil.WriteFile(fmt.Sprintf("%v/.config/gj/alias.yaml", home), out, fs.ModePerm)
}
