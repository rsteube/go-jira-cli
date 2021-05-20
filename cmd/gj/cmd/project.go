package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		return project_viewCmd.RunE(project_viewCmd, []string{})
	},
}

func init() {
	projectCmd.PersistentFlags().String("host", config.Default().Host, "jira host")
	rootCmd.AddCommand(projectCmd)

	carapace.Gen(projectCmd).FlagCompletion(carapace.ActionMap{
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
