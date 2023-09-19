package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
)

var projectOpts struct {
	Category []string
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		return project_viewCmd.RunE(project_viewCmd, []string{})
	},
}

func init() {
	projectCmd.PersistentFlags().Bool("web", false, "view in browser")
	projectCmd.PersistentFlags().String("host", config.Default().Host, "jira host")
	projectCmd.PersistentFlags().StringSliceVar(&projectOpts.Category, "category", []string{}, "filter category")
	rootCmd.AddCommand(projectCmd)

	carapace.Gen(projectCmd).FlagCompletion(carapace.ActionMap{
		"category": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			host := projectCmd.Flag("host").Value.String()
			return action.ActionProjectCategories(&host).FilterParts().NoSpace()
		}),
		"host": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if hosts, err := config.Hosts(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				vals := make([]string, 0)
				for host := range hosts {
					vals = append(vals, host)
				}
				return carapace.ActionValues(vals...)
			}
		}),
	})
}
