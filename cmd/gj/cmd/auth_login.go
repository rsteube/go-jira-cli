package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
)

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		method, err := chooseAuthMethod()
		if err != nil {
			return err
		}

		hostConfig := &config.HostConfig{}
		switch method {
		case "anonymous":

		case "basic":
			hostConfig.User, hostConfig.Token, err = enterCredentials()
			if err != nil {
				return err
			}
		case "cookie":
			user, pass, err := enterCredentials()
			if err != nil {
				return err
			}
			hostConfig.Cookie, err = api.CookieAuth(args[0], user, pass)
			if err != nil {
				return err
			}
		}

		config.AddHost(args[0], hostConfig)

		return nil
	},
}

func init() {
	authCmd.AddCommand(authLoginCmd)

	carapace.Gen(authLoginCmd).PositionalCompletion(
		action.ActionConfigHosts(),
	)
}

func chooseAuthMethod() (method string, err error) {
	prompt := &survey.Select{
		Message: "Choose authentication method:",
		Options: []string{"anonymous", "basic", "cookie"},
	}
	err = survey.AskOne(prompt, &method)
	return
}

func enterCredentials() (name, password string, err error) {
	var qs = []*survey.Question{
		{
			Name:   "name",
			Prompt: &survey.Input{Message: "Enter username:"},
		},
		{
			Name:   "password",
			Prompt: &survey.Password{Message: "Enter token:"},
		},
	}
	answers := struct {
		Name     string
		Password string
	}{}

	err = survey.Ask(qs, &answers)
	return answers.Name, answers.Password, err
}
