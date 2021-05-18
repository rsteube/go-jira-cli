package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

func ActionPriorities(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		host := cmd.Flag("host").Value.String()

        return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if priorities, err := api.ListPriorities(host); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				vals := make([]string, 0)
				for _, priority := range priorities {
					vals = append(vals, priority.Name, priority.Description)
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		}).Cache(24*time.Hour, cache.String(host))
	})
}
