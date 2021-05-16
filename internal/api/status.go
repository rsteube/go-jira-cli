package api

import (
	"github.com/andygrunwald/go-jira"
)

func ListStatuses(host string) ([]jira.Status, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	statuses, _, err := client.Status.GetAllStatuses()
	return statuses, ApiError(err)
}
