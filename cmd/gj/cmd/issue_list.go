package cmd

import (
	"github.com/cli/cli/pkg/iostreams"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/output"
	"github.com/spf13/cobra"
)

var issueListOpts api.ListIssuesOptions

var issue_listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		issueListOpts.Fields = []string{"key", "status", "type", "summary", "components", "updated"}
		issues, err := api.ListIssues(cmd.Flag("host").Value.String(), &issueListOpts)
		if err != nil {
			return err
		}
		return output.Pager(func(io *iostreams.IOStreams) error {
			return output.PrintIssues(io, issues)
		})
	},
}

func init() {
	issue_listCmd.Flags().StringSliceVarP(&issueListOpts.Project, "project", "p", nil, "filter project")
	issue_listCmd.Flags().StringSliceVarP(&issueListOpts.Type, "type", "t", nil, "filter type")
	issue_listCmd.Flags().StringSliceVarP(&issueListOpts.Status, "status", "s", nil, "filter status")
	issue_listCmd.Flags().StringSliceVarP(&issueListOpts.Assignee, "assignee", "a", nil, "filter assignee")
	issue_listCmd.Flags().StringSliceVarP(&issueListOpts.Component, "component", "c", nil, "filter component")
	issue_listCmd.Flags().StringVarP(&issueListOpts.Query, "query", "q", "", "filter text")

	issueCmd.AddCommand(issue_listCmd)

	carapace.Gen(issue_listCmd).FlagCompletion(carapace.ActionMap{
		"component": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionComponents(issue_listCmd, issueListOpts.Project).Invoke(c).Filter(c.Parts).ToA()
		}),
		"project": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjects(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"status": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatuses(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"type": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionIssueTypes(issue_listCmd, issueListOpts.Project).Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
