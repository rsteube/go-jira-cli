package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

var issue_setAliasCmd = &cobra.Command{
	Use:   "set-alias",
	Short: "Create alias commmand for current issue flag values",
	Run: func(cmd *cobra.Command, args []string) {

		flagValues := make(map[string]string)
		cmd.Parent().Flags().VisitAll(func(f *pflag.Flag) {
			flagValues[f.Name] = f.Value.String()
		})
		if out, err := yaml.Marshal(flagValues); err == nil {
			println(string(out))
		}
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
		carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			if cmd, _, err := issue_setAliasCmd.Root().Find(c.Parts); err != nil {
				return carapace.ActionValues()
			} else {
				vals := make([]string, 0)
				for _, subcommand := range cmd.Commands() {
					if !subcommand.Hidden {
						vals = append(vals, subcommand.Name())
					}
				}
				return carapace.ActionValues(vals...).Invoke(c).Suffix("/").ToA()
			}
		}),
	)
}
