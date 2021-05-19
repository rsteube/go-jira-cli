package cmd

import (
	"fmt"

	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

var meta_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		statuses, err := api.ListStatuses(cmd.Flag("host").Value.String())
		if err != nil {
			return err
		}
		for _, status := range statuses {
			fmt.Printf("%v %v %v\n", status.Name, status.StatusCategory.ColorName, status.StatusCategory.Name)
		}
		return nil
	},
}

func init() {
	metaCmd.AddCommand(meta_statusCmd)
}
