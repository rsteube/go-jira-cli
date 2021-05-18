package api

import (
	"github.com/andygrunwald/go-jira"
)

func ListPriorities(host string) ([]jira.Priority, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	priorities, _, err := client.Priority.GetList()
	if err != nil {
		return nil, ApiError(err)
	}
	return priorities, nil
}
