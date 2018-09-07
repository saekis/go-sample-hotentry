package entry

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const REQEST_URL = "http://b.hatena.ne.jp/hotentry/"

type Entries struct {
	List *goquery.Document
}

type Entry struct {
	fulltitle     string
	bookmarkcount int
	url           string
}

func Fetch(category string) (*Entries, error) {
	req, err := http.NewRequest("GET", REQEST_URL+category, nil)
	if err != nil {
		return nil, errors.New("Error occured")
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, errors.New("Error occured")
	}

	return &Entries{List: doc}, nil
}
