package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	issueCmd.PersistentFlags().String("host", "", "jira host")
	rootCmd.AddCommand(issueCmd)

	carapace.Gen(issueCmd).FlagCompletion(carapace.ActionMap{
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
