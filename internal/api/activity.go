package api

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ActivityStream struct {
	XMLName xml.Name `xml:"feed"`
	Entry   []struct {
		Title   string `xml:"title"`
		Content string `xml:"content"`
		Author  struct {
			Name string `xml:"name"`
		} `xml:"author"`
		Updated  string `xml:"updated"`
		Category struct {
			Term string `xml:"term,attr"`
		} `xml:"category"`
		Object struct {
			Title   string `xml:"title"`
			Summary string `xml:"summary"`
		} `xml:"object"`
	} `xml:"entry"`
}

func ListActivities(host string, project string) (*ActivityStream, error) {
	client, err := NewClient(host)
	if err != nil {
		return nil, ApiError(err)
	}
	req, err := client.NewRequest("GET", fmt.Sprintf("https://%v/activity?maxResults=10&streams=key+IS+%v", host, project), nil)
	if err != nil {
		return nil, ApiError(err)
	}

	resp, err := client.Do(req, nil)
	if err != nil {
		return nil, ApiError(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
        var as ActivityStream
        err = xml.Unmarshal(bodyBytes, &as)
		if err != nil {
			return nil, err
		}
        return &as, nil
	}
    return nil, errors.New("")
}
