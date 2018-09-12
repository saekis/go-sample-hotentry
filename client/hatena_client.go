package client

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/saekis/go-sample-hotentry/config"
	"github.com/saekis/go-sample-hotentry/model/hatena"
)

const HatenaBaseURL = "http://b.hatena.ne.jp/hotentry/"

type HTTPClient interface {
	Get(string) (*http.Response, error)
}

type Parser interface {
	parseResponseToEntryList(*http.Response, int) (*hatena.EntryList, error)
}

type hatenaclient struct {
	URL        string
	HTTPClient HTTPClient
	Config     *config.Hatena
	parser     Parser
}

// NewHatenaClient initialize client struct
func NewHatenaClient(c *http.Client, config *config.Hatena) (*hatenaclient, error) {
	u := HatenaBaseURL + config.Category
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, err
	}

	return &hatenaclient{URL: u, HTTPClient: c, Config: config, parser: Responseparser{}}, nil
}

func (c *hatenaclient) GetHotentryList() (*hatena.EntryList, error) {
	res, err := c.HTTPClient.Get(c.URL)
	if err != nil {
		return nil, errors.New("http request error")
	}

	el, err := c.parser.parseResponseToEntryList(res, c.Config.LowerLimit)
	if err != nil {
		return nil, errors.New("http response parse error")
	}
	el.SortByBookmarkUser()

	return el, nil
}
