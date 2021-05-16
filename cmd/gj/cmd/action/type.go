package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

func ActionIssueTypes(cmd *cobra.Command, projects []string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		host := cmd.Flag("host").Value.String()
		action := carapace.ActionValues().Invoke(c)

		for _, project := range projects {
			subAction := carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if issueTypes, err := api.ListIssueTypes(host, project); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					vals := make([]string, 0)
					for _, issueType := range issueTypes {
						vals = append(vals, issueType.Name, issueType.Description)
					}
					return carapace.ActionValuesDescribed(vals...)
				}
			}).Cache(24*time.Hour, cache.String(host, project))
			action = action.Merge(subAction.Invoke(c))
		}
		return action.ToA()
	})
}
