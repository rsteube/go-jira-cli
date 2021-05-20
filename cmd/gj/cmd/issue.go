package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
)

var issueOpts api.ListIssuesOptions

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Manage issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		return issue_viewCmd.RunE(issue_viewCmd, []string{})
	},
}

func init() {
	issueCmd.PersistentFlags().Bool("web", false, "view in browser")
	issueCmd.PersistentFlags().StringVar(&issueOpts.Host, "host", config.Default().Host, "jira host") // TODO maybe pass var to actions
	issueCmd.PersistentFlags().IntVarP(&issueOpts.Filter, "filter", "f", -1, "predefined filter")
	issueCmd.PersistentFlags().IntVarP(&issueOpts.Limit, "limit", "l", 50, "limit results (default: 50)")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Label, "label", nil, "filter label")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Priority, "priority", nil, "filter priority")
	issueCmd.PersistentFlags().StringSliceVarP(&issueOpts.Assignee, "assignee", "a", nil, "filter assignee")
	issueCmd.PersistentFlags().StringSliceVarP(&issueOpts.Component, "component", "c", nil, "filter component")
	issueCmd.PersistentFlags().StringSliceVarP(&projectOpts.Category, "project", "p", nil, "filter project")
	issueCmd.PersistentFlags().StringSliceVarP(&issueOpts.Resolution, "resolution", "r", nil, "filter resolution")
	issueCmd.PersistentFlags().StringSliceVarP(&issueOpts.Status, "status", "s", nil, "filter status")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.StatusCategory, "status-category", nil, "filter status-category")
	issueCmd.PersistentFlags().StringSliceVarP(&issueOpts.Type, "type", "t", nil, "filter type")
	issueCmd.PersistentFlags().StringVarP(&issueOpts.Jql, "jql", "j", "", "custom jql")
	issueCmd.PersistentFlags().StringVarP(&issueOpts.Query, "query", "q", "", "filter text")
	rootCmd.AddCommand(issueCmd)

	carapace.Gen(issueCmd).FlagCompletion(carapace.ActionMap{
		"assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionUsers(&issueOpts.Host).Invoke(c).Filter(c.Parts).ToA() // TODO assignable users
		}),
		"component": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionComponents(&issueOpts.Host, issueOpts.Project).Invoke(c).Filter(c.Parts).ToA()
		}),
		"resolution": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionResolutions(&issueOpts.Host).Invoke(c).Filter(c.Parts).ToA()
		}),
		"priority": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPriorities(&issueOpts.Host).Invoke(c).Filter(c.Parts).ToA()
		}),
		"project": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjects(&issueOpts.Host, projectOpts.Category).Invoke(c).Filter(c.Parts).ToA()
		}),
		"status": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatuses(&issueOpts.Host).Invoke(c).Filter(c.Parts).ToA()
		}),
		"status-category": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatusCategories(&issueOpts.Host).Invoke(c).Filter(c.Parts).ToA()
		}),
		"type": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionIssueTypes(&issueOpts.Host, issueOpts.Project).Invoke(c).Filter(c.Parts).ToA()
		}),
		"filter": action.ActionFilters(&issueOpts.Host),
		"host":   action.ActionConfigHosts(),
	})
}
