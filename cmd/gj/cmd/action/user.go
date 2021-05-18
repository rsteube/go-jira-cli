package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

func ActionUsers(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		host := cmd.Flag("host").Value.String()

		if users, err := api.FindUsers(host, c.CallbackValue); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for _, user := range users {
				vals = append(vals, user.EmailAddress, user.DisplayName)
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
