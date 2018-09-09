package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/saekis/go-sample-hotentry/client"
	"github.com/saekis/go-sample-hotentry/config"

	"github.com/pkg/errors"
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

	o := &config.Hatena{LowerLimit: *lawerlimit, Category: *category}
	c, err := client.NewHatenaClient(&http.Client{}, o)
	if err != nil {
		log.Fatal(errors.Wrap(err, "hatena client initialize error"))
	}

	resp, err := c.GetHotentryList()
	if err != nil {
		log.Fatal(errors.Wrap(err, "url fetch error"))
		return 1
	}

	var output string
	for _, v := range *resp {
		output += v.ToOutputLine()
	}

	print(output)

	return 0
}
