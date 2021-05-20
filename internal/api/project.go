package api

import (
	"github.com/andygrunwald/go-jira"
)

func ListProjects(host string) ([]jira.ProjectInfo, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	projects, _, err := client.Project.GetList()
	return projects, ApiError(err)
}

func GetProject(host string, project string) (*jira.Project, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	jiraProject, _, err := client.Project.Get(project)

	return jiraProject, ApiError(err)
}
