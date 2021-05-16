package api

import (
	"github.com/andygrunwald/go-jira"
)

func ListIssueTypes(host string, project string) ([]jira.IssueType, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	jiraProject, _, err := client.Project.Get(project)
	if err != nil {
		return nil, ApiError(err)
	}
	return jiraProject.IssueTypes, nil
}
