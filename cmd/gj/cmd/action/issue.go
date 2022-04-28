package action

import (
	"fmt"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/api"
)

func ActionIssues(opts *api.ListIssuesOptions) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts.Fields = []string{"key", "summary", "status"}

		if issues, err := api.ListIssues(opts); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for _, issue := range issues {
				vals = append(vals, issue.Key, issue.Fields.Summary, statusColor(issue.Fields.Status.StatusCategory.ColorName))
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		}
	})
}

func ActionIssueTransitions(host string, issueID string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if transitions, err := api.GetIssueTransitions(host, issueID); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for _, transition := range transitions {
				vals = append(vals, transition.Name, fmt.Sprintf("%v (%v)", transition.To.Name, transition.To.Description))
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
