package cmd

import (
	"github.com/cli/cli/pkg/iostreams"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/output"
	"github.com/spf13/cobra"
)

var project_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View project(s)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			projects, err := api.ListProjects(cmd.Flag("host").Value.String())
			if err != nil {
				return err
			}
			return output.Pager(func(io *iostreams.IOStreams) error {
				return output.PrintProjectList(io, projects)
			})
		}
		project, err := api.GetProject(cmd.Flag("host").Value.String(), args[0])
		if err != nil {
			return err
		}
		return output.Pager(func(io *iostreams.IOStreams) error {
			return output.PrintProject(io, project)
		})
	},
}

func init() {
	projectCmd.AddCommand(project_viewCmd)

	carapace.Gen(project_viewCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			host := project_viewCmd.Flag("host").Value.String()
			return action.ActionProjects(&host)
		}),
	)
}
