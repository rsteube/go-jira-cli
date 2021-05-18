package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

func ActionFilters(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		host := cmd.Flag("host").Value.String()
		if filters, err := api.ListFilters(host); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for _, filter := range filters {
				vals = append(vals, filter.ID, filter.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
