package api

import (
	"github.com/andygrunwald/go-jira"
)

func FindUsers(host string, username string) ([]jira.User, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	users, _, err := client.User.Find(username)
	return users, ApiError(err)
}
