package cmd

import (
	"fmt"

	"github.com/cli/cli/pkg/iostreams"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/output"
	"github.com/spf13/cobra"
)

var issue_transitionCmd = &cobra.Command{
	Use:   "transition",
	Short: "Change issue state",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output.Output(func(io *iostreams.IOStreams, cs *iostreams.ColorScheme) error {
			if len(args) == 1 {
				transitions, err := api.GetIssueTransitions(issueOpts.Host, args[0])
				if err != nil {
					return err
				}
				return output.Output(func(io *iostreams.IOStreams, cs *iostreams.ColorScheme) error {
					return output.PrintIssueTransitions(io, transitions)
				})
			}

			if state, err := api.DoTransition(issueOpts.Host, args[0], args[1]); err != nil {
				fmt.Fprintf(io.ErrOut, "%s Transition failed %v (%s)\n", cs.Yellow("!"), args[0], err.Error())
			} else {
				fmt.Fprintf(io.ErrOut, "%s Transition successful %s (%s)\n", cs.SuccessIconWithColor(cs.Green), args[0], state.Name)
			}
			return nil
		})
	},
}

func init() {
	issueCmd.AddCommand(issue_transitionCmd)

	carapace.Gen(issue_transitionCmd).PositionalCompletion(
		action.ActionIssues(&issueOpts),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionIssueTransitions(issueOpts.Host, c.Args[0])
		}),
	)
}
