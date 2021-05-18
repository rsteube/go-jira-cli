package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

func ActionResolutions(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		host := cmd.Flag("host").Value.String()
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if resolutions, err := api.ListResolutions(host); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				vals := make([]string, 0)
				for _, resolution := range resolutions {
					vals = append(vals, resolution.Name, resolution.Description)
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		}).Cache(24*time.Hour, cache.String(host))
	})
}
