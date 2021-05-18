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

					cs := io.ColorScheme()
					fmt.Fprintf(io.Out, "%v %v [%v]\n%v %v • opened %v • %v comment(s)\nComponents: %v\nLabels: %v\n%v\n",
						cs.Bold(issue.Key),
						cs.Bold(issue.Fields.Summary),
						cs.ColorFromString(strings.SplitN(issue.Fields.Priority.StatusColor, "-", 2)[0])(issue.Fields.Priority.Name),
						cs.ColorFromString(strings.SplitN(issue.Fields.Status.StatusCategory.ColorName, "-", 2)[0])(issue.Fields.Status.Name),
						issue.Fields.Type.Name,
						utils.FuzzyAgo(time.Since(time.Time(issue.Fields.Created))),
						len(issue.Fields.Comments.Comments),
						cs.Gray(strings.Join(components, ", ")),
						cs.Gray(strings.Join(issue.Fields.Labels, ", ")),
						description)

					for index, comment := range issue.Fields.Comments.Comments {
						// TODO optimize
						newest := ""
						if index == len(issue.Fields.Comments.Comments)-1 {
							newest = fmt.Sprintf(" • %v", cs.Cyan("Newest comment"))
						} else if !cmd.Flag("comments").Changed {
							continue
						}

						body, err := markdown.Render(j2m.JiraToMD(comment.Body), "dark") // TODO glamour style from config
						if err != nil {
							return err
						}

						updated, err := time.Parse("2006-01-02T15:04:05Z0700", comment.Updated)
						if err != nil {
							return err
						}

						fmt.Fprintf(io.Out, "%v • %v%v\n%v\n",
							cs.Bold(comment.Author.DisplayName),
							utils.FuzzyAgo(time.Since(time.Time(updated))),
							newest,
							body,
						)
					}

					return nil
				})
			}
		}
	},
}

func init() {
	issue_viewCmd.Flags().Bool("comments", false, "fiew issue comments")
	issueCmd.AddCommand(issue_viewCmd)

	carapace.Gen(issue_viewCmd).PositionalCompletion(
		action.ActionIssues(&issueOpts),
	)
}
