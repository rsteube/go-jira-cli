package api

import (
	"fmt"
	"strings"

	"github.com/andygrunwald/go-jira"
)

type ListIssuesOptions struct {
	Host           string
	Project        []string
	Type           []string
	Status         []string
	StatusCategory []string
	Assignee       []string
	Component      []string
	Label          []string
	Priority       []string
	Resolution     []string
	Fields         []string
	Filter         int
	Jql            string
	Query          string
	Limit          int
}

func (o *ListIssuesOptions) ToJql() (string, error) {
	//project in (SZOPS, BA) AND issuetype in (Bug, CVE) AND status in ("In Progress", Reopened) AND assignee in (membersOf("Interner Benutzer"), membersOf(jira-developers))
	jql := make([]string, 0)

	if o.Filter > 0 {
		filter, err := GetFilter(o.Host, o.Filter)
		if err != nil {
			return "", err
		}
		jql = append(jql, filter.Jql)
	}

	if o.Jql != "" {
		jql = append(jql, o.Jql)
	}
	if o.Assignee != nil && len(o.Assignee) > 0 {
		jql = append(jql, fmt.Sprintf(`assignee in ("%v")`, strings.Join(o.Assignee, `","`)))
	}
	if o.Component != nil && len(o.Component) > 0 {
		jql = append(jql, fmt.Sprintf(`component in ("%v")`, strings.Join(o.Component, `","`)))
	}
	if o.Label != nil && len(o.Label) > 0 {
		jql = append(jql, fmt.Sprintf(`labels in ("%v")`, strings.Join(o.Label, `","`)))
	}
	if o.Priority != nil && len(o.Priority) > 0 {
		jql = append(jql, fmt.Sprintf(`priority in ("%v")`, strings.Join(o.Priority, `","`)))
	}
	if o.Project != nil && len(o.Project) > 0 {
		jql = append(jql, fmt.Sprintf(`project in ("%v")`, strings.Join(o.Project, `","`)))
	}
	if o.Type != nil && len(o.Type) > 0 {
		jql = append(jql, fmt.Sprintf(`type in ("%v")`, strings.Join(o.Type, `","`)))
	}
	if o.Resolution != nil && len(o.Resolution) > 0 {
		jql = append(jql, fmt.Sprintf(`resolution in ("%v")`, strings.Join(o.Resolution, `","`)))
	}
	if o.Status != nil && len(o.Status) > 0 {
		jql = append(jql, fmt.Sprintf(`status in ("%v")`, strings.Join(o.Status, `","`)))
	}
	if o.StatusCategory != nil && len(o.StatusCategory) > 0 {
		jql = append(jql, fmt.Sprintf(`statusCategory in ("%v")`, strings.Join(o.StatusCategory, `","`)))
	}
	if o.Query != "" {
		jql = append(jql, fmt.Sprintf(`text ~ "%v"`, o.Query))
	}
	result := strings.Join(jql, " AND ")
	if !strings.Contains(result, "ORDER") {
		result = result + " ORDER BY updated DESC"
	}
	return result, nil
}

func ListIssues(opts *ListIssuesOptions) ([]jira.Issue, error) {
	client, err := NewClient(opts.Host)
	if err != nil {
		return nil, ApiError(err)
	}
	jql, err := opts.ToJql()
	if err != nil {
		return nil, err
	}

	if opts.Limit <= 0 {
		opts.Limit = 50
	}

	issues := make([]jira.Issue, 0)
	for i := opts.Limit; i > 0; i = i - 50 {
		maxResults := 50
		if i < 50 {
			maxResults = i % 50
		}
		issuesPage, _, err := client.Issue.Search(jql, &jira.SearchOptions{Fields: opts.Fields, MaxResults: maxResults})
		if err != nil {
			return nil, ApiError(err)
		}
		issues = append(issues, issuesPage...)
	}
	return issues, nil
}

func GetIssue(host string, issueID string, opts *jira.GetQueryOptions) (*jira.Issue, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	issue, _, err := client.Issue.Get(issueID, opts)
	return issue, ApiError(err)
}
