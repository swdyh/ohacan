package ohacan

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Client struct {
	GroupID string
	Code    string
	URL     string
}

func (c *Client) Request(uv url.Values) (*http.Response, error) {
	var err error
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar}

	u := c.URL
	if u == "" {
		u = "https://ssl.jobcan.jp/"
	}
	authRes, err := client.Get(u + "/m/?code=" + c.Code)
	if err != nil {
		return nil, err
	}
	defer authRes.Body.Close()
	postRes, err := client.PostForm(u+"/m/work/simplestamp", uv)
	if err != nil {
		return nil, err
	}
	defer postRes.Body.Close()
	return postRes, err
}

func (c *Client) ClockIn() (*http.Response, error) {
	uv := url.Values{
		"adit_item": {"出勤"},
		"gps":       {"0"},
		"group_id":  {c.GroupID},
		"reason":    {""},
	}
	return c.Request(uv)
}

func (c *Client) ClockOut() (*http.Response, error) {
	uv := url.Values{
		"adit_item": {"退勤"},
		"gps":       {"0"},
		"group_id":  {c.GroupID},
		"reason":    {""},
	}
	return c.Request(uv)
}
