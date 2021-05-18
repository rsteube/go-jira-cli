package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/api"
)

func ActionIssues(opts *api.ListIssuesOptions) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts.Fields = []string{"key", "summary"}

		if issues, err := api.ListIssues(opts); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for _, issue := range issues {
				vals = append(vals, issue.Key, issue.Fields.Summary)
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
