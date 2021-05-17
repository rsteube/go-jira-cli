package cmd

import (
	"github.com/spf13/cobra"
)

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	authCmd.AddCommand(authLoginCmd)
}
