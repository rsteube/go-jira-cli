package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/go-jira-cli/internal/api"
)

func ActionComponents(host *string, projects []string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		action := carapace.ActionValues().Invoke(c)

		for _, project := range projects {
			subAction := carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if components, err := api.ListComponents(*host, project); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					vals := make([]string, 0)
					for _, component := range components {
						vals = append(vals, component.Name, component.Description)
					}
					return carapace.ActionValuesDescribed(vals...)
				}
			}).Cache(24*time.Hour, cache.String(*host, project))
			action = action.Merge(subAction.Invoke(c))
		}
		return action.ToA()
	})
}
