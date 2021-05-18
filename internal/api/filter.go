package api

import (
	"github.com/andygrunwald/go-jira"
)

func GetFilter(host string, filterID int) (*jira.Filter, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	filter, _, err := client.Filter.Get(filterID)
	if err != nil {
		return nil, ApiError(err)
	}
	return filter, nil
}

func ListFilters(host string) ([]*jira.Filter, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	filters, _, err := client.Filter.GetFavouriteList()
	if err != nil {
		return nil, ApiError(err)
	}
	return filters, nil
}
