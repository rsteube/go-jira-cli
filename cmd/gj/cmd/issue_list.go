package cmd

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/spf13/cobra"
)

var opts api.ListIssuesOptions

var issue_listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts.Fields = []string{"key", "status", "type", "summary"}
		issues, err := api.ListIssues(cmd.Flag("host").Value.String(), &opts)
		if err != nil {
			return err
		}
		for _, issue := range issues {
			fmt.Printf("%v %v %v %v\n", issue.Key, issue.Fields.Status.Name, issue.Fields.Type.Name, issue.Fields.Summary)
		}
		return nil
	},
}

func init() {
	issue_listCmd.Flags().StringSliceVarP(&opts.Project, "project", "p", nil, "filter project")
	issue_listCmd.Flags().StringSliceVarP(&opts.Type, "type", "t", nil, "filter project")
	issue_listCmd.Flags().StringSliceVarP(&opts.Status, "status", "s", nil, "filter project")
	issue_listCmd.Flags().StringSliceVarP(&opts.Assignee, "assignee", "a", nil, "filter project")

	issueCmd.AddCommand(issue_listCmd)

	carapace.Gen(issue_listCmd).FlagCompletion(carapace.ActionMap{
		"project": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjects(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"status": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatuses(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
