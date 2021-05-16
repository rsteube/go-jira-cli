package action

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

func ActionIssues(cmd *cobra.Command) carapace.Action {
	return carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionProjects(cmd).Invoke(c).Suffix("-").ToA()
		case 1:
			host := cmd.Flag("host").Value.String()
			if issues, err := api.ListIssues(host, &api.ListIssuesOptions{
				Project: []string{c.Parts[0]},
				Fields:  []string{"key", "summary"},
				Search:  api.String(fmt.Sprintf("issue >= %v-%v ORDER BY updatedDate DESC", c.Parts[0], c.CallbackValue))}); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				vals := make([]string, 0)
				for _, issue := range issues {
					vals = append(vals, strings.SplitN(issue.Key, "-", 2)[1], issue.Fields.Summary)
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		default:
			return carapace.ActionValues()
		}
	})
}
