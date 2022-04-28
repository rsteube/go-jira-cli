package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/go-jira-cli/internal/api"
)

func ActionPriorities(host *string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {

		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if priorities, err := api.ListPriorities(*host); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				vals := make([]string, 0)
				for _, priority := range priorities {
					vals = append(vals, priority.Name, priority.Description, priority.StatusColor)
				}
				return carapace.ActionStyledValuesDescribed(vals...)
			}
		}).Cache(24*time.Hour, cache.String(*host))
	})
}
