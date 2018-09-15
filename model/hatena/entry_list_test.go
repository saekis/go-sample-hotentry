package hatena_test

import (
	"testing"

	"github.com/saekis/hotentry-cli/model/hatena"
)

func TestSortByBookmarkUser_Success(t *testing.T) {
	el := hatena.EntryList{
		hatena.Entry{Bookmarkcount: 100},
		hatena.Entry{Bookmarkcount: 300},
		hatena.Entry{Bookmarkcount: 50},
	}

	el.SortByBookmarkUser()
	if el[0].Bookmarkcount != 50 {
		t.Errorf("got: %v\nwant: %v", el[0].Bookmarkcount, 50)
	}
	if el[1].Bookmarkcount != 100 {
		t.Errorf("got: %v\nwant: %v", el[1].Bookmarkcount, 100)
	}
	if el[2].Bookmarkcount != 300 {
		t.Errorf("got: %v\nwant: %v", el[2].Bookmarkcount, 300)
	}
}
