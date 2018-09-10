package hatena

import (
	"fmt"
)

type Entry struct {
	Fulltitle     string
	Bookmarkcount int
	URL           string
}

func (e *Entry) ToOutputLine() string {
	return fmt.Sprintf("B!%d %s\n%s\n\n", e.Bookmarkcount, e.Fulltitle, e.URL)
}
