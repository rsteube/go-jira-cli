package cmd

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rsteube/carapace"
	"github.com/rsteube/go-jira-cli/cmd/gj/cmd/action"
	"github.com/rsteube/go-jira-cli/internal/api"
	"github.com/rsteube/go-jira-cli/internal/config"
	"github.com/spf13/cobra"
)

var authLoginCmd = &cobra.Command{
	Use:   "login [host]",
	Short: "Authenticate with a Jira host",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		host := trimHost(args[0])

		method, err := chooseAuthMethod()
		if err != nil {
			return err
		}

		credentials := &config.Credentials{}
		switch method {
		case "anonymous":

		case "basic":
			credentials.User, credentials.Token, err = enterCredentials()
			if err != nil {
				return err
			}
		case "cookie":
			user, pass, err := enterCredentials()
			if err != nil {
				return err
			}
			credentials.Cookie, err = api.AquireCookie(host, user, pass)
			if err != nil {
				return err
			}
		}

		config, err := config.Hosts()
		if err != nil {
			return err
		}
		config.Add(host, credentials)

		return nil
	},
}

func init() {
	authCmd.AddCommand(authLoginCmd)

	carapace.Gen(authLoginCmd).PositionalCompletion(
		action.ActionConfigHosts(),
	)
}

// trimHost trims prefix/suffix to support copy&and past from browser url
func trimHost(host string) string {
	host = strings.TrimSpace(host)
	host = strings.TrimPrefix(host, "https://")
	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimSuffix(host, "/login.jsp")
	host = strings.TrimSuffix(host, "/secure/Dashboard.jspa")
	return host
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
