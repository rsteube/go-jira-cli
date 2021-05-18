package api

import (
	"fmt"
	"strings"

	"github.com/andygrunwald/go-jira"
)

type ListIssuesOptions struct {
	Project        []string
	Type           []string
	Status         []string
	StatusCategory []string
	Assignee       []string
	Component      []string
	Priority       []string
	Fields         []string
	Filter         int
	Jql            string
	Query          string
}

func (o *ListIssuesOptions) toJql(host string) (string, error) {
	//project in (SZOPS, BA) AND issuetype in (Bug, CVE) AND status in ("In Progress", Reopened) AND assignee in (membersOf("Interner Benutzer"), membersOf(jira-developers))
	jql := make([]string, 0)

	if o.Filter > 0 {
		filter, err := GetFilter(host, o.Filter)
		if err != nil {
			return "", err
		}
		jql = append(jql, filter.Jql)
	}

	if o.Jql != "" {
		jql = append(jql, o.Jql)
	}
	if o.Component != nil && len(o.Component) > 0 {
		jql = append(jql, fmt.Sprintf(`component in ('%v')`, strings.Join(o.Component, "','")))
	}
	if o.Priority != nil && len(o.Priority) > 0 {
		jql = append(jql, fmt.Sprintf(`priority in ('%v')`, strings.Join(o.Priority, "','")))
	}
	if o.Project != nil && len(o.Project) > 0 {
		jql = append(jql, fmt.Sprintf(`project in ('%v')`, strings.Join(o.Project, "','")))
	}
	if o.Type != nil && len(o.Type) > 0 {
		jql = append(jql, fmt.Sprintf(`type in ('%v')`, strings.Join(o.Type, "','")))
	}
	if o.Status != nil && len(o.Status) > 0 {
		jql = append(jql, fmt.Sprintf(`status in ('%v')`, strings.Join(o.Status, "','")))
	}
	if o.StatusCategory != nil && len(o.StatusCategory) > 0 {
		jql = append(jql, fmt.Sprintf(`statusCategory in ('%v')`, strings.Join(o.StatusCategory, "','")))
	}
	if o.Query != "" {
		jql = append(jql, fmt.Sprintf(`text ~ '%v'`, o.Query))
	}
	result := strings.Join(jql, " AND ")
	if !strings.Contains(result, "ORDER") {
		result = result + " ORDER BY updated DESC"
	}
	return result, nil
}

func ListIssues(host string, opts *ListIssuesOptions) ([]jira.Issue, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	jql, err := opts.toJql(host)
	if err != nil {
		return nil, err
	}

	issues, _, err := client.Issue.Search(jql, &jira.SearchOptions{Fields: opts.Fields})
	return issues, ApiError(err)
}

func GetIssue(host string, issueID string, opts *jira.GetQueryOptions) (*jira.Issue, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	issue, _, err := client.Issue.Get(issueID, opts)
	return issue, ApiError(err)
}
