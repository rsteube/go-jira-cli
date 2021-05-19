package cmd

import (
	"fmt"

	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

var meta_prioritiesCmd = &cobra.Command{
	Use:   "priorities",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		priorities, err := api.ListPriorities(cmd.Flag("host").Value.String())
		if err != nil {
			return err
		}
		for _, priority := range priorities {
			fmt.Printf("%v %v %v\n", priority.Name, priority.StatusColor, priority.Description)
		}
		return nil
	},
}

func init() {
	metaCmd.AddCommand(meta_prioritiesCmd)
}
