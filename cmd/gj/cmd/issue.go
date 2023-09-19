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
	issueCmd.PersistentFlags().IntVarP(&issueOpts.Filter, "filter", "f", -1, "predefined filter")
	issueCmd.PersistentFlags().IntVarP(&issueOpts.Limit, "limit", "l", 50, "limit results (default: 50)")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Label, "label", nil, "filter label")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Priority, "priority", nil, "filter priority")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.StatusCategory, "status-category", nil, "filter status-category")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.NotStatusCategory, "not-status-category", nil, "filter status-category")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Assignee, "assignee", nil, "filter assignee")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Component, "component", nil, "filter component")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Project, "project", nil, "filter project")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Resolution, "resolution", nil, "filter resolution")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Status, "status", nil, "filter status")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.NotStatus, "not-status", nil, "filter status")
	issueCmd.PersistentFlags().StringSliceVar(&issueOpts.Type, "type", nil, "filter type")
	issueCmd.PersistentFlags().StringVar(&issueOpts.Host, "host", config.Default().Host, "jira host") // TODO maybe pass var to actions
	issueCmd.PersistentFlags().StringVarP(&issueOpts.Jql, "jql", "j", "", "custom jql")
	issueCmd.PersistentFlags().StringVarP(&issueOpts.Query, "query", "q", "", "filter text")
	rootCmd.AddCommand(issueCmd)

	carapace.Gen(issueCmd).FlagCompletion(carapace.ActionMap{
		"assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionUsers(&issueOpts.Host).FilterParts().NoSpace() // TODO assignable users
		}),
		"component": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionComponents(&issueOpts.Host, issueOpts.Project).FilterParts().NoSpace()
		}),
		"resolution": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionResolutions(&issueOpts.Host).FilterParts().NoSpace()
		}),
		"priority": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPriorities(&issueOpts.Host).FilterParts().NoSpace()
		}),
		"project": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjects(&issueOpts.Host, projectOpts.Category).FilterParts().NoSpace()
		}),
		"status": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatuses(&issueOpts.Host).FilterParts().NoSpace()
		}),
		"not-status": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatuses(&issueOpts.Host).FilterParts().NoSpace()
		}),
		"status-category": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatusCategories(&issueOpts.Host).FilterParts().NoSpace()
		}),
		"not-status-category": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionStatusCategories(&issueOpts.Host).FilterParts().NoSpace()
		}),
		"type": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionIssueTypes(&issueOpts.Host, issueOpts.Project).FilterParts().NoSpace()
		}),
		"filter": action.ActionFilters(&issueOpts.Host),
		"host":   action.ActionConfigHosts(),
	})
}
