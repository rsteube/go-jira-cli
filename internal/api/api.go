package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/rsteube/go-jira-cli/internal/config"
)

func CookieAuth(host string, user string, password string) (map[string]string, error) {
	tp := jira.CookieAuthTransport{
		Username: user,
		Password: password,
		AuthURL:  fmt.Sprintf("https://%v/rest/auth/1/session", host),
	}
	client, err := jira.NewClient(tp.Client(), "https://"+host)
	if err != nil {
		return nil, err
	}

	_, _, err = client.User.GetSelf()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, cookie := range tp.SessionObject {
		if cookie.Name == "atlassian.xsrf.token" || cookie.Name == "JSESSIONID" {
			result[cookie.Name] = cookie.Value
		}
	}
	return result, nil
}

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
		if auth.Cookie != nil {
			sessionObject := make([]*http.Cookie, 0)
			for key, value := range auth.Cookie {
				sessionObject = append(sessionObject, &http.Cookie{Name: key, Value: value})
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
