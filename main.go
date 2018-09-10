package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/saekis/go-sample-hotentry/client"
	"github.com/saekis/go-sample-hotentry/config"
	"github.com/saekis/go-sample-hotentry/validator"

	"github.com/pkg/errors"
)

var (
	category   = flag.String("c", "all", "article category")
	lowerlimit = flag.Int("ll", 0, "lower limit of bookmark number")
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()

	if err := validator.Validate(*lowerlimit, *category); err != nil {
		log.Fatal(err)
		return 1
	}

	o := &config.Hatena{LowerLimit: *lowerlimit, Category: *category}
	c, err := client.NewHatenaClient(&http.Client{}, o)
	if err != nil {
		log.Fatal(errors.Wrap(err, "hatena client initialize error"))
		return 1
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
