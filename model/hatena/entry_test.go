package hatena_test

import (
	"testing"

	"github.com/saekis/hotentry-cli/model/hatena"
)

func TestToOutputLine_Success(t *testing.T) {
	e := hatena.Entry{
		Fulltitle:     "Google",
		Bookmarkcount: 100,
		URL:           "https://google.com",
	}

	expected := "B!100 Google\nhttps://google.com\n\n"
	if res := e.ToOutputLine(); res != expected {
		t.Errorf("got: %v\nwant: %v", res, expected)
	}
}
