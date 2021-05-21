package cmd

import (
	"fmt"

	"github.com/cli/browser"
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
		host := cmd.Flag("host").Value.String()

		if len(args) == 0 {
			if cmd.Flag("web").Changed {
				return browser.OpenURL(fmt.Sprintf("https://%v/secure/BrowseProjects.jspa?selectedCategory=all&selectedProjectType=all", host))
			}

			projects, err := api.ListProjects(host, projectOpts.Category)
			if err != nil {
				return err
			}
			return output.Pager(func(io *iostreams.IOStreams) error {
				return output.PrintProjectList(io, projects)
			})
		}

		if cmd.Flag("web").Changed {
			return browser.OpenURL(fmt.Sprintf("https://%v/projects/%v/summary", host, args[0]))
		}

		project, err := api.GetProject(host, args[0])
		if err != nil {
			return err
		}

		activities, err := api.ListActivities(host, args[0])
		if err != nil {
			return err
		}

		return output.Pager(func(io *iostreams.IOStreams) error {
			return output.PrintProject(io, project, activities)
		})
	},
}

func init() {
	projectCmd.AddCommand(project_viewCmd)

	carapace.Gen(project_viewCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			host := project_viewCmd.Flag("host").Value.String()
			return action.ActionProjects(&host, projectOpts.Category)
		}),
	)
}
