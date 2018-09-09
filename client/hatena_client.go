package client

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/saekis/go-sample-hotentry/config"
	"github.com/saekis/go-sample-hotentry/model/hatena"

	"github.com/PuerkitoBio/goquery"
)

const HatenaBaseURL = "http://b.hatena.ne.jp/hotentry/"

type hatenaclient struct {
	URL        string
	HTTPClient *http.Client
	config     *config.Hatena
}

// NewHatenaClient initialize client struct
func NewHatenaClient(c *http.Client, config *config.Hatena) (*hatenaclient, error) {
	u := HatenaBaseURL + config.Category
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, err
	}

	return &hatenaclient{URL: u, HTTPClient: c, config: config}, nil
}

func (c *hatenaclient) GetHotentryList() (*hatena.EntryList, error) {
	res, err := c.HTTPClient.Get(c.URL)
	if err != nil {
		return nil, errors.New("http request error")
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, errors.New("document parse error")
	}

	i := 0
	selection := doc.Find("div.entrylist-contents")
	el := make(hatena.EntryList, selection.Length())
	selection.Each(func(index int, s *goquery.Selection) {
		titleanker := s.Find("h3 > a")
		usercount := s.Find(".entrylist-contents-users > a > span").Text()
		count, _ := strconv.Atoi(usercount)
		titlelink, _ := titleanker.Attr("href")
		fulltitle, _ := titleanker.Attr("title")

		if count > c.config.LowerLimit {
			el[i] = hatena.Entry{fulltitle, count, titlelink}
			i++
		}
	})
	el = el[:i]

	return &el, nil
}
