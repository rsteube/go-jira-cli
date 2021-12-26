package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gj",
	Short: "simple jira command line client",
}

func Execute() {
	aliasOverride()
	rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd)
}

func aliasOverride() {
	aliases := config.Aliases()
	aliasCommands := make(map[*cobra.Command]*config.Alias)

	aliases.TraverseSorted(func(name string, alias *config.Alias) error {
		targetCmd := rootCmd
		segments := strings.Split(name, "/")
	segment:
		for index, segment := range segments {
			for _, existingSubCmd := range targetCmd.Commands() {
				if existingSubCmd.Name() == segment {
					targetCmd = existingSubCmd
					continue segment
				}
			}
			subCmd := &cobra.Command{
				Use:                segment,
				Short:              "Alias",
				Run:                func(cmd *cobra.Command, args []string) {},
				DisableFlagParsing: true,
			}
			if index == len(segments)-1 {
				aliasCommands[subCmd] = alias
				subCmd.Short = alias.Description
				subCmd.Run = func(cmd *cobra.Command, args []string) {
					aliasedCmd, aliasedCmdArgs, err := cmd.Root().Find(append(alias.Command, args...))
					if err != nil {
						panic(err.Error()) // TODO
					}

					err = aliasedCmd.Flags().Parse(aliasedCmdArgs)
					if err != nil {
						panic(err.Error())
					}

					for flagName, flagValue := range alias.Flags {
						if flag := aliasedCmd.Flag(flagName); flag != nil && !flag.Changed {
							if err := flag.Value.Set(flagValue); err != nil {
								panic(err.Error())
							}
                            flag.Changed = true
						}
					}

					err = aliasedCmd.RunE(aliasedCmd, aliasedCmd.Flags().Args())
					if err != nil {
						panic(err.Error())
					}
				}
			}
			targetCmd.AddCommand(subCmd)
			targetCmd = subCmd
		}
		targetCmd.Short = alias.Description
		return nil
	})

	if carapace.IsCallback() && len(os.Args) > 3 {
		if targetCmd, targetArgs, err := rootCmd.Find(os.Args[5:]); err == nil {
			if alias := aliasCommands[targetCmd]; alias != nil {
				newArgs := os.Args[:5]
				newArgs = append(newArgs, alias.Command...)
				newArgs = append(newArgs, targetArgs...)
				os.Args = newArgs

				overrideFlags(alias, targetArgs)
			}
		}
	}
}

func overrideFlags(alias *config.Alias, args []string) {
	aliasedCmd, _, err := rootCmd.Find(append(alias.Command, args...))
	if err != nil {
		panic(err.Error)
	}
	for flagName, flagValue := range alias.Flags {
		if flag := aliasedCmd.Flag(flagName); flag != nil {
			if err := flag.Value.Set(flagValue); err != nil {
			}
		}
	}
}
