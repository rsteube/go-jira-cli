package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
)

var authLogoutCmd = &cobra.Command{
	Use:   "logout [host]",
	Short: "Remove authentication for a Jira host",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.Hosts()
		if err != nil {
			return nil
		}
		return config.Remove(args[0])
	},
}

func init() {
	authCmd.AddCommand(authLogoutCmd)

	carapace.Gen(authLogoutCmd).PositionalCompletion(
		action.ActionConfigHosts(),
	)
}
