package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gj",
	Short: "A brief description of your application",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd)
}
