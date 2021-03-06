package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/go-jira-cli/internal/api"
)

func ActionStatuses(host *string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if statuses, err := api.ListStatuses(*host); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					vals := make([]string, 0)
					for _, status := range statuses {
						vals = append(vals, status.Name, status.Description, statusColor(status.StatusCategory.ColorName))
					}
					return carapace.ActionStyledValuesDescribed(vals...)
				}
			}).Cache(1*time.Hour, cache.String(*host))
		})
	})
}

func ActionStatusCategories(host *string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if categories, err := api.ListStatusCategories(*host); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					vals := make([]string, 0)
					for _, category := range categories {
						vals = append(vals, category.Name, category.Key, statusColor(category.ColorName))
					}
					return carapace.ActionStyledValuesDescribed(vals...)
				}
			}).Cache(1*time.Hour, cache.String(*host))
		})
	})
}
