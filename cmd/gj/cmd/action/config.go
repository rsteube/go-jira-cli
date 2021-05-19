package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/config"
)

func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if hosts, err := config.Hosts(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for host := range hosts {
				vals = append(vals, host)
			}
			return carapace.ActionValues(vals...)
		}
	})
}
