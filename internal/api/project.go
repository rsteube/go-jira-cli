package api

import (
	"github.com/andygrunwald/go-jira"
)

type ListProjectOptions struct {
	Category []string
}

type ProjectListEntry struct {
	Expand          string               `json:"expand" structs:"expand"`
	Self            string               `json:"self" structs:"self"`
	ID              string               `json:"id" structs:"id"`
	Key             string               `json:"key" structs:"key"`
	Name            string               `json:"name" structs:"name"`
	AvatarUrls      jira.AvatarUrls      `json:"avatarUrls" structs:"avatarUrls"`
	ProjectTypeKey  string               `json:"projectTypeKey" structs:"projectTypeKey"`
	ProjectCategory jira.ProjectCategory `json:"projectCategory,omitempty" structs:"projectsCategory,omitempty"`
	IssueTypes      []jira.IssueType     `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
}

func ListProjects(host string, category []string) ([]ProjectListEntry, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	projects, _, err := client.Project.GetList()
	result := make([]ProjectListEntry, 0)
	for _, project := range *projects {
		if category == nil || len(category) == 0 {
			result = append(result, ProjectListEntry(project))
			continue
		}

		for _, c := range category {
			if c == "" || c == project.ProjectCategory.Name {
				result = append(result, ProjectListEntry(project))
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
