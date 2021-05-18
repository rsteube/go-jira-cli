package api

import (
	"github.com/andygrunwald/go-jira"
)

func ListResolutions(host string) ([]jira.Resolution, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	resolutions, _, err := client.Resolution.GetList()
	if err != nil {
		return nil, ApiError(err)
	}
	return resolutions, nil
}
