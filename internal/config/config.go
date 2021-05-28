package config

import (
	"fmt"
	"os"
)

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
