package hatena

import (
	"fmt"
)

type Entry struct {
	Fulltitle     string
	Bookmarkcount int
	URL           string
}

type EntryList []Entry

func (e *Entry) ToOutputLine() string {
	return fmt.Sprintf("%du %s\n%s\n\n", e.Bookmarkcount, e.Fulltitle, e.URL)
}
