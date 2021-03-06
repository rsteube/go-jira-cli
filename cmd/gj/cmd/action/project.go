package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/go-jira-cli/internal/api"
)

func ActionProjects(host *string, categories []string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(categories) == 0 {
			categories = []string{""}
		}

		action := carapace.ActionValues().Invoke(c)
		for _, category := range categories {
			subaction := carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					if projects, err := api.ListProjects(*host, []string{category}); err != nil {
						return carapace.ActionMessage(err.Error())
					} else {
						vals := make([]string, 0)
						for _, project := range projects {
							vals = append(vals, project.Key, project.Name)
						}
						return carapace.ActionValuesDescribed(vals...)
					}
				}).Cache(24*time.Hour, cache.String(*host, category))
			})
			action = action.Merge(subaction.Invoke(c))
		}
		return action.ToA()
	})
}

func ActionProjectCategories(host *string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if categories, err := api.ListProjectCategories(*host); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				vals := make([]string, 0)
				for _, category := range categories {
					vals = append(vals, category.Name, category.Description)
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		}).Cache(24*time.Hour, cache.String(*host))
	})
}
