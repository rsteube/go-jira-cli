package cmd

import (
	"fmt"
	"net/url"

	"github.com/andygrunwald/go-jira"
	"github.com/cli/browser"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/output"
	"github.com/spf13/cobra"
)

var issue_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View issue",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if cmd.Flag("web").Changed { // open in browser
				jql, err := issueOpts.ToJql()
				if err != nil {
					return err
				}
				return browser.OpenURL(fmt.Sprintf("https://%v/issues/?jql=%v", issueOpts.Host, url.QueryEscape(jql)))
			}

			issueOpts.Fields = []string{"key", "status", "type", "summary", "components", "updated"}
			issues, err := api.ListIssues(&issueOpts)
			if err != nil {
				return err
			}

			return output.Pager(func(io *iostreams.IOStreams) error {
				return output.PrintIssueList(io, issues)
			})
		} else {
			if cmd.Flag("web").Changed { // open in browser
				return browser.OpenURL(fmt.Sprintf("https://%v/browse/%v", issueOpts.Host, args[0]))
			}
			issue, err := api.GetIssue(issueOpts.Host, args[0], &jira.GetQueryOptions{})
			if err != nil {
				return err
			}

			priorities, err := api.ListPriorities(issueOpts.Host) // TODO cache
			if err != nil {
				return err
			}

			return output.Pager(func(io *iostreams.IOStreams) error {
				return output.PrintIssue(io, issue, priorities, cmd.Flag("comments").Changed)
			})
		}
	},
}

func init() {
	issue_viewCmd.Flags().Bool("comments", false, "view issue comments")
	issueCmd.AddCommand(issue_viewCmd)

	carapace.Gen(issue_viewCmd).PositionalCompletion(
		action.ActionIssues(&issueOpts),
	)
}
