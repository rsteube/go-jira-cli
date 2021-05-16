package api

import (
	"fmt"
	"strings"

	"github.com/andygrunwald/go-jira"
)

type ListIssuesOptions struct {
	Project  []string
	Type     []string
	Status   []string
	Assignee []string
    Fields   []string
	Search   *string
}

func (o *ListIssuesOptions) Jql() string {
	//project in (SZOPS, BA) AND issuetype in (Bug, CVE) AND status in ("In Progress", Reopened) AND assignee in (membersOf("Interner Benutzer"), membersOf(jira-developers))
	jql := make([]string, 0)
	if o.Project != nil && len(o.Project) > 0 {
		jql = append(jql, fmt.Sprintf(`project in (%v)`, strings.Join(o.Project, ",")))
	}
	return strings.Join(jql, " ")
}

func ListIssues(host string, opts *ListIssuesOptions) ([]jira.Issue, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	issues, _, err := client.Issue.Search(opts.Jql(), &jira.SearchOptions{Fields: opts.Fields})
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
