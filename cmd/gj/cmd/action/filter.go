package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/api"
)

func ActionFilters(host *string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if filters, err := api.ListFilters(*host); err != nil {
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
