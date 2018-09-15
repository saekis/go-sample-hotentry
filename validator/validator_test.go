package validator_test

import (
	"testing"

	"github.com/saekis/hotentry-cli/validator"
)

func TestValidate_success(t *testing.T) {
	success_cases := []struct {
		category   string
		lowerlimit int
	}{
		{category: "all", lowerlimit: 100},
		{category: "general", lowerlimit: 100},
		{category: "social", lowerlimit: 100},
		{category: "economics", lowerlimit: 100},
		{category: "life", lowerlimit: 100},
		{category: "knowlegde", lowerlimit: 100},
		{category: "it", lowerlimit: 100},
		{category: "fun", lowerlimit: 100},
		{category: "entertaiment", lowerlimit: 100},
		{category: "game", lowerlimit: 100},
	}

	for _, c := range success_cases {
		if err := validator.Validate(c.lowerlimit, c.category); err != nil {
			t.Errorf("got: %v\nwant: %v", err, nil)
		}
	}

}
