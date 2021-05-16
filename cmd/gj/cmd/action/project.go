package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

func ActionProjects(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		host := cmd.Flag("host").Value.String()
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if projects, err := api.ListProjects(host); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				vals := make([]string, 0)
				for _, project := range *projects {
					vals = append(vals, project.Key, project.Name)
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		}).Cache(1*time.Hour, cache.String(host))
	})
}
