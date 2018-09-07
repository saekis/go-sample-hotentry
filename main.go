package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/saekis/go-sample-hotentry/entry"
)

var (
	category   = flag.String("c", "all", "article category")
	lawerlimit = flag.Int("ll", 100, "lower limit of bookmark number")
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	ent, err := entry.Fetch(*category)
	if err != nil {
		log.Fatal(errors.Wrap(err, "url fetch error"))
		return 1
	}

	outputs := make([]string, 100)
	selection := ent.List.Find("div.entrylist-contents")
	selection.Each(func(index int, s *goquery.Selection) {
		titleanker := s.Find("h3 > a")
		usercount := s.Find(".entrylist-contents-users > a > span").Text()
		count, _ := strconv.Atoi(usercount)
		titlelink, _ := titleanker.Attr("href")
		fulltitle, _ := titleanker.Attr("title")

		if count < *lawerlimit {
			outputs = append(outputs, fmt.Sprintf("%du %s\n%s\n\n", count, fulltitle, titlelink))
		}
	})
	sort.Sort(sort.Reverse(sort.StringSlice(outputs)))
	for _, v := range outputs {
		fmt.Print(v)
	}
	return 0
}
