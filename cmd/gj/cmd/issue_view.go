package cmd

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/StevenACoffman/j2m"
	"github.com/andygrunwald/go-jira"
	"github.com/cli/browser"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/pkg/markdown"
	"github.com/cli/cli/utils"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/output"
	"github.com/spf13/cobra"
)

var issue_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view issue",
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

				return output.PrintIssues(io, issues)
			})
		} else {
			if cmd.Flag("web").Changed { // open in browser
				return browser.OpenURL(fmt.Sprintf("https://%v/browse/%v", issueOpts.Host, args[0]))
			} else {
				issue, err := api.GetIssue(issueOpts.Host, args[0], &jira.GetQueryOptions{})
				if err != nil {
					return err
				}

				return output.Pager(func(io *iostreams.IOStreams) error {
					description, err := markdown.Render(j2m.JiraToMD(issue.Fields.Description), "dark") // TODO glamour style from config
					if err != nil {
						return err
					}

					components := make([]string, len(issue.Fields.Components))
					for index, component := range issue.Fields.Components {
						components[index] = component.Name
					}

					fmt.Fprintf(io.Out, "%v %v [%v]\n%v %v • opened %v • %v comment(s)\nComponents: %v\nLabels: %v\n%v\n",
						io.ColorScheme().Bold(issue.Key),
						io.ColorScheme().Bold(issue.Fields.Summary),
						io.ColorScheme().Gray(issue.Fields.Priority.Name),
						io.ColorScheme().ColorFromString(issue.Fields.Status.StatusCategory.ColorName)(issue.Fields.Status.Name),
						issue.Fields.Type.Name,
						utils.FuzzyAgo(time.Since(time.Time(issue.Fields.Created))),
						len(issue.Fields.Comments.Comments),
						io.ColorScheme().Gray(strings.Join(components, ", ")),
						io.ColorScheme().Gray(strings.Join(issue.Fields.Labels, ", ")),
						description)
					return nil
				})
			}
		}
	},
}

func init() {
	issueCmd.AddCommand(issue_viewCmd)

	carapace.Gen(issue_viewCmd).PositionalCompletion(
		action.ActionIssues(&issueOpts),
	)
}
