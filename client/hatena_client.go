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

type Hatenaclient struct {
	URL        string
	HTTPClient HTTPClient
	Config     *config.Hatena
	Parser     Parser
}

// NewHatenaClient initialize client struct
func NewHatenaClient(c HTTPClient, config *config.Hatena, p Parser) (*Hatenaclient, error) {
	u := HatenaBaseURL + config.Category
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, err
	}

	return &Hatenaclient{URL: u, HTTPClient: c, Config: config, Parser: p}, nil
}

func (c *Hatenaclient) GetHotentryList() (*hatena.EntryList, error) {
	res, err := c.HTTPClient.Get(c.URL)
	if err != nil {
		return nil, errors.New("http request error")
	}

	el, err := c.Parser.ParseToEntryList(res, c.Config.LowerLimit)
	if err != nil {
		return nil, errors.New("http response parse error")
	}
	el.SortByBookmarkUser()

	return el, nil
}
