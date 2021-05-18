package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/rsteube/go-jira-cli/internal/config"
)

func NewClient(host string) (*jira.Client, error) {
	hosts, err := config.Hosts()
	if err != nil {
		return nil, err
	}
	if auth, ok := hosts[host]; !ok {
		return nil, errors.New("unknown host")
	} else {
		if auth.User != "" && auth.Token != "" {
			tp := &jira.BasicAuthTransport{
				Username: auth.User,
				Password: auth.Token,
			}
			return jira.NewClient(tp.Client(), "https://"+host)
		}
		if auth.Cookie != "" {
			sessionObject := make([]*http.Cookie, 0)
			for _, cookie := range strings.Split(auth.Cookie, ";") {
				splitted := strings.SplitN(strings.TrimSpace(cookie), "=", 2)
				sessionObject = append(sessionObject, &http.Cookie{Name: splitted[0], Value: splitted[1]})
			}

			tp := jira.CookieAuthTransport{
				SessionObject: sessionObject,
			}
			return jira.NewClient(tp.Client(), "https://"+host)
		}
		return jira.NewClient(nil, "https://"+host)
	}
}

func ApiError(err error) error {
	if err != nil {
		return errors.New(strings.SplitN(err.Error(), "\n", 2)[0])
	}
	return err
}

func String(s string) *string {
	return &s
}
