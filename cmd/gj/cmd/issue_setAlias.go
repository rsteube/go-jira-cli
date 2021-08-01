package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var issue_setAliasCmd = &cobra.Command{
	Use:   "set-alias [name] [description]",
	Args:  cobra.ExactArgs(2),
	Short: "Create alias commmand for current issue flag values",
	RunE: func(cmd *cobra.Command, args []string) error {
		flagValues := make(map[string]string)
		cmd.Parent().Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				if strings.HasSuffix(f.Value.Type(), "Slice") ||
					strings.HasSuffix(f.Value.Type(), "Array") {
					v := f.Value.String()
					flagValues[f.Name] = v[1 : len(v)-1]
				} else {
					flagValues[f.Name] = f.Value.String()
				}
			}
		})
		return config.AddAlias(args[0], &config.Alias{
			Command:     []string{"issue"},
			Description: args[1],
			Flags:       flagValues,
		})
		//if out, err := yaml.Marshal(flagValues); err == nil {
		//	println(string(out))
		//}
	},
}

type Alias struct {
	Command     []string
	Description string
	Flags       map[string]string
}

func init() {
	issueCmd.AddCommand(issue_setAliasCmd)

	carapace.Gen(issue_setAliasCmd).PositionalCompletion(
		action.ActionConfigAliases(),
	)
}
