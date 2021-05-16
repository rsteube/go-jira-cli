package api

import (
	"github.com/andygrunwald/go-jira"
)

func ListProjects(host string) (*jira.ProjectList, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	projects, _, err := client.Project.GetList()
	return projects, ApiError(err)
}
