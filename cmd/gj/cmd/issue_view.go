package cmd

import (
	"fmt"

	"github.com/StevenACoffman/j2m"
	"github.com/andygrunwald/go-jira"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/pkg/markdown"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/output"
	"github.com/spf13/cobra"
)

var issueViewOpts api.ListIssuesOptions

var issue_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		issue, err := api.GetIssue(cmd.Flag("host").Value.String(), args[0], &jira.GetQueryOptions{})
		if err != nil {
			return err
		}

		return output.Pager(func(io *iostreams.IOStreams) error {
			description, err := markdown.Render(j2m.JiraToMD(issue.Fields.Description), "dark")
			if err != nil {
				return err
			}
			fmt.Fprintf(io.Out, "%v %v %v %v\n%v\n", issue.Key, issue.Fields.Status.Name, issue.Fields.Type.Name, issue.Fields.Summary, description)
			return nil
		})
	},
}

func init() {
	issue_viewCmd.Flags().StringSliceVarP(&issueViewOpts.Project, "project", "p", nil, "filter project")
	issue_viewCmd.Flags().StringSliceVarP(&issueViewOpts.Type, "type", "t", nil, "filter type")
	issue_viewCmd.Flags().StringSliceVarP(&issueViewOpts.Status, "status", "s", nil, "filter status")
	issue_viewCmd.Flags().StringSliceVarP(&issueViewOpts.Assignee, "assignee", "a", nil, "filter assignee")
	issue_viewCmd.Flags().StringVarP(&issueViewOpts.Query, "query", "q", "", "filter text")
	issueCmd.AddCommand(issue_viewCmd)

	carapace.Gen(issue_viewCmd).FlagCompletion(carapace.ActionMap{
		"project": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjects(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"status": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatuses(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"type": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionIssueTypes(issue_listCmd, issueViewOpts.Project).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(issue_viewCmd).PositionalCompletion(
		action.ActionIssues(issue_viewCmd, &issueViewOpts),
	)
}
