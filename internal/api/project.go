package api

import (
	"github.com/andygrunwald/go-jira"
)

type ListProjectOptions struct {
	Category []string
}

func ListProjects(host string, category []string) ([]jira.ProjectInfo, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	projects, _, err := client.Project.GetList()
	if category == nil || len(category) == 0 {
		return projects, nil
	}

	result := make([]jira.ProjectInfo, 0)
	for _, project := range projects {
		for _, c := range category {
			if c == "" || c == project.ProjectCategory.Name {
				result = append(result, project)
			}
		}
	}
	return result, ApiError(err)
}

func GetProject(host string, project string) (*jira.Project, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	jiraProject, _, err := client.Project.Get(project)

	return jiraProject, ApiError(err)
}

func ListProjectCategories(host string) ([]jira.ProjectCategory, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	req, _ := client.NewRequest("GET", "rest/api/2/projectCategory", nil)

	projectCategories := new([]jira.ProjectCategory)
	_, err = client.Do(req, projectCategories)
	if err != nil {
		return nil, err
	}
	return *projectCategories, nil
}
