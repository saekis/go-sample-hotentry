package client

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/saekis/hotentry-cli/model/hatena"
)

type Parser interface {
	ParseToEntryList(*http.Response, int) (*hatena.EntryList, error)
}

type Responseparser struct{}

func (*Responseparser) ParseToEntryList(res *http.Response, ll int) (*hatena.EntryList, error) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, errors.New("document parse error")
	}

	selection := doc.Find("div.entrylist-contents")

	i := 0
	el := make(hatena.EntryList, selection.Length())
	selection.Each(func(index int, s *goquery.Selection) {
		titleanker := s.Find("h3 > a")
		usercount := s.Find(".entrylist-contents-users > a > span").Text()
		count, _ := strconv.Atoi(usercount)
		titlelink, _ := titleanker.Attr("href")
		fulltitle, _ := titleanker.Attr("title")

		if count > ll {
			el[i] = hatena.Entry{Fulltitle: fulltitle, Bookmarkcount: count, URL: titlelink}
			i++
		}
	})

	el = el[:i]
	return &el, nil
}
