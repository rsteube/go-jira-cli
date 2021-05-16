package api

import (
	"errors"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/rsteube/go-jira-cli/internal/config"
)

func NewClient(host string) (*jira.Client, error) {
	hosts, err := config.Hosts()
	if err != nil {
		return nil, err
	}
	if _, ok := hosts[host]; !ok {
		return nil, errors.New("unknown host")
	} else {
		// TODO basic auth
		return jira.NewClient(nil, "https://"+host)
	}
}

func ApiError(err error) error {
	if err != nil {
		return errors.New(strings.SplitN(err.Error(), "\n", 2)[0])
	}
	return err
}

func String(s string) *string {
	return &s
}
