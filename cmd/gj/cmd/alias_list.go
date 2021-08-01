package cmd

import (
	"fmt"

	"github.com/cli/cli/pkg/iostreams"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/rsteube/go-jira-cli/internal/output"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var alias_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List aliases",
	Run: func(cmd *cobra.Command, args []string) {
		output.Output(func(io *iostreams.IOStreams, cs *iostreams.ColorScheme) error {
			formatted, err := yaml.Marshal(config.Aliases())
			if err != nil {
				return err
			}
			fmt.Fprintln(io.Out, string(formatted))
			return nil
		})
	},
}

func init() {
	aliasCmd.AddCommand(alias_listCmd)
}
