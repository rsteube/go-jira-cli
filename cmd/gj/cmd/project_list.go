package cmd

import (
	"fmt"

	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

var project_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, err := api.ListProjects(cmd.Flag("host").Value.String())
		if err != nil {
			return err
		}
		for _, project := range *projects {
			fmt.Printf("%v %v\n", project.Key, project.Name)
		}
		return nil
	},
}

func init() {
	projectCmd.AddCommand(project_listCmd)
}
